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
	"net/url"
)

type Webhook struct {
	Description string  `json:"description"`
	CallbackURL string  `json:"callbackURL"`
	IdModel     string  `json:"idModel"`
	client      *Client `json:"-"`
}

func (c *Client) NewWebhook(cb, i string) (webhook *Webhook) {
	webhook = &Webhook{
		Description: fmt.Sprintf("Events for model \"%s\"", i),
		CallbackURL: cb,
		IdModel:     i,
		client:      c,
	}
	return
}

func (c *Client) Webhook(webhookId string) (webhook *Webhook, err error) {
	req, err := c.NewRequest("GET", c.endpoint+"/webhooks/"+webhookId, nil)
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

	err = json.Unmarshal(body, &webhook)
	webhook.client = c
	return
}

func (w *Webhook) Post() (body []byte, err error) {
	payload := url.Values{}
	payload.Set("description", w.Description)
	payload.Add("callbackURL", w.CallbackURL)
	payload.Add("idModel", w.IdModel)

	body, err = w.client.PostForm("/webhooks/", payload)
	return
}
