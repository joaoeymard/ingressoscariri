package eventos

import (
	"github.com/JoaoEymard/ingressoscariri/api/utils/database/postgres"
)

// FindAll Retorna os eventos via json
func FindAll() ([]map[string]interface{}, error) {

	var (
		eventos     = make(map[string]interface{})
		periodos    = make(map[string]interface{})
		jsonEventos []map[string]interface{}
	)

	attributes := []string{
		"evento.id AS id_evento", "evento.titulo", "evento.imagem", "evento.cidade", "evento.uf", "evento.localidade",
		"periodo.id AS id_periodo", "periodo.data_periodo",
	}

	table := "t_ingressoscariri_evento AS evento"

	join := []string{
		"LEFT JOIN t_ingressoscariri_evento_periodo AS periodo ON evento.id = periodo.id_evento",
	}

	rows, err := postgres.Select(attributes, table, join, nil, nil, "")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			idEvento, idPeriodo                    string
			titulo, imagem, cidade, uf, localidade interface{}
			dataPeriodo                            interface{}
		)

		rows.Scan(&idEvento, &titulo, &imagem, &cidade, &uf, &localidade, &idPeriodo, &dataPeriodo)

		eventos[idEvento] = map[string]interface{}{
			"id":         idEvento,
			"titulo":     titulo,
			"imagem":     imagem,
			"cidade":     cidade,
			"uf":         uf,
			"localidade": localidade,
		}

		if idPeriodo != "" {
			periodos[idPeriodo] = map[string]interface{}{
				"idEvento":     idEvento,
				"id":           idPeriodo,
				"data_periodo": dataPeriodo,
			}
		}

	}

	var (
		eventoAux  []map[string]interface{}
		periodoAux []map[string]interface{}
	)

	for idEvento, evento := range eventos {
		for _, periodo := range periodos {
			if periodo.(map[string]interface{})["idEvento"] == idEvento {
				periodoAux = append(periodoAux, map[string]interface{}{
					"data_periodo": periodo.(map[string]interface{})["data_periodo"],
				})
			}
		}

		eventoAux = append(eventoAux, map[string]interface{}{
			"id":         evento.(map[string]interface{})["id"],
			"titulo":     evento.(map[string]interface{})["titulo"],
			"imagem":     evento.(map[string]interface{})["imagem"],
			"cidade":     evento.(map[string]interface{})["cidade"],
			"uf":         evento.(map[string]interface{})["uf"],
			"localidade": evento.(map[string]interface{})["localidade"],
			"periodo":    periodoAux,
		})

		periodoAux = nil
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
		"data": eventoAux,
	})

	return jsonEventos, nil

}

// FindByID Retorna o evento correspondente ao id via json
func FindByID(id string) (map[string]interface{}, error) {

	var (
		eventos    = make(map[string]interface{})
		periodos   = make(map[string]interface{})
		categorias = make(map[string]interface{})
		galerias   = make(map[string]interface{})
		jsonEvento map[string]interface{}
	)

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

	where := []string{
		"evento.id = " + id,
	}

	order := []string{
		"periodo.data_periodo ASC",
	}

	rows, err := postgres.Select(attributes, table, join, where, order, "")
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

	var (
		eventoAux    map[string]interface{}
		periodoAux   []map[string]interface{}
		categoriaAux []map[string]interface{}
		galeriaAux   []map[string]interface{}
	)

	for idEvento, evento := range eventos {
		periodoAux = nil
		for idPeriodo, periodo := range periodos {
			categoriaAux = nil
			if periodo.(map[string]interface{})["idEvento"] == idEvento {
				for _, categoria := range categorias {
					if categoria.(map[string]interface{})["idPeriodo"] == idPeriodo {
						categoriaAux = append(categoriaAux, map[string]interface{}{
							"nome":               categoria.(map[string]interface{})["nome"],
							"valor":              categoria.(map[string]interface{})["valor"],
							"quantidade":         categoria.(map[string]interface{})["quantidade"],
							"quantidade_vendida": categoria.(map[string]interface{})["quantidade_vendida"],
							"lote":               categoria.(map[string]interface{})["lote"],
						})
					}
				}
				periodoAux = append(periodoAux, map[string]interface{}{
					"atracao":   periodo.(map[string]interface{})["atracao"],
					"data":      periodo.(map[string]interface{})["data_periodo"],
					"categoria": categoriaAux,
				})
			}
		}

		galeriaAux = nil
		for _, galeria := range galerias {
			if galeria.(map[string]interface{})["idEvento"] == idEvento {
				// fmt.Println("Galeria", galeria)
				galeriaAux = append(galeriaAux, map[string]interface{}{
					"imagem": galeria.(map[string]interface{})["imagem"],
				})
			}
		}

		eventoAux = map[string]interface{}{
			"id":           evento.(map[string]interface{})["id"],
			"titulo":       evento.(map[string]interface{})["titulo"],
			"imagem":       evento.(map[string]interface{})["imagem"],
			"cidade":       evento.(map[string]interface{})["cidade"],
			"estado":       evento.(map[string]interface{})["uf"],
			"data_criacao": evento.(map[string]interface{})["localidade"],
			"local":        evento.(map[string]interface{})["localidade"],
			"taxa":         evento.(map[string]interface{})["taxa"],
			"mapa":         evento.(map[string]interface{})["mapa"],
			"descricao":    evento.(map[string]interface{})["descricao"],
			"periodo":      periodoAux,
			"galeria":      galeriaAux,
		}
	}

	jsonEvento = eventoAux

	return jsonEvento, nil

}

// Insert Retorna os eventos via json
func Insert() ([]map[string]interface{}, error) {

	attributes := []string{
		"titulo", "imagem", "cidade", "uf", "localidade", "taxa", "mapa", "descricao",
	}

	table := "t_ingressoscariri_evento"

	values := [][]interface{}{
		[]interface{}{"titulo1", "hash(imagem1)", "cidade1", "u1", "localidade1", 19.0, "mapa1", "descricao1"},
		[]interface{}{"titulo2", "hash(imagem2)", "cidade2", "u2", "localidade2", 19.0, "mapa2", "descricao2"},
	}

	_, err := postgres.Insert(attributes, table, values)
	if err != nil {
		return nil, err
	}

	return nil, nil

}
