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
	"net/url"
	"strconv"
	"strings"
)

type List struct {
	client  *Client
	Id      string  `json:"id"`
	Name    string  `json:"name"`
	Closed  bool    `json:"closed"`
	IdBoard string  `json:"idBoard"`
	Pos     float32 `json:"pos"`
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

func (l *List) Cards() (cards []Card, err error) {
	body, err := l.client.Get("/lists/" + l.Id + "/cards")
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &cards)
	for i := range cards {
		cards[i].client = l.client
	}
	return
}

func (l *List) Actions() (actions []Action, err error) {
	body, err := l.client.Get("/lists/" + l.Id + "/actions")
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &actions)
	for i := range actions {
		actions[i].client = l.client
	}
	return
}

// AddCard creates with the attributes of the supplied Card struct
// https://developers.trello.com/advanced-reference/card#post-1-cards
func (l *List) AddCard(opts Card) (*Card, error) {
	opts.IdList = l.Id

	payload := url.Values{}
	payload.Set("name", opts.Name)
	payload.Set("desc", opts.Desc)
	payload.Set("pos", strconv.FormatFloat(opts.Pos, 'g', -1, 64))
	payload.Set("due", opts.Due)
	payload.Set("idList", opts.IdList)
	payload.Set("idMembers", strings.Join(opts.IdMembers, ","))

	body, err := l.client.Post("/cards", payload)
	if err != nil {
		return nil, err
	}

	var card Card
	if err = json.Unmarshal(body, &card); err != nil {
		return nil, err
	}
	card.client = l.client
	return &card, nil
}

//If mode is true, list is archived, otherwise it's unarchived (returns to the board)
func (l *List) Archive(mode bool) error {
	payload := url.Values{}
	payload.Set("value", strconv.FormatBool(mode))

	_, err := l.client.Put("/lists/" + l.Id + "/closed", payload)
	return err
}