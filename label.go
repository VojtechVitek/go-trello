package trello

import (
	"net/url"
	"encoding/json"
)

type Label struct {
	client *Client

	Id string `json:"id"`
	IdBoard string `json:"idBoard"`
	Name string `json:"name"`
	Color string `json:"color"`
	Uses int `json:"uses"`
}

func (l *Label) UpdateName(name string) (err error) {
	payload := url.Values{}
	payload.Set("value", name)

	body, err := l.client.Put("/labels/" + l.Id + "/name", payload)
	if err != nil {
		return
	}
	return json.Unmarshal(body, l)
}

//Color can be null
func (l *Label) UpdateColor(color string) (err error) {
	payload := url.Values{}
	payload.Set("value", color)

	body, err := l.client.Put("/labels/" + l.Id + "/color", payload)
	if err != nil {
		return
	}
	return json.Unmarshal(body, l)
}

func (l *Label) DeleteLabel() error {
	_, err := l.client.Delete("/labels/" + l.Id)
	return err
}
