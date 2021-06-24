package entity

type ClientSignal struct {
	ItemName string            `json:"item_name"`
	Params   map[string]string `json:"params"`
}
