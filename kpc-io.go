package kpc

import (
	"encoding/json"
	"io/ioutil"
	"os"
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

func ReadKPCFile(filepath string) (error, *KPC) {

	jsonFile, err := os.Open(filepath)
	// if we os.Open returns an error then handle it
	if err != nil {
		return err, nil
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	return FromJSON(byteValue)
}
