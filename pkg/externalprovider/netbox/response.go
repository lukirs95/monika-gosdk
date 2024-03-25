package netbox

import "encoding/json"

type Response struct {
	Count   int             `json:"count"`
	Next    string          `json:"next"`
	Results json.RawMessage `json:"results"`
}
