package share

import (
	"encoding/json"
)

type Symbol struct {
	Quote     string `json:"quote"`
	Avalaible string `json:"avalaible"`
}

type Balance struct {
	Currency string   `json:"currency"`
	Symbols  []Symbol `json:"symbols"`
}

func (b Balance) String() string {
	byte, _ := json.Marshal(b)

	return string(byte)
}
