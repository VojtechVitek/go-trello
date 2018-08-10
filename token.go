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

// Token is a struct that contains information about a Trello oauth1 access token
type Token struct {
	client   *Client
	ID       string `json:"id"`
	MemberID string `json:"idMember"`
	Created  string `json:"dateCreated"`
	Expiry   string `json:"dateExpires"`
}

// GetToken gets information on the token and the retriever of the token with the given tokenID
func (c *Client) GetTokenInfo(tokenID string) (token *Token, err error) {
	body, err := c.Get("/tokens/" + tokenID)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &token)
	token.client = c
	return
}
