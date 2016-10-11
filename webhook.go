package trello

type Webhook struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	IDModel     string `json:"idModel"`
	CallbackURL string `json:"callbackURL"`
	Active      bool   `json:"active"`
}
