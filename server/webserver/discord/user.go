package discord

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// User -
type User struct {
	Username      string `json:"username"`
	Verified      bool   `json:"verified"`
	MFAEnabled    bool   `json:"mfa_enabled"`
	ID            string `json:"id"`
	Avatar        string `json:"avatar"`
	Discriminator string `json:"discriminator"`
	Email         string `json:"email"`
}

// GetUserInfo - get user info
func GetUserInfo(accessToken string) (User, error) {
	req, err := http.NewRequest("GET", discordAPI+"/users/@me", nil)

	if err != nil {
		log.Error(err)
		return User{}, err
	}

	req.Header.Add("Authorization", "Bearer "+accessToken)

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Error(err)
		return User{}, err
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Error(err)
		return User{}, err
	}

	var userInfo User

	err = json.Unmarshal(data, &userInfo)

	if err != nil {
		log.Error(err)
		return User{}, err
	}

	// filter guild based on id
	return userInfo, nil
}
