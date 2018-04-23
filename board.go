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

import "encoding/json"

type Board struct {
	client   *Client
	Id       string `json:"id"`
	Name     string `json:"name"`
	Desc     string `json:"desc"`
	DescData struct {
		Emoji struct{} `json:"emoji"`
	} `json:"descData"`
	Closed         bool   `json:"closed"`
	IdOrganization string `json:"idOrganization"`
	Pinned         bool   `json:"pinned"`
	Url            string `json:"url"`
	ShortUrl       string `json:"shortUrl"`
	Prefs          struct {
		PermissionLevel       string            `json:"permissionLevel"`
		Voting                string            `json:"voting"`
		Comments              string            `json:"comments"`
		Invitations           string            `json:"invitations"`
		SelfJoin              bool              `json:"selfjoin"`
		CardCovers            bool              `json:"cardCovers"`
		CardAging             string            `json:"cardAging"`
		CalendarFeedEnabled   bool              `json:"calendarFeedEnabled"`
		Background            string            `json:"background"`
		BackgroundColor       string            `json:"backgroundColor"`
		BackgroundImage       string            `json:"backgroundImage"`
		BackgroundImageScaled []BoardBackground `json:"backgroundImageScaled"`
		BackgroundTile        bool              `json:"backgroundTile"`
		BackgroundBrightness  string            `json:"backgroundBrightness"`
		CanBePublic           bool              `json:"canBePublic"`
		CanBeOrg              bool              `json:"canBeOrg"`
		CanBePrivate          bool              `json:"canBePrivate"`
		CanInvite             bool              `json:"canInvite"`
	} `json:"prefs"`
	LabelNames struct {
		Red    string `json:"red"`
		Orange string `json:"orange"`
		Yellow string `json:"yellow"`
		Green  string `json:"green"`
		Blue   string `json:"blue"`
		Purple string `json:"purple"`
	} `json:"labelNames"`
}

type BoardBackground struct {
	width  int    `json:"width"`
	height int    `json:"height"`
	url    string `json:"url"`
}

func (c *Client) Boards() (boards []Board, err error) {
	body, err := c.Get("/members/me/boards/")
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &boards)
	for i := range boards {
		boards[i].client = c
	}
	return
}

func (c *Client) Board(boardId string) (board *Board, err error) {
	body, err := c.Get("/boards/" + boardId)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &board)
	board.client = c
	return
}

func (b *Board) Lists() (lists []List, err error) {
	body, err := b.client.Get("/boards/" + b.Id + "/lists")
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &lists)
	for i := range lists {
		lists[i].client = b.client
	}
	return
}

func (b *Board) Members() (members []Member, err error) {
	body, err := b.client.Get("/boards/" + b.Id + "/members")
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &members)
	for i := range members {
		members[i].client = b.client
	}
	return
}

func (b *Board) Cards() (cards []Card, err error) {
	body, err := b.client.Get("/boards/" + b.Id + "/cards")
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &cards)
	for i := range cards {
		cards[i].client = b.client
	}
	return
}

func (b *Board) Card(IdCard string) (card *Card, err error) {
	body, err := b.client.Get("/boards/" + b.Id + "/cards/" + IdCard)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &card)
	card.client = b.client
	return
}

func (b *Board) Checklists() (checklists []Checklist, err error) {
	body, err := b.client.Get("/boards/" + b.Id + "/checklists")
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &checklists)
	for i := range checklists {
		checklists[i].client = b.client
	}
	return
}

func (b *Board) MemberCards(IdMember string) (cards []Card, err error) {
	body, err := b.client.Get("/boards/" + b.Id + "/members/" + IdMember + "/cards")
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &cards)
	for i := range cards {
		cards[i].client = b.client
	}
	return
}

func (b *Board) Actions(arg ...*Argument) (actions []Action, err error) {
	ep := "/boards/" + b.Id + "/actions"
	if query := EncodeArgs(arg); query != "" {
		ep += "?" + query
	}

	body, err := b.client.Get(ep)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &actions)
	for i := range actions {
		actions[i].client = b.client
	}
	return
}
