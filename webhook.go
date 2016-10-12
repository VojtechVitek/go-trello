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

func (c *Client) Webhook(webhookID string) (webhook Webhook, err error) {

	url := fmt.Sprintf("/webhooks/%s/", webhookID)
	body, err := c.Get(url)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &webhook)
	return
}

func webhookURL(token string) (url string) {

	return fmt.Sprintf("/tokens/%s/webhooks/", token)
}
