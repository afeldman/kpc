package kpc

import (
	"encoding/json"
)

func (karelpackageconfig *KPC) ToJSON() (error, []byte) {
	b, err := json.Marshal(karelpackageconfig)
	if err != nil {
		return err, []byte{}
	}
	return nil, b
}

func FromJSON(json_str []byte) (error, *KPC) {
	var kpc KPC
	err := json.Unmarshal(json_str, &kpc)
	if err != nil {
		return err, nil
	}
	return nil, &kpc
}
