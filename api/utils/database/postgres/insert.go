package postgres

import (
	// _ Importanto apenas o init
	_ "github.com/lib/pq"
)

// Insert Insert Universal para o Banco Postgres
func Insert(query string, params ...interface{}) (map[string]interface{}, error) {
	var (
		values map[string]interface{}
	)

	stmt, err := postgres.Prepare(query)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(params...)
	if err != nil {
		return nil, err
	}

	columns, _ := rows.Columns()

	for rows.Next() {

		var (
			rowsValues = make(map[string]interface{}, len(columns))
			refs       = make([]interface{}, 0, len(columns))
		)

		for _, column := range columns {
			var ref interface{}
			rowsValues[column] = &ref
			refs = append(refs, &ref)
		}

		rows.Scan(refs...)

		values = rowsValues

	}

	return values, nil
}
