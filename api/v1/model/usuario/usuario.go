package usuario

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/JoaoEymard/ingressoscariri/api/utils/database/postgres"
	"github.com/JoaoEymard/ingressoscariri/api/v1/utils"
)

const (
	// Tabela referente ao usuario
	tableUsuario = "t_ingressoscariri_usuario"
	// Tabela referente ao contato
	tableContato = "t_ingressoscariri_usuario_contato"
	// Para fazer a consulta do total de registro
	countTotal = "COUNT(*) AS total"
)

// Insert Adiciona um registro
func Insert(content io.ReadCloser) ([]byte, int, error) {

	var paramsJSON map[string]interface{}

	params, err := ioutil.ReadAll(content)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	err = json.Unmarshal(params, &paramsJSON)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	if !validParams(paramsJSON) {
		return nil, http.StatusBadRequest, errors.New(`{"erro": "Parametros inválidos"}`)
	}

	sql := fmt.Sprintf(`INSERT INTO %v( nome, senha, ativo, cpf, data_nascimento, sexo, nivel )
						VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id;`, tableUsuario)

	rows, err := postgres.Insert(sql, paramsJSON["nome"], paramsJSON["senha"], paramsJSON["ativo"], paramsJSON["cpf"], paramsJSON["data_nascimento"], paramsJSON["sexo"], paramsJSON["nivel"])
	if err != nil {
		return nil, http.StatusBadRequest, utils.BancoDados(err)
	}

	if rows == nil {
		return nil, http.StatusBadRequest, errors.New(`{"erro": "Não conseguiu cadastrar o usuario"}`)
	}

	if validParamsContato(paramsJSON) {

		contato := paramsJSON["contato"].(map[string]interface{})

		sql := fmt.Sprintf(`INSERT INTO %v( id_usuario, endereco, complemento, referencia, bairro, cep, cidade, uf, telefone_principal, telefone_secundario, telefone_terciario, email )
							VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) RETURNING id;`, tableContato)

		rowsContato, err := postgres.Insert(sql, rows["id"], contato["endereco"], contato["complemento"], contato["referencia"], contato["bairro"], contato["cep"], contato["cidade"], contato["uf"], contato["telefone_principal"], contato["telefone_secundario"], contato["telefone_terciario"], contato["email"])
		if err != nil {
			return nil, http.StatusBadRequest, utils.BancoDados(err)
		}

		if rowsContato == nil {
			return nil, http.StatusBadRequest, errors.New(`{"erro": "Não conseguiu cadastrar o contato"}`)
		}
	}

	retorno, err := json.Marshal(rows)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	return retorno, http.StatusCreated, nil

}

// Find Retorna os eventos via json
func Find(params url.Values) ([]byte, int, error) {

	var dadosRows []map[string]interface{}

	// Consulta para saber o total de registro
	sqlTotal := fmt.Sprintf(`SELECT %v
	FROM %v`, countTotal, tableUsuario)

	// Retorna um []map com as colunas e valores vindo do banco de dados
	rowsTotal, err := postgres.SelectOne(sqlTotal)
	if err != nil {
		return nil, http.StatusBadRequest, utils.BancoDados(err)
	}

	// Verifica se o retorno está nulo
	if rowsTotal == nil {
		return nil, http.StatusNotFound, utils.Errors["NOT_FOUND"]
	}

	// Tratamento dos paramentros e filtro recebidos pela URL
	filter, order, limit, offset, err := postgres.SetParams(params, filtros)
	if err != nil {
		return nil, http.StatusBadRequest, utils.ParamsInvalidos(err)
	}

	// Consulta para coletar os registro
	sql := fmt.Sprintf(`SELECT USR.id AS id, USR.nome, USR.senha, USR.ultimo_acesso, USR.ativo, USR.cpf, USR.data_nascimento, USR.sexo, USR.nivel,
	CONTATO.id AS contato_id, CONTATO.endereco, CONTATO.complemento, CONTATO.referencia, CONTATO.bairro, CONTATO.cep, CONTATO.cidade, CONTATO.uf, CONTATO.telefone_principal, CONTATO.telefone_secundario, CONTATO.telefone_terciario, CONTATO.email
	FROM %v USR
	LEFT JOIN %v CONTATO ON USR.id = CONTATO.id_usuario
	%v %v %v %v`, tableUsuario, tableContato, filter, order, limit, offset)

	// Retorna um []map com as colunas e valores vindo do banco de dados
	rows, err := postgres.Select(sql)
	if err != nil {
		return nil, http.StatusBadRequest, utils.BancoDados(err)
	}

	// Verifica se o retorno está nulo
	if rows == nil {
		return nil, http.StatusNotFound, utils.Errors["NOT_FOUND"]
	}

	for _, row := range rows {
		dadosRows = append(dadosRows, map[string]interface{}{
			"id":              row["id"],
			"nome":            row["nome"],
			"senha":           row["senha"],
			"ultimo_acesso":   row["ultimo_acesso"],
			"ativo":           row["ativo"],
			"cpf":             row["cpf"],
			"data_nascimento": row["data_nascimento"],
			"sexo":            row["sexo"],
			"nivel":           row["nivel"],
			"contato": map[string]interface{}{
				"id":                  row["contato_id"],
				"endereco":            row["endereco"],
				"complemento":         row["complemento"],
				"referencia":          row["referencia"],
				"bairro":              row["bairro"],
				"cep":                 row["cep"],
				"cidade":              row["cidade"],
				"uf":                  row["uf"],
				"telefone_principal":  row["telefone_principal"],
				"telefone_secundario": row["telefone_secundario"],
				"telefone_terciario":  row["telefone_terciario"],
				"email":               row["email"],
			},
		})
	}

	// Monta a estrutura de retorno
	dados := map[string]interface{}{
		"dados": dadosRows,
		"total": rowsTotal["total"],
	}

	// Converte a estrutura para json
	retorno, err := json.Marshal(dados)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	return retorno, http.StatusOK, nil
}

// Update Adiciona um registro
func Update() ([]map[string]interface{}, int, error) {

	return nil, 0, nil

}

// Delete Adiciona um registro
func Delete() ([]map[string]interface{}, int, error) {

	return nil, 0, nil

}
