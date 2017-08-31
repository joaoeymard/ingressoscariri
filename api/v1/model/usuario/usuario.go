package usuario

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/JoaoEymard/ingressoscariri/api/utils/database/postgres"
	"github.com/JoaoEymard/ingressoscariri/api/v1/utils"
)

const (
	// TableUsuario Tabela referente ao usuario
	TableUsuario = "t_ingressoscariri_usuario USR"
)

// Insert Adiciona um registro
func Insert() ([]map[string]interface{}, int, error) {

	return nil, 0, nil

}

// Find Retorna os eventos via json
func Find(params url.Values) ([]byte, int, error) {

	// Consulta para saber o total de registro
	sqlTotal := fmt.Sprintf(`SELECT COUNT(*) AS total
	FROM %v
	LEFT JOIN t_ingressoscariri_usuario_contato AS CONTATO ON USR.id = CONTATO.id_usuario`, TableUsuario)

	// Retorna um []map com as colunas e valores vindo do banco de dados
	rowsTotal, err := postgres.Select(sqlTotal)
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
	sql := fmt.Sprintf(`SELECT USR.id AS id, USR.nome, USR.login, USR.senha, USR.ultimo_acesso, USR.administrador, USR.ativo, USR.cpf, USR.data_nascimento, USR.sexo,
	CONTATO.id AS contato_id, CONTATO.endereco, CONTATO.complemento, CONTATO.referencia, CONTATO.bairro, CONTATO.cep, CONTATO.cidade, CONTATO.uf, CONTATO.telefone_principal, CONTATO.telefone_secundario, CONTATO.telefone_terciario, CONTATO.email
	FROM %v
	LEFT JOIN t_ingressoscariri_usuario_contato AS CONTATO ON USR.id = CONTATO.id_usuario
	%v %v %v %v`, TableUsuario, filter, order, limit, offset)

	// Retorna um []map com as colunas e valores vindo do banco de dados
	rows, err := postgres.Select(sql)
	if err != nil {
		return nil, http.StatusBadRequest, utils.BancoDados(err)
	}

	// Verifica se o retorno está nulo
	if rows == nil {
		return nil, http.StatusNotFound, utils.Errors["NOT_FOUND"]
	}

	// Monta a estrutura de retorno
	dados := map[string]interface{}{
		"dados": rows,
		"total": rowsTotal[0]["total"],
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
