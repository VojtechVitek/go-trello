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
		Todo string `json:"todo"`
	} `json:"descData"`
	Closed         bool   `json:"closed"`
	IdOrganization string `json:"idOrganization"`
	Pinned         bool   `json:"pinned"`
	Starred        bool   `json:"starred"`
	Url            string `json:"url"`
	ShortUrl       string `json:"shortUrl`
	Prefs          struct {
		PermissionLevel      string `json:"permissionLevel"`
		Voting               string `json:"voting"`
		Comments             string `json:"comments"`
		Invitations          string `json:"invitations"`
		SelfJoin             bool   `json:"selfjoin"`
		CardCovers           bool   `json:"cardCovers"`
		CardAging            string `json:"cardAging"`
		CalendarFeedEnabled  bool   `json:"calendarFeedEnabled"`
		Background           string `json:"background"`
		BackgroundColor      string `json:"backgroundColor"`
		BackgroundImage      string `json:"backgroundImage"`
		BackgroundTile       bool   `json:"backgroundTile"`
		BackgroundBrightness string `json:"backgroundBrightness"`
		CanBePublic          bool   `json:"canBePublic"`
		CanBeOrg             bool   `json:"canBeOrg"`
		CanBePrivate         bool   `json:"canBePrivate"`
		CanInvite            bool   `json:"canInvite"`
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

func (c *Client) Boards() (boards []Board, err error) {
	body, err := c.Get("/boards/")
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &boards)
	for i, _ := range boards {
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

func (c *Client) List(listId string) (list *List, err error) {
	body, err := c.Get("/lists/" + listId)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &list)
	list.client = c
	return
}

func (b *Board) Lists() (lists []List, err error) {
	body, err := b.client.Get("/boards/" + b.Id + "/lists")
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &lists)
	for i, _ := range lists {
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
	for i, _ := range members {
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
	for i, _ := range cards {
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
	for i, _ := range checklists {
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
	for i, _ := range cards {
		cards[i].client = b.client
	}
	return
}
