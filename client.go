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
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	client   *http.Client
	endpoint string
	version  string
	key      string
	token    string
}

func (c *Client) NewRequest(method, url string, body io.Reader) (*http.Request, error) {
	if c.key != "" && c.token != "" {
		return http.NewRequest(method, fmt.Sprintf("%s?key=%s&token=%s", url, c.key, c.token), nil)
	}
	return http.NewRequest(method, url, nil)
}

func NewClient(key, token string) (*Client, error) {
	version := "1"
	endpoint := "https://api.trello.com/" + version

	return &Client{
		client:   http.DefaultClient,
		endpoint: endpoint,
		version:  version,
		key:      key,
		token:    token,
	}, nil
}
