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
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Client struct {
	client   *http.Client
	endpoint string
	version  string
}

func (c *Client) do(req *http.Request) ([]byte, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		err = fmt.Errorf("Received unexpected status %d while trying to retrieve the server data with \"%s\"", resp.StatusCode, string(body))
		return nil, err
	}
	return body, nil
}

func (c *Client) Get(resource string) ([]byte, error) {
	req, err := http.NewRequest("GET", c.endpoint+resource, nil)
	if err != nil {
		return nil, err
	}
	return c.do(req)
}

func (c *Client) Post(resource string, data url.Values) ([]byte, error) {
	req, err := http.NewRequest("POST", c.endpoint+resource, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	return c.do(req)
}

func (c *Client) Put(resource string, data url.Values) ([]byte, error) {
	req, err := http.NewRequest("PUT", c.endpoint+resource, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	return c.do(req)
}

func (c *Client) Delete(resource string) ([]byte, error) {
	req, err := http.NewRequest("DELETE", c.endpoint+resource, nil)
	if err != nil {
		return nil, err
	}

	return c.do(req)
}

type bearerRoundTripper struct {
	Delegate http.RoundTripper
	key      string
	token    *string
}

func (b *bearerRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	if b.Delegate == nil {
		b.Delegate = http.DefaultTransport
	}
	values := req.URL.Query()
	values.Add("key", b.key)
	values.Add("token", *b.token)
	req.URL.RawQuery = values.Encode()
	return b.Delegate.RoundTrip(req)
}

// NewBearerTokenTransport will return an http.RoundTripper which will add the
// provided application id and token to API calls.
//   If Delegate is left unset the http.DefaultTransport will be used.
// See https://trello.com/app-key to get your applicationKey
// See https://trello.com/1/connect?key=MYKEYFROMABOVE&name=MYAPPNAME&response_type=token&scope=read,write&expiration=1d
// to get a read/write token good for 1 day
func NewBearerTokenTransport(applicationKey string, token *string) *bearerRoundTripper {
	return &bearerRoundTripper{
		key:   applicationKey,
		token: token,
	}
}

// NewCustomClient can be used to implement your own client
func NewCustomClient(client *http.Client) (*Client, error) {
	version := "1"
	endpoint := "https://api.trello.com/" + version

	return &Client{
		client:   client,
		endpoint: endpoint,
		version:  version,
	}, nil
}

// NewAuthClient will create a trello client which allows authentication. It uses
// NewBearerTokenTransport to create an http.Client which can be used as a trello
// client.
func NewAuthClient(applicationKey string, token *string) (*Client, error) {
	rr := NewBearerTokenTransport(applicationKey, token)
	client := &http.Client{
		Transport: rr,
	}
	return NewCustomClient(client)
}

// NewClient returns a client needed to make trello API calls. If transport is nil
// all API calls will be unauthenticated. If you have a bearer token, NewBearerTokenTransport()
// may be helpful in making calls authenticated.
func NewClient() (*Client, error) {
	return NewCustomClient(http.DefaultClient)
}
