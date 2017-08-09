package usuario

import (
	"errors"
	"net/http"

	"github.com/JoaoEymard/ingressoscariri/api/utils/database/postgres"
)

// Insert Adiciona um registro
func Insert() ([]map[string]interface{}, int, error) {

	return nil, 0, nil

}

// FindAll Retorna os eventos via json
func FindAll() ([]byte, int, error) {

	attributes := `usuario.id AS id_usuario, usuario.nome, usuario.login, usuario.senha, usuario.ultimo_acesso, usuario.administrador, usuario.ativo, usuario.cpf, usuario.data_nascimento, usuario.sexo,
				   contato.id, contato.endereco, contato.complemento, contato.referencia, contato.bairro, contato.cep, contato.cidade, contato.uf, contato.telefone_principal, contato.telefone_secundario, contato.telefone_terciario, contato.email`

	table := "t_ingressoscariri_usuario AS usuario"

	join := "LEFT JOIN t_ingressoscariri_usuario_contato AS contato ON usuario.id = contato.id_usuario"

	order := "usuario.nome ASC"

	rows, err := postgres.Select(attributes, table, join, "", order, "")
	if err != nil {
		return nil, http.StatusBadRequest, errors.New("Bad Request: " + err.Error())
	}

	for rows.Next() {
		var (
			idUsuario                                                                                                                                                                                                string
			usuarioNome, usuarioLogin, usuarioSenha, usuarioUltimoAcesso, usuarioAdministrador, usuarioAtivo, usuarioCpf, usuarioDataNascimento, usuarioSexo                                                         interface{}
			idContato                                                                                                                                                                                                string
			contatoEndereco, contatoComplemento, contatoReferencia, contatoBairro, contatoCep, contatoCidade, contatoUf, contatoTelefonePrincipal, contatoTelefoneSecundario, contatoTelefoneTerciario, contatoEmail interface{}
		)

		rows.Scan(&idUsuario, &usuarioNome, &usuarioLogin, &usuarioSenha, &usuarioUltimoAcesso, &usuarioAdministrador, &usuarioAtivo, &usuarioCpf, &usuarioDataNascimento, &usuarioSexo, &idContato, &contatoEndereco, &contatoComplemento, &contatoReferencia, &contatoBairro, &contatoCep, &contatoCidade, &contatoUf, &contatoTelefonePrincipal, &contatoTelefoneSecundario, &contatoTelefoneTerciario, &contatoEmail)
	}

	return nil, 0, nil

}

// FindByID Retorna o evento correspondente ao id via json
func FindByID(link string) ([]byte, int, error) {

	return nil, 0, nil

}

// Update Adiciona um registro
func Update() ([]map[string]interface{}, int, error) {

	return nil, 0, nil

}

// Delete Adiciona um registro
func Delete() ([]map[string]interface{}, int, error) {

	return nil, 0, nil

}
