package postgres

import (
	"database/sql"
	"errors"

	"strings"

	"fmt"

	_ "github.com/lib/pq"
)

// Insert Insert Universal para o Banco Postgres
func Insert(attributes []string, table string, values [][]interface{}) (*sql.Rows, error) {
	var query []string

	query = []string{"INSERT INTO"}

	if table != "" {
		query = append(query, table)
	} else {
		return nil, errors.New("Parametro - 'Tabela' está vazio")
	}

	if attributes != nil {
		query = append(query, "(")
		query = append(query, strings.Join(attributes, ","))
		query = append(query, ")")
	} else {
		return nil, errors.New("Parametro - 'Atributos' está vazio")
	}

	query = append(query, "VALUES")

	var valuesAux []interface{}
	if values != nil {
		for y, value := range values {
			query = append(query, "(")
			for x, v := range value {
				valuesAux = append(valuesAux, v)
				if x+1 != len(value) {
					query = append(query, fmt.Sprintf("$%d,", len(valuesAux)))
				} else {
					query = append(query, fmt.Sprintf("$%d", len(valuesAux)))
				}
			}
			query = append(query, ")")
			if y+1 != len(values) {
				query = append(query, ",")
			}
		}
	} else {
		return nil, errors.New("Parametro - 'Values' está vazio")
	}

	query = append(query, ";")

	stmt, err := postgres.Prepare(strings.Join(query, " "))
	if err != nil {
		return nil, err
	}

	return stmt.Query(valuesAux...)
}
