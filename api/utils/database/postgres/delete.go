package postgres

import (
	// _ Importanto apenas o init
	"errors"
	"fmt"

	_ "github.com/lib/pq"
)

// DeleteOne Delete Universal para o Banco Postgres
func DeleteOne(tabela string, where string) (map[string]interface{}, error) {

	var (
		dados map[string]interface{}
	)

	query := fmt.Sprintf(`DELETE FROM %v
	WHERE %v
	RETURNING id;`, tabela, where)

	rows, err := postgres.Query(query)
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

		dados = rowsValues

	}

	if dados == nil {
		return nil, errors.New(`{"erro": "Ocorreu um erro ao deletar o registro"}`)
	}

	return dados, nil
}
