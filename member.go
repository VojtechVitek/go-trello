package trello

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Member struct {
	Id                       string   `json:"id"`
	AvatarHash               string   `json:"avatarHash"`
	Bio                      string   `json:"bio"`
	BioData                  string   `json:"bioData"`
	Confirmed                bool     `json:"confirmed"`
	FullName                 string   `json:"fullName"`
	IdPremOrgsAdmin          string   `json:"idPremOrgsAdmin"`
	Initials                 string   `json:"initials"`
	MemberType               string   `json:"memberType"`
	Products                 []string `json:"products"`
	Status                   string   `json:"status"`
	Url                      string   `json:"url"`
	Username                 string   `json:"username"`
	AvatarSource             string   `json:"avatarSource"`
	Email                    string   `json:"email"`
	GravatarHash             string   `json:"gravatarHash"`
	IdBoards                 []string `json:"idBoards"`
	IdBoardsPinned           []string `json:"idBoardsPinned"`
	IdOrganizations          []string `json:"idOrganizations"`
	LoginTypes               string   `json:"loginTypes"`
	NewEmail                 string   `json:"newEmail"`
	OneTimeMessagesDismissed string   `json:"oneTimeMessagesDismissed"`
	Prefs                    string   `json:"prefs"`
	Trophies                 []string `json:"trophies"`
	UploadedAvatarHash       string   `json:"uploadedAvatarHash"`
	PremiumFeatures          []string `json:"premiumFeatures"`
}

func (c *Client) Member(nick string) (member *Member, err error) {
	req, err := http.NewRequest("GET", c.endpoint+"/member/"+nick, nil)
	if err != nil {
		return
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	} else if resp.StatusCode != 200 {
		err = fmt.Errorf("Received unexpected status %d while trying to retrieve the server data", resp.StatusCode)
		return
	}

	err = json.Unmarshal(body, &member)
	return
}

// TODO: Avatar sizes [170, 30]
func (m *Member) AvatarUrl() string {
	return "https://trello-avatars.s3.amazonaws.com/" + m.AvatarHash + "/170.png"
}
