package vk

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

var (
	// NameCases is a list of name cases available for VK
	NameCases = []string{"nom", "gen", "dat", "acc", "ins", "abl"}
)

type (
	// Response from users.get
	Response struct {
		Response []UserInfo
	}
	// UserInfo contains user information
	// TODO improve fields list from here: http://vk.com/dev/fields
	UserInfo struct {
		ID                     int          `json:"id"`
		FirstName              string       `json:"first_name"`
		LastName               string       `json:"last_name"`
		ScreenName             string       `json:"screen_name"`
		NickName               string       `json:"nickname"`
		Sex                    int          `json:"sex,omitempty"`
		Domain                 string       `json:"domain,omitempty"`
		Birthdate              string       `json:"bdate,omitempty"`
		City                   GeoPlace     `json:"city,omitempty"`
		Country                GeoPlace     `json:"country,omitempty"`
		Photo50                string       `json:"photo_50,omitempty"`
		Photo100               string       `json:"photo_100,omitempty"`
		Photo200               string       `json:"photo_200,omitempty"`
		PhotoMax               string       `json:"photo_max,omitempty"`
		Photo200Orig           string       `json:"photo_200_orig,omitempty"`
		PhotoMaxOrig           string       `json:"photo_max_orig,omitempty"`
		HasMobile              bool         `json:"has_mobile,omitempty"`
		Online                 bool         `json:"online,omitempty"`
		CanPost                bool         `json:"can_post,omitempty"`
		CanSeeAllPosts         bool         `json:"can_see_all_posts,omitempty"`
		CanSeeAudio            bool         `json:"can_see_audio,omitempty"`
		CanWritePrivateMessage bool         `json:"can_write_private_message,omitempty"`
		Site                   string       `json:"site,omitempty"`
		Status                 string       `json:"status,omitempty"`
		LastSeen               PlatformInfo `json:"last_seen,omitempty"`
		CommonCount            int          `json:"common_count,omitempty"`
		University             int          `json:"university,omitempty"`
		UniversityName         string       `json:"university_name,omitempty"`
		Faculty                int          `json:"faculty,omitempty"`
		FacultyName            int          `json:"faculty_name,omitempty"`
		Graduation             int          `json:"graduation,omitempty"`
		Relation               int          `json:"relation,omitempty"`
		Universities           []University `json:"universities,omitempty"`
		Schools                []School     `json:"schools,omitempty"`
		Relatives              []Relative   `json:"relatives,omitempty"`
	}
	// GeoPlace contains geographical information like City, Country
	GeoPlace struct {
		ID    int    `json:"id"`
		Title string `json:"title"`
	}
	// PlatformInfo contains information about time and platform
	PlatformInfo struct {
		Time     EpochTime `json:"time"`
		Platform int       `json:"platform"`
	}
	// University contains information about the university
	University struct {
		ID              int    `json:"id"`
		Country         int    `json:"country"`
		City            int    `json:"city"`
		Name            string `json:"name"`
		Faculty         int    `json:"faculty"`
		FacultyName     string `json:"faculty_name"`
		Chair           int    `json:"chair"`
		ChairName       string `json:"chair_name"`
		Graduation      int    `json:"graduation"`
		EducationForm   string `json:"education_form"`
		EducationStatus string `json:"education_status"`
	}
	// School contains information about schools
	School struct {
		ID         int    `json:"id"`
		Country    int    `json:"country"`
		City       int    `json:"city"`
		Name       string `json:"name"`
		YearFrom   int    `json:"year_from"`
		YearTo     int    `json:"year_to"`
		Class      string `json:"class"`
		TypeStr    string `json:"type_str,omitempty"`
		Speciality string `json:"speciality,omitempty"`
	}
	// Relative contains information about relative to the user
	Relative struct {
		ID   int    `json:"id"`   // negative id describes non-existing users (possibly prepared id if they will register)
		Type string `json:"type"` // like `parent`, `grandparent`, `sibling`
		Name string `json:"name,omitempty"`
	}
)

// UsersGet implements method http://vk.com/dev/users.get
//
//     userIds - no more than 1000, use `user_id` or `screen_name`
//     fields - sex, bdate, city, country, photo_50, photo_100, photo_200_orig,
//     photo_200, photo_400_orig, photo_max, photo_max_orig, online,
//     online_mobile, lists, domain, has_mobile, contacts, connections, site,
//     education, universities, schools, can_post, can_see_all_posts,
//     can_see_audio, can_write_private_message, status, last_seen,
//     common_count, relation, relatives, counters
//     name_case - choose one of nom, gen, dat, acc, ins, abl.
//     nom is default
//
func (api *API) UsersGet(userIds []string, fields []string, nameCase string) ([]UserInfo, error) {
	if len(userIds) == 0 {
		return nil, errors.New("you must pass at least one id or screen_name")
	}
	if !ElemInSlice(nameCase, NameCases) {
		return nil, errors.New("the only available name cases are: " + strings.Join(NameCases, ", "))
	}

	endpoint := api.getAPIURL("users.get")
	query := endpoint.Query()
	query.Set("user_ids", strings.Join(userIds, ","))
	query.Set("fields", strings.Join(fields, ","))
	query.Set("name_case", nameCase)
	endpoint.RawQuery = query.Encode()

	var err error
	var resp *http.Response
	var response Response

	if resp, err = http.Get(endpoint.String()); err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response.Response, nil
}
