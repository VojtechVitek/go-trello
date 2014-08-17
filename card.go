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
	"fmt"
	"io/ioutil"
	"net/http"
)

type Card struct {
	client                *Client
	Id                    string   `json:"id"`
	CheckItemStates       string   `json:"checkItemStates"`
	Closed                bool     `json:"closed"`
	DateLastActivity      string   `json:"dateLastActivity"`
	Desc                  string   `json:"desc"`
	DescData              []string `json:"descData"`
	Email                 string   `json:"email"`
	IdBoard               string   `json:"idBoard"`
	IdList                string   `json:"idList"`
	IdMembersVoted        []string `json:"idMembersVoted"`
	IdShort               int      `json:"idShort"`
	IdAttachmentCover     string   `json:"idAttachmentCover"`
	ManualCoverAttachment bool     `json:"manualCoverAttachment"`
	Name                  string   `json:"name"`
	Pos                   float32  `json:"pos"`
	ShortLink             string   `json:"shortLink"`
	Badges                struct {
		Votes              int    `json:"votes"`
		ViewingMemberVoted bool   `json:"viewingMemberVoted"`
		Subscribed         bool   `json:"subscribed"`
		Fogbugz            string `json:"fogbugz"`
		CheckItems         int    `json:"checkItems"`
		CheckItemsChecked  int    `json:"checkItemsChecked"`
		Comments           int    `json:"comments"`
		Attachments        int    `json:"attachments"`
		Description        string `json:"description"`
		Due                string `json:"due"`
	} `json:"badges"`
	Due          string   `json:"due"`
	IdCheckLists []string `json:"idCheckLists"`
	IdMembers    []string `json:"idMembers"`
	Labels       []string `json:"labels"`
	ShortUrl     string   `json:"shortUrl"`
	Subscribed   bool     `json:"subscribed"`
	Url          string   `json:"url"`
}

func (c *Client) Card(CardId string) (card *Card, err error) {
	req, err := http.NewRequest("GET", c.endpoint+"/card/"+CardId, nil)
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

	err = json.Unmarshal(body, &card)
	card.client = c
	return
}

func (c *Card) Checklists() (checklists []Checklist, err error) {
	req, err := http.NewRequest("GET", c.client.endpoint+"/card/"+c.Id+"/checklists", nil)
	if err != nil {
		return
	}

	resp, err := c.client.client.Do(req)
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

	err = json.Unmarshal(body, &checklists)
	for i, _ := range checklists {
		checklists[i].client = c.client
	}
	return
}
