package eventos

import (
	"fmt"

	"github.com/JoaoEymard/ingressoscariri/api/utils/database/postgres"
)

// FindAll Retorna os eventos via json
func FindAll() ([]map[string]interface{}, error) {

	var (
		eventos     = make(map[string]interface{})
		periodos    = make(map[string]interface{})
		categorias  = make(map[string]interface{})
		galerias    = make(map[string]interface{})
		jsonEventos []map[string]interface{}
	)

	if err := postgres.Open(); err != nil {
		return nil, err
	}
	defer postgres.Close()

	attributes := []string{
		"evento.id AS id_evento", "evento.titulo", "evento.imagem", "evento.cidade", "evento.uf", "evento.localidade", "evento.taxa", "evento.mapa", "evento.descricao",
		"periodo.id AS id_periodo", "periodo.data_periodo", "periodo.atracao", "periodo.lote",
		"categoria.id AS id_categoria", "categoria.nome", "categoria.valor", "categoria.quantidade", "categoria.quantidade_vendida",
		"galeria.id AS id_galeria", "galeria.imagem",
	}

	table := "t_ingressoscariri_evento AS evento"

	join := []string{
		"LEFT JOIN t_ingressoscariri_evento_periodo AS periodo ON evento.id = periodo.id_evento",
		"LEFT JOIN t_ingressoscariri_evento_categoria AS categoria ON periodo.id = categoria.id_periodo",
		"LEFT JOIN t_ingressoscariri_evento_galeria AS galeria ON evento.id = galeria.id_evento",
	}

	rows, err := postgres.Select(attributes, table, join, nil, nil, "")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			idEvento, idPeriodo, idCategoria, idGaleria                   string
			titulo, imagem, cidade, uf, localidade, taxa, mapa, descricao interface{}
			dataPeriodo, atracao, lote                                    interface{}
			nomeCategoria, valor, quantidade, quantidadeVendida           interface{}
			imagemGaleria                                                 interface{}
		)

		rows.Scan(&idEvento, &titulo, &imagem, &cidade, &uf, &localidade, &taxa, &mapa, &descricao, &idPeriodo, &dataPeriodo, &atracao, &lote, &idCategoria, &nomeCategoria, &valor, &quantidade, &quantidadeVendida, &idGaleria, &imagemGaleria)

		eventos[idEvento] = map[string]interface{}{
			"id":         idEvento,
			"titulo":     titulo,
			"imagem":     imagem,
			"cidade":     cidade,
			"uf":         uf,
			"localidade": localidade,
			"taxa":       taxa,
			"mapa":       mapa,
			"descricao":  descricao,
		}

		if idPeriodo != "" {
			periodos[idPeriodo] = map[string]interface{}{
				"idEvento":     idEvento,
				"id":           idPeriodo,
				"data_periodo": dataPeriodo,
				"atracao":      atracao,
			}
		}

		if idCategoria != "" {
			categorias[idCategoria] = map[string]interface{}{
				"idPeriodo":          idPeriodo,
				"id":                 idCategoria,
				"nome":               nomeCategoria,
				"valor":              valor,
				"quantidade":         quantidade,
				"quantidade_vendida": quantidadeVendida,
				"lote":               lote,
			}
		}

		if idGaleria != "" {
			galerias[idGaleria] = map[string]interface{}{
				"idEvento": idEvento,
				"id":       idGaleria,
				"imagem":   imagemGaleria,
			}
		}

	}

	for idEvento, evento := range eventos {
		fmt.Println("Eventos", idEvento, evento)
		for idPeriodo, periodo := range periodos {
			if periodo.(map[string]interface{})["idEvento"] == idEvento {
				fmt.Println("Periodo", periodo)
				for _, categoria := range categorias {
					if categoria.(map[string]interface{})["idPeriodo"] == idPeriodo {
						fmt.Println("Categoria", categoria)
					}
				}
			}
		}
		for _, galeria := range galerias {
			if galeria.(map[string]interface{})["idEvento"] == idEvento {
				fmt.Println("Galeria", galeria)
			}
		}
	}

	attributes = []string{"COUNT(*)"}
	table = "t_ingressoscariri_evento"

	rows, err = postgres.Select(attributes, table, nil, nil, nil, "")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var count interface{}

		rows.Scan(&count)

		jsonEventos = append(jsonEventos, map[string]interface{}{
			"total": count,
		})

	}

	jsonEventos = append(jsonEventos, map[string]interface{}{
		"data": jsonEventos,
	})

	return jsonEventos, nil

}
