package mapa

import (
	"encoding/json"
	"io/ioutil"

	"errors"
	"fmt"
)

// Find Retorna o mapa via json
func Find() (map[string]interface{}, error) {

	var (
		jsonMap = make(map[string]interface{})
	)

	jsonFile, err := ioutil.ReadFile("api/v1/model/mapa/mapa.json")

	if err != nil {
		fmt.Println(err)
		return nil, errors.New("Localizar arquivo JSON")
	}

	json.Unmarshal(jsonFile, &jsonMap)

	return jsonMap, nil

}
