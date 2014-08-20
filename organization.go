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

type Organization struct {
	client      *Client
	Id          string   `json:"id"`
	Name        string   `json:"name"`
	DisplayName string   `json:"displayName"`
	Desc        string   `json:"desc"`
	DescData    string   `json:"descData"`
	Url         string   `json:"url"`
	Website     string   `json:"website"`
	LogoHash    string   `json:"logoHash"`
	Products    []string `json:"products"`
	PowerUps    []string `json:"powerUps"`
}

func (c *Client) Organization(orgId string) (organization *Organization, err error) {
	req, err := http.NewRequest("GET", c.endpoint+"/organization/"+orgId, nil)
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

	err = json.Unmarshal(body, &organization)
	organization.client = c
	return
}

func (o *Organization) Members() (members []Member, err error) {
	req, err := http.NewRequest("GET", o.client.endpoint+"/organization/"+o.Id+"/members", nil)
	if err != nil {
		return
	}

	resp, err := o.client.client.Do(req)
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

	err = json.Unmarshal(body, &members)
	for i, _ := range members {
		members[i].client = o.client
	}
	return
}
