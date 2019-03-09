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
)

type Organization struct {
	client      *Client
	Id          string `json:"id"`
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
	Desc        string `json:"desc"`
	Url         string `json:"url"`
	Website     string `json:"website"`
	LogoHash    string `json:"logoHash"`
}

func (c *Client) Organization(orgId string) (organization *Organization, err error) {
	body, err := c.Get("/organization/" + orgId)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &organization)
	organization.client = c
	return
}

func (o *Organization) Members() (members []Member, err error) {
	body, err := o.client.Get("/organization/" + o.Id + "/members")
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &members)
	for i := range members {
		members[i].client = o.client
	}
	return
}

func (o *Organization) Boards() (boards []Board, err error) {
	body, err := o.client.Get("/organizations/" + o.Id + "/boards")
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &boards)
	for i := range boards {
		boards[i].client = o.client
	}
	return
}
