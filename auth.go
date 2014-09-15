// Package vk implements VKontakte API (including OAuth)
package vk

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// AccessToken response from VK
type AccessToken struct {
	AccessToken      string        `json:"access_token"`
	ExpiresIn        time.Duration `json:"expires_in"`
	UserID           int           `json:"user_id"`
	UserEmail        string        `json:"email"`
	Error            string        `json:"error"`
	ErrorDescription string        `json:"error_description"`
}

// AuthURL generates URL to authenticate via OAuth
func (api *API) AuthURL(state string) string {

	query := api.requestTokenURL.Query()
	query.Set("client_id", api.AppID)
	if len(api.Scope) > 0 {
		query.Set("scope", strings.Join(api.Scope, ","))
	}
	query.Set("redirect_uri", api.callbackURL.String())
	query.Set("display", "page")
	query.Set("v", Version)
	query.Set("response_type", "code")
	api.requestTokenURL.RawQuery = query.Encode()

	return api.requestTokenURL.String()
}

// Authenticate with API
func (api *API) Authenticate(code string) error {
	var resp *http.Response
	var err error
	var tok AccessToken

	query := api.accessTokenURL.Query()
	query = url.Values{
		"client_id":     {api.AppID},
		"client_secret": {api.Secret},
		"code":          {code},
		"redirect_uri":  {api.callbackURL.String()},
	}
	api.accessTokenURL.RawQuery = query.Encode()

	if resp, err = http.Get(api.accessTokenURL.String()); err != nil {
		return err
	}
	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(&tok); err != nil {
		return err
	}

	if tok.Error != "" {
		return errors.New(tok.ErrorDescription)
	}

	tok.ExpiresIn *= time.Second
	api.UserID = strconv.Itoa(tok.UserID)
	api.UserEmail = tok.UserEmail
	api.AccessToken = tok.AccessToken
	api.Expiry = time.Now().Add(tok.ExpiresIn)

	return nil
}
