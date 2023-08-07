package models

import (
	"encoding/json"
)

type WhetherConfig struct {
	ChatID      int64 `json:"chat_id"`
	Temperature int64 `json:"temperature"`
}

func (wc *WhetherConfig) Marshal() ([]byte, error) {
	return json.Marshal(wc)
}

func (wc *WhetherConfig) Unmarshal(blob []byte) error {
	return json.Unmarshal(blob, wc)
}
