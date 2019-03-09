package trello

type Plugin struct {
	ID       string `json:"id"`
	BoardID  string `json:"idBoard"`
	PluginID string `json:"idPlugin"`
}

type PluginData struct {
	ID       string `json:"id"`
	PluginID string `json:"idPlugin"`
	Scope    string `json:"scope"`
	ModelID  string `json:"idModel"`
	Value    string `json:"value"`
	Access   string `json:"access"`
}
