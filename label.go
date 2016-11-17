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
)

type Label struct {
	client  *Client
	Id      string `json:"id"`
	IdBoard string `json:"idBoard"`
	Name    string `json:"name"`
	Color   string `json:"color"`
	Uses    int    `json:"uses"`
}

func (c *Client) Label(labelId string) (label *Label, err error) {
	body, err := c.Get("/labels/" + labelId)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &label)
	label.client = c
	return
}

func (c *Client) NewLabel(opts Label) (*Label, error) {
	payload := url.Values{}
	payload.Set("name", opts.Name)
	payload.Set("color", opts.Color)
	payload.Set("idBoard", opts.IdBoard)

	body, err := c.Post("/labels", payload)
	if err != nil {
		return nil, err
	}

	var label Label
	if err = json.Unmarshal(body, &label); err != nil {
		return nil, err
	}
	label.client = c
	return &label, nil
}

func (c *Client) DeleteLabel(id string) error {
	_, err := c.Delete("/labels/" + id)
	if err != nil {
		return err
	}
	return nil
}

func (l *Label) Delete() error {
	return l.client.DeleteLabel(l.Id)
}
