package usuario

var (
	atributos = []string{
		"id",
		"nome",
		"login",
		"senha",
		"ultimo_acesso",
		"administrador",
		"ativo",
		"cpf",
		"data_nascimento",
		"sexo",
	}

	filtros = map[string]string{
		"id":   "USR.id = %v",
		"nome": "USR.nome ILIKE '%%'|| '%v' ||'%%'",
	}
)
