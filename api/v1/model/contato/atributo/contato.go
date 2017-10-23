package atributo

var (
	// Filtros lista de filtro para consulta
	Filtros = map[string]string{
		"id": "CONTATO.id = %v",
		// "nome": "CONTATO.nome ILIKE '%%'|| '%v' ||'%%'",
	}
)

// ValidValues Responsavel por validar os atributos recebidos
func ValidValues(params map[string]interface{}) bool {

	var (
		valid map[string]func(interface{}) bool
	)

	for key, value := range params {
		if retorno := valid[key](value); !retorno {
			return false
		}
	}

	return true

}

func endereco(value interface{}) {

}

func complemento(value interface{}) {

}

func referencia(value interface{}) {

}

func bairro(value interface{}) {

}

func cep(value interface{}) {

}

func cidade(value interface{}) {

}

func uf(value interface{}) {

}

func telefone_principal(value interface{}) {

}

func telefone_secundario(value interface{}) {

}

func telefone_terciario(value interface{}) {

}

func email(value interface{}) {

}
