package trello

import (
	"encoding/json"
	"fmt"
)

type Webhook struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	IDModel     string `json:"idModel"`
	CallbackURL string `json:"callbackURL"`
	Active      bool   `json:"active"`
}

func (c *Client) Webhooks(token string) (webhooks []Webhook, err error) {

	body, err := c.Get(webhookURL(token))
	if err != nil {
		return []Webhook{}, err
	}
	err = json.Unmarshal(body, &webhooks)
	return
}

func webhookURL(token string) (url string) {

	return fmt.Sprintf("/tokens/%s/webhooks/", token)
}
