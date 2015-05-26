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
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Client struct {
	client   *http.Client
	endpoint string
	version  string
	key      string
	token    string
}

func (c *Client) QueryURL(url string) (authurl string) {
	if c.key != "" && c.token != "" {
		authurl = fmt.Sprintf("%s?key=%s&token=%s", url, c.key, c.token)
		return
	}
	authurl = url
	return
}

func (c *Client) Get(resource string) (body []byte, err error) {
	req, err := c.NewRequest(
		"GET",
		c.endpoint+resource,
		nil,
	)
	if err != nil {
		return
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	} else if resp.StatusCode != 200 {
		err = fmt.Errorf("Received unexpected status %d while trying to retrieve the server data with \"%s\"", resp.StatusCode, string(body))
		return
	}

	return
}

func (c *Client) PostForm(resource string, data url.Values) (body []byte, err error) {
	req, err := c.NewRequest(
		"POST",
		c.endpoint+resource,
		strings.NewReader(data.Encode()),
	)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.client.Do(req)
	if err != nil {
		return
	}
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		err = fmt.Errorf("Received unexpected status %d while trying to retrieve the server data with: \"%s\"", resp.StatusCode, body)
		return
	}
	return
}

func (c *Client) NewRequest(method, url string, body io.Reader) (*http.Request, error) {
	return http.NewRequest(method, c.QueryURL(url), body)
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
