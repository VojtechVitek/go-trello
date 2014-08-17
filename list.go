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

type List struct {
	client  *Client
	Id      string  `json:"id"`
	Name    string  `json:"name"`
	Closed  bool    `json:"closed"`
	IdBoard string  `json:"idBoard"`
	Pos     float32 `json:"pos"`
}

func (l *List) Cards() (cards []Card, err error) {
	req, err := http.NewRequest("GET", l.client.endpoint+"/lists/"+l.Id+"/cards", nil)
	if err != nil {
		return
	}

	resp, err := l.client.client.Do(req)
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

	err = json.Unmarshal(body, &cards)
	for i, _ := range cards {
		cards[i].client = l.client
	}
	return
}
