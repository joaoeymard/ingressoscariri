package atributo

import (
	"github.com/JoaoEymard/ingressoscariri/api/v1/utils"
)

var (
	// Filtros lista de filtro para consulta
	Filtros = map[string]string{
		"usuarioID": "USUARIO.id = %v",
		"nome":      "USUARIO.nome ILIKE '%%'|| '%v' ||'%%'",
	}
)

// ValidValues Responsavel por validar os atributos recebidos
func ValidValues(params map[string]interface{}) error {

	var (
		valid = map[string]func(interface{}) error{
			"nome":            nome,
			"senha":           senha,
			"ativo":           ativo,
			"cpf":             cpf,
			"data_nascimento": dataNascimento,
			"sexo":            sexo,
			"nivel":           nivel,
		}
	)

	for key, value := range params {
		if retorno := valid[key](value); retorno != nil {
			return retorno
		}
	}

	return nil

}

func nome(value interface{}) error {

	data, valueType := value.(string)
	if !valueType {
		return utils.ValueInvalidos("nome", "string")
	}

	if data == "" {
		return utils.ValueRequired("nome")
	}

	if len(data) < 3 {
		return utils.ValueMinino("nome", 3)
	}

	if len(data) > 50 {
		return utils.ValueMinino("nome", 50)
	}

	return nil
}

func senha(value interface{}) error {

	data, valueType := value.(string)
	if !valueType {
		return utils.ValueInvalidos("senha", "string")
	}

	if data == "" {
		return utils.ValueRequired("senha")
	}

	return nil
}

func ativo(value interface{}) error {

	_, valueType := value.(bool)
	if !valueType {
		return utils.ValueInvalidos("ativo", "string")
	}

	return nil
}

func cpf(value interface{}) error {

	data, valueType := value.(string)
	if !valueType {
		return utils.ValueInvalidos("cpf", "string")
	}

	if data == "" {
		return utils.ValueRequired("cpf")
	}

	return nil
}

func dataNascimento(value interface{}) error {

	data, valueType := value.(string)
	if !valueType {
		return utils.ValueInvalidos("data_nascimento", "string")
	}

	if data == "" {
		return utils.ValueRequired("data_nascimento")
	}

	return nil
}

func sexo(value interface{}) error {

	data, valueType := value.(string)
	if !valueType {
		return utils.ValueInvalidos("sexo", "string")
	}

	if data == "" {
		return utils.ValueRequired("sexo")
	}

	return nil
}

func nivel(value interface{}) error {

	_, valueType := value.(float64)
	if !valueType {
		return utils.ValueInvalidos("nome", "string")
	}

	return nil
}
