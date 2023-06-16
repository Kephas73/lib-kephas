package document

import "encoding/json"

type ResponseElastic struct {
	Took int
	Hits struct {
		Total struct {
			Value int `json:"value"`
		}
		Hits []struct {
			ID         string          `json:"_id"`
			Source     json.RawMessage `json:"_source"`
			Highlights json.RawMessage `json:"highlight"`
			Sort       []interface{}   `json:"sort"`
		}
	}
}
