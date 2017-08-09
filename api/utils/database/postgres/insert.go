package postgres

import (
	"database/sql"
	"errors"

	"strings"

	"fmt"

	// _ Importanto apenas o init
	_ "github.com/lib/pq"
)

// Insert Insert Universal para o Banco Postgres
func Insert(attributes, table string, values []interface{}) (*sql.Rows, error) {
	var query []string

	query = []string{"INSERT INTO"}

	if table != "" {
		query = append(query, table)
	} else {
		return nil, errors.New("Parametro - 'Tabela' está vazio")
	}

	if attributes != "" {
		query = append(query, "(")
		query = append(query, attributes)
		query = append(query, ")")
	} else {
		return nil, errors.New("Parametro - 'Atributos' está vazio")
	}

	query = append(query, "VALUES")

	if values != nil {
		query = append(query, "(")
		for y := range values {
			if y+1 != len(values) {
				query = append(query, fmt.Sprintf("$%d,", y))
			} else {
				query = append(query, fmt.Sprintf("$%d", y))
			}
		}
		query = append(query, ")")
	} else {
		return nil, errors.New("Parametro - 'Values' está vazio")
	}

	query = append(query, ";")

	stmt, err := postgres.Prepare(strings.Join(query, " "))
	if err != nil {
		return nil, err
	}

	return stmt.Query(values...)
}
