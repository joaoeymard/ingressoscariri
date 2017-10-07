package utils

import (
	"errors"
	"fmt"
	"strings"
)

var (
	// Errors Conteudo de erros
	Errors = map[string]error{
		"NOT_FOUND":      errors.New(`{"erro":"Registro não encontrado!"}`),
		"METHOD_DEFAULT": errors.New(`{"erro":"O método recebido não foi encontrado!"}`),
	}
)

// ParamsInvalidos Tratamento de erro
func ParamsInvalidos(err error) error {
	return err
}

// ParamsRequired Tratamento para campos obrigatórios
func ParamsRequired(atributo string) error {
	return fmt.Errorf("O atributo %v é obrigatório", atributo)
}

// BancoDados Tratamento de erro
func BancoDados(err error) error {
	if strings.Contains(err.Error(), "pq") {
		retorno := strings.Split(err.Error(), ":")[1][1:]
		retorno = strings.Replace(retorno, "\"", "'", -1)
		err = fmt.Errorf(`{"postgres": "%v"}`, retorno)
	}
	return err
}
