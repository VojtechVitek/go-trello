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
	"net/url"
	"strconv"
)

type ChecklistItem struct {
	client   *Client
	listID   string // back pointer to the parent Id
	State    string `json:"state"`
	Id       string `json:"id"`
	Name     string `json:"name"`
	NameData struct {
		Emoji struct{} `json:"emoji"`
	} `json:"nameData"`
	Pos int `json:"pos"`
}

func (i *ChecklistItem) Delete() error {
	_, err := i.client.Delete("/checklists/" + i.listID + "/checkItems/" + i.Id)
	return err
}

// Checklist is a representation of a checklist on a trello card
// https://developers.trello.com/advanced-reference/checklist
type Checklist struct {
	client     *Client
	Id         string          `json:"id"`
	Name       string          `json:"name"`
	IdBoard    string          `json:"idBoard"`
	IdCard     string          `json:"idCard"`
	Pos        float32         `json:"pos"`
	CheckItems []ChecklistItem `json:"checkItems"`
}

// Delete will delete the checklist
// https://developers.trello.com/advanced-reference/checklist#delete-1-checklists-idchecklist
func (c *Checklist) Delete() error {
	_, err := c.client.Delete("/checklists/" + c.Id)
	return err
}

// AddItem will add a new item to the given checklist. The position will default to 'bottom'
// if nil and the item will default to 'unchecked'.
//   name must have a length 1 <= length <= 16384
//   pos can take the values 'top', 'bottom', or a positive integer
// https://developers.trello.com/advanced-reference/checklist#post-1-checklists-idchecklist-checkitems
func (c *Checklist) AddItem(name string, pos *string, checked *bool) (*ChecklistItem, error) {
	payload := url.Values{}
	if len(name) < 1 || len(name) > 16384 {
		return nil, fmt.Errorf("Checklist item name %q has invalid length. 1 <= length <= 16384", name)
	}
	payload.Set("name", name)
	if pos != nil {
		if *pos != "top" && *pos != "bottom" {
			i, err := strconv.Atoi(*pos)
			if err != nil {
				return nil, err
			}
			if i < 1 {
				return nil, fmt.Errorf("Checklist item position %q is invalid. Only 'top', 'bottom', or a positive integer", *pos)
			}
		}
		payload.Set("pos", *pos)
	}
	if checked != nil {
		payload.Set("checked", strconv.FormatBool(*checked))
	}
	body, err := c.client.Post("/checklist/"+c.Id+"/checkItems", payload)
	if err != nil {
		return nil, err
	}

	item := &ChecklistItem{}
	if err = json.Unmarshal(body, item); err != nil {
		return nil, err
	}
	item.client = c.client
	item.listID = c.Id

	return item, err
}
