package utils

import (
	"errors"
	"fmt"
	"strings"
)

var (
	// Errors Conteudo de erros
	Errors = map[string]error{
		"NOT_FOUND": errors.New(`{"erro":"Registro não encontrado"}`),
	}
)

// ParamsInvalidos Tratamento de erro
func ParamsInvalidos(err error) error {
	return err
}

// BancoDados Tratamento de erro
func BancoDados(err error) error {
	if strings.Contains(err.Error(), "pq") {
		err = fmt.Errorf(`{"postgres": "%v"}`, strings.Split(err.Error(), ":")[1][1:])
	}
	return err
}
