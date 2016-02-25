/*
Copyright 2014 go-trello authors. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package trello

import (
	"encoding/json"
	"strings"
)

type Member struct {
	client     *Client
	Id         string `json:"id"`
	AvatarHash string `json:"avatarHash"`
	Bio        string `json:"bio"`
	BioData    struct {
		Emoji interface{} `json:"emoji,omitempty"`
	} `json:"bioData"`
	Confirmed                bool     `json:"confirmed"`
	FullName                 string   `json:"fullName"`
	IdPremOrgsAdmin          []string `json:"idPremOrgsAdmin"`
	Initials                 string   `json:"initials"`
	MemberType               string   `json:"memberType"`
	Products                 []int    `json:"products"`
	Status                   string   `json:"status"`
	Url                      string   `json:"url"`
	Username                 string   `json:"username"`
	AvatarSource             string   `json:"avatarSource"`
	Email                    string   `json:"email"`
	GravatarHash             string   `json:"gravatarHash"`
	IdBoards                 []string `json:"idBoards"`
	IdBoardsPinned           []string `json:"idBoardsPinned"`
	IdOrganizations          []string `json:"idOrganizations"`
	LoginTypes               []string `json:"loginTypes"`
	NewEmail                 string   `json:"newEmail"`
	OneTimeMessagesDismissed []string `json:"oneTimeMessagesDismissed"`
	Prefs                    struct {
		SendSummaries                 bool   `json:"sendSummaries"`
		MinutesBetweenSummaries       int    `json:"minutesBetweenSummaries"`
		MinutesBeforeDeadlineToNotify int    `json:"minutesBeforeDeadlineToNotify"`
		ColorBlind                    bool   `json:"colorBlind"`
		Locale                        string `json:"locale"`
	} `json:"prefs"`
	Trophies           []string `json:"trophies"`
	UploadedAvatarHash string   `json:"uploadedAvatarHash"`
	PremiumFeatures    []string `json:"premiumFeatures"`
}

func (c *Client) Member(nick string) (member *Member, err error) {
	body, err := c.Get("/members/" + nick)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &member)
	member.client = c
	return
}

func (m *Member) Boards(field ...string) (boards []Board, err error) {
	fields := ""
	if len(field) == 0 {
		fields = "all"
	} else {
		fields = strings.Join(field, ",")
	}

	body, err := m.client.Get("/members/" + m.Id + "/boards?fields=" + fields)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &boards)
	for i := range boards {
		boards[i].client = m.client
	}
	return
}

func (m *Member) Notifications() (notifications []Notification, err error) {
	body, err := m.client.Get("/members/" + m.Id + "/notifications")
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &notifications)
	for i := range notifications {
		notifications[i].client = m.client
	}
	return
}

// TODO: Avatar sizes [170, 30]
func (m *Member) AvatarUrl() string {
	return "https://trello-avatars.s3.amazonaws.com/" + m.AvatarHash + "/170.png"
}
