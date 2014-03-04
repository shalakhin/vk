package vk

import (
	"net/url"
	"time"
)

var (
	// Version of VK API
	Version = "5.12"
	// APIURL is a base to make API calls
	APIURL = "https://api.vk.com/method/"
)

// API holds data to use for communication
type API struct {
	AppID           string
	Secret          string
	Scope           []string
	AccessToken     string
	Expiry          time.Time
	UserID          string
	callbackURL     *url.URL
	requestTokenURL *url.URL
	accessTokenURL  *url.URL
}

// NewAPI creates instance of API
func NewAPI(appID, secret string, scope []string, callback string) *API {
	var err error
	var callbackURL *url.URL

	if appID == "" {
		return nil
	}
	if secret == "" {
		return nil
	}
	if callbackURL, err = url.Parse(callback); err != nil {
		return nil
	}
	reqTokURL, _ := url.Parse("https://oauth.vk.com/authorize")
	accTokURL, _ := url.Parse("https://oauth.vk.com/access_token")

	return &API{
		AppID:           appID,
		Secret:          secret,
		Scope:           scope,
		callbackURL:     callbackURL,
		requestTokenURL: reqTokURL,
		accessTokenURL:  accTokURL,
	}
}

// getAPIURL prepares URL instance with defined method
func getAPIURL(method string) *url.URL {
	apiURL, _ := url.Parse(APIURL + method)
	return apiURL
}
