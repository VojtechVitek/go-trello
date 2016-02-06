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

type Notification struct {
	client *Client
	Id     string `json:"id"`
	Unread bool   `json:"unread"`
	Type   string `json:"type"`
	Date   string `json:"date"`
	Data   struct {
		ListBefore struct {
			Id   string `json:"id"`
			Name string `json:"name"`
		} `json:"listBefore"`
		ListAfter struct {
			Id   string `json:"id"`
			Name string `json:"name"`
		} `json:"listAfter"`
		Board struct {
			Id        string `json:"id"`
			Name      string `json:"name"`
			ShortLink string `json:"shortLink"`
		} `json:"board"`
		Card struct {
			Id        string `json:"id"`
			Name      string `json:"name"`
			ShortLink string `json:"shortLink"`
			IdShort   int    `json:"idShort"`
		} `json:"card"`
		Old struct {
			IdList string `json:"idList"`
		} `json:"old"`
	} `json:"data"`
	IdMemberCreator string `json:"idMemberCreator"`
	MemberCreator   struct {
		Id         string `json:"id"`
		AvatarHash string `json:"avatarHash"`
		FullName   string `json:"fullName"`
		Initials   string `json:"initials"`
		Username   string `json:"username"`
	} `json:"memberCreator"`
}

func (c *Client) Notification(notificationId string) (notification *Notification, err error) {
	body, err := c.Get("/notifications/" + notificationId)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &Notification)
	notification.client = c
	return
}
