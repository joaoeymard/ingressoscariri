package eventos

import (
	"encoding/json"
	"io/ioutil"

	"errors"
	"fmt"
)

// FindAll Retorna os eventos via json
func FindAll() (map[string]interface{}, error) {

	var (
		jsonEventos = make(map[string]interface{})
	)

	jsonFile, err := ioutil.ReadFile("api/v1/model/eventos/eventos.json")

	if err != nil {
		fmt.Println(err)
		return nil, errors.New("Localizar arquivo JSON")
	}

	json.Unmarshal(jsonFile, &jsonEventos)

	return jsonEventos, nil

}
