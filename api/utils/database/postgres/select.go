package postgres

import (
	"database/sql"
	"errors"

	"strings"

	// _ Importanto apenas o init
	_ "github.com/lib/pq"
)

// Select Select Universal para o Banco Postgres
func Select(attributes, table, join, where, order, limit string) (*sql.Rows, error) {

	var query []string

	query = []string{"SELECT"}

	if attributes != "" {
		query = append(query, attributes)
	} else {
		query = append(query, "*")
	}

	query = append(query, "FROM")

	if table != "" {
		query = append(query, table)
	} else {
		return nil, errors.New("Parametro - 'Tabela' está vazio")
	}

	if join != "" {
		query = append(query, join)
	}

	if where != "" {
		query = append(query, "WHERE")
		query = append(query, where)
	}

	if order != "" {
		query = append(query, "ORDER BY")
		query = append(query, order)
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
