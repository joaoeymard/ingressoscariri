package postgres

import (
	"database/sql"
	"errors"

	"strings"

	_ "github.com/lib/pq"
)

// Select Select Universal para o Banco Postgres
func Select(attributes []string, table string, join, where, order []string, limit string) (*sql.Rows, error) {

	var query []string

	query = []string{"SELECT"}

	if attributes != nil {
		query = append(query, strings.Join(attributes, ","))
	} else {
		query = append(query, "*")
	}

	query = append(query, "FROM")

	if table != "" {
		query = append(query, table)
	} else {
		return nil, errors.New("Parametro - 'Tabela' está vazio")
	}

	if join != nil {
		query = append(query, strings.Join(join, " "))
	}

	if where != nil {
		query = append(query, "WHERE")
		query = append(query, strings.Join(where, " "))
	}

	if order != nil {
		query = append(query, "ORDER BY")
		query = append(query, strings.Join(order, " "))
	}

	if limit != "" {
		query = append(query, "LIMIT")
		query = append(query, limit)
	}

	query = append(query, ";")

	stmt, err := postgres.Prepare(strings.Join(query, " "))
	if err != nil {
		return nil, err
	}

	return stmt.Query()
}
