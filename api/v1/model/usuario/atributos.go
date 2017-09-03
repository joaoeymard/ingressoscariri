package usuario

var (
	filtros = map[string]string{
		"id":   "USR.id = %v",
		"nome": "USR.nome ILIKE '%%'|| '%v' ||'%%'",
	}
)

func validParams(params map[string]interface{}) bool {

	if value, valueOf := params["nome"].(string); value == "" && !valueOf {
		return false
	}

	if value, valueOf := params["senha"].(string); value == "" && !valueOf {
		return false
	}

	if _, valueOf := params["ativo"].(bool); !valueOf {
		return false
	}

	if value, valueOf := params["cpf"].(string); value == "" && !valueOf {
		return false
	}

	if value, valueOf := params["data_nascimento"].(string); value == "" && !valueOf {
		return false
	}

	if value, valueOf := params["sexo"].(string); value == "" && !valueOf {
		return false
	}

	if _, valueOf := params["nivel"].(float64); !valueOf {
		return false
	}

	return true

}

func validParamsContato(params map[string]interface{}) bool {

	contato, valueOf := params["contato"].(map[string]interface{})

	if contato == nil && !valueOf {
		return false
	}

	if value, valueOf := contato["endereco"].(string); value == "" && !valueOf {
		return false
	}
	if value, valueOf := contato["complemento"].(string); value == "" && !valueOf {
		return false
	}
	if value, valueOf := contato["referencia"].(string); value == "" && !valueOf {
		return false
	}
	if value, valueOf := contato["bairro"].(string); value == "" && !valueOf {
		return false
	}
	if value, valueOf := contato["cep"].(string); value == "" && !valueOf {
		return false
	}
	if value, valueOf := contato["cidade"].(string); value == "" && !valueOf {
		return false
	}
	if value, valueOf := contato["uf"].(string); value == "" && !valueOf {
		return false
	}
	if value, valueOf := contato["telefone_principal"].(string); value == "" && !valueOf {
		return false
	}
	if value, valueOf := contato["telefone_secundario"].(string); value == "" && !valueOf {
		return false
	}
	if value, valueOf := contato["telefone_terciario"].(string); value == "" && !valueOf {
		return false
	}
	if value, valueOf := contato["email"].(string); value == "" && !valueOf {
		return false
	}

	return true

}
