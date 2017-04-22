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

type Attachment struct {
	client    *Client
	Id        string `json:"id"`
	Bytes     int    `json:"bytes"`
	Date      string `json:"date"`
	EdgeColor string `json:"edgeColor"`
	IdMember  string `json:"idMember"`
	IsUpload  bool   `json:"isUpload"`
	MimeType  string `json:"mimeType"`
	Name      string `json:"name"`
	Previews  []struct {
		Width  int    `json:"width"`
		Height int    `json:"height"`
		Url    string `json:"url"`
		Bytes  int    `json:"bytes"`
		Id     string `json:"_id"`
		Scaled bool   `json:"scaled"`
	} `json:"previews"`
	Url string `json:"url"`
}
