package evento

import (
	"net/http"

	"github.com/JoaoEymard/ingressoscariri/api/utils/database/postgres"
)

// Insert Retorna os eventos via json
func Insert() ([]map[string]interface{}, int, error) {

	attributes := "titulo, imagem, cidade, uf, localidade, taxa, mapa, descricao"

	table := "t_ingressoscariri_evento"

	values := []interface{}{"titulo1", "hash(imagem1)", "cidade1", "u1", "localidade1", 19.0, "mapa1", "descricao1"}

	_, err := postgres.Insert(attributes, table, values)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	return nil, http.StatusOK, nil

}

// FindAll Retorna os eventos via json
func FindAll() ([]byte, int, error) {

	// 	var (
	// 		eventos     []map[string]interface{}
	// 		periodos    []map[string]interface{}
	// 		jsonEventos []map[string]interface{}
	// 	)

	// 	attributes := `DISTINCT ON (evento.id) evento.id AS id_evento, evento.titulo, evento.imagem, evento.cidade, evento.uf, evento.localidade, evento.link,
	// 				   periodo.id, periodo.data_periodo`

	// 	table := "t_ingressoscariri_evento AS evento"

	// 	join := "LEFT JOIN t_ingressoscariri_evento_periodo AS periodo ON evento.id = periodo.id_evento"

	// 	where := "periodo.data_periodo >= NOW()"

	// 	order := "evento.id ASC, periodo.data_periodo ASC"

	// 	rows, err := postgres.Select(attributes, table, join, where, order, "")
	// 	if err != nil {
	// 		return nil, http.StatusBadRequest, errors.New("Bad Request: " + err.Error())
	// 	}

	// 	for rows.Next() {
	// 		var (
	// 			idEvento                                                                         string
	// 			eventoTitulo, eventoImagem, eventoCidade, eventoUf, eventoLocalidade, eventoLink interface{}
	// 			idPeriodo                                                                        string
	// 			periodoDataPeriodo                                                               interface{}
	// 		)

	// 		rows.Scan(&idEvento, &eventoTitulo, &eventoImagem, &eventoCidade, &eventoUf, &eventoLocalidade, &eventoLink, &idPeriodo, &periodoDataPeriodo)

	// 		eventos = append(eventos, map[string]interface{}{
	// 			"id":         idEvento,
	// 			"titulo":     eventoTitulo,
	// 			"imagem":     eventoImagem,
	// 			"cidade":     eventoCidade,
	// 			"uf":         eventoUf,
	// 			"localidade": eventoLocalidade,
	// 			"link":       eventoLink,
	// 		})

	// 		if idPeriodo != "" {
	// 			periodos = append(periodos, map[string]interface{}{
	// 				"idEvento":     idEvento,
	// 				"id":           idPeriodo,
	// 				"data_periodo": periodoDataPeriodo,
	// 			})
	// 		}

	// 	}

	// 	var (
	// 		periodoAux interface{}
	// 	)

	// 	for _, evento := range eventos {
	// 		periodoAux = nil
	// 		for _, periodo := range periodos {
	// 			if periodo["idEvento"] == evento["id"] {
	// 				periodoAux = periodo["data_periodo"]
	// 				break
	// 			}
	// 		}

	// 		jsonEventos = append(jsonEventos, map[string]interface{}{
	// 			"titulo": evento["titulo"],
	// 			"imagem": evento["imagem"],
	// 			"cidade": evento["cidade"],
	// 			"uf":     evento["uf"],
	// 			"local":  evento["localidade"],
	// 			"link":   evento["link"],
	// 			"data":   periodoAux,
	// 		})
	// 	}

	// 	if jsonEventos == nil {
	// 		return nil, http.StatusNotFound, errors.New("Not Found: Id não encontrado")
	// 	}

	// 	byteEvento, err := json.Marshal(jsonEventos)
	// 	if err != nil {
	// 		return nil, http.StatusBadRequest, errors.New("Not Found: Id não encontrado")
	// 	}

	return nil, http.StatusOK, nil

}

// FindByID Retorna o evento correspondente ao id via json
func FindByID(link string) ([]byte, int, error) {

	// 	var (
	// 		eventos    = make(map[string]interface{})
	// 		periodos   = make(map[string]interface{})
	// 		categorias = make(map[string]interface{})
	// 		galerias   = make(map[string]interface{})
	// 		jsonEvento map[string]interface{}
	// 	)

	// 	attributes := `evento.id AS id_evento, evento.titulo, evento.imagem, evento.cidade, evento.uf, evento.localidade, evento.taxa, evento.mapa, evento.descricao, evento.link,
	// 				   periodo.id AS id_periodo, periodo.data_periodo, periodo.atracao,
	// 				   categoria.id, categoria.nome, categoria.valor, categoria.quantidade, categoria.quantidade_vendida, categoria.lote,
	// 				   galeria.id, galeria.imagem`

	// 	table := "t_ingressoscariri_evento AS evento"

	// 	join := `LEFT JOIN t_ingressoscariri_evento_periodo AS periodo ON evento.id = periodo.id_evento
	// 		 	 LEFT JOIN t_ingressoscariri_periodo_categoria AS categoria ON periodo.id = categoria.id_periodo
	// 		 	 LEFT JOIN t_ingressoscariri_evento_galeria AS galeria ON evento.id = galeria.id_evento`

	// 	where := "evento.link LIKE '" + link + "'"

	// 	order := "periodo.data_periodo ASC"

	// 	rows, err := postgres.Select(attributes, table, join, where, order, "")
	// 	if err != nil {
	// 		return nil, http.StatusBadRequest, errors.New("Bad Request: " + err.Error())
	// 	}

	// 	for rows.Next() {
	// 		var (
	// 			idEvento                                                                                                                  string
	// 			eventoTitulo, eventoImagem, eventoCidade, eventoUf, eventoLocalidade, eventoTaxa, eventoMapa, eventoDescricao, eventoLink interface{}
	// 			idPeriodo                                                                                                                 string
	// 			periodoDataPeriodo, periodoAtracao                                                                                        interface{}
	// 			idCategoria                                                                                                               string
	// 			categoriaNome, categoriaValor, categoriaQuantidade, categoriaQuantidadeVendida, categoriaLote                             interface{}
	// 			idGaleria                                                                                                                 string
	// 			galeriaImagem                                                                                                             interface{}
	// 		)

	// 		rows.Scan(&idEvento, &eventoTitulo, &eventoImagem, &eventoCidade, &eventoUf, &eventoLocalidade, &eventoTaxa, &eventoMapa, &eventoDescricao, &eventoLink, &idPeriodo, &periodoDataPeriodo, &periodoAtracao, &idCategoria, &categoriaNome, &categoriaValor, &categoriaQuantidade, &categoriaQuantidadeVendida, &categoriaLote, &idGaleria, &galeriaImagem)

	// 		eventos[idEvento] = map[string]interface{}{
	// 			"id":         idEvento,
	// 			"titulo":     eventoTitulo,
	// 			"imagem":     eventoImagem,
	// 			"cidade":     eventoCidade,
	// 			"uf":         eventoUf,
	// 			"localidade": eventoLocalidade,
	// 			"taxa":       eventoTaxa,
	// 			"mapa":       eventoMapa,
	// 			"descricao":  eventoDescricao,
	// 			"link":       eventoLink,
	// 		}

	// 		if idPeriodo != "" {
	// 			periodos[idPeriodo] = map[string]interface{}{
	// 				"idEvento":     idEvento,
	// 				"id":           idPeriodo,
	// 				"data_periodo": periodoDataPeriodo,
	// 				"atracao":      periodoAtracao,
	// 			}
	// 		}

	// 		if idCategoria != "" {
	// 			categorias[idCategoria] = map[string]interface{}{
	// 				"idPeriodo":          idPeriodo,
	// 				"id":                 idCategoria,
	// 				"nome":               categoriaNome,
	// 				"valor":              categoriaValor,
	// 				"quantidade":         categoriaQuantidade,
	// 				"quantidade_vendida": categoriaQuantidadeVendida,
	// 				"lote":               categoriaLote,
	// 			}
	// 		}

	// 		if idGaleria != "" {
	// 			galerias[idGaleria] = map[string]interface{}{
	// 				"idEvento": idEvento,
	// 				"id":       idGaleria,
	// 				"imagem":   galeriaImagem,
	// 			}
	// 		}

	// 	}

	// 	var (
	// 		eventoAux    map[string]interface{}
	// 		periodoAux   []map[string]interface{}
	// 		categoriaAux []map[string]interface{}
	// 		galeriaAux   []map[string]interface{}
	// 	)

	// 	for idEvento, evento := range eventos {
	// 		periodoAux = nil
	// 		for idPeriodo, periodo := range periodos {
	// 			categoriaAux = nil
	// 			if periodo.(map[string]interface{})["idEvento"] == idEvento {
	// 				for _, categoria := range categorias {
	// 					if categoria.(map[string]interface{})["idPeriodo"] == idPeriodo {
	// 						categoriaAux = append(categoriaAux, map[string]interface{}{
	// 							"nome":               categoria.(map[string]interface{})["nome"],
	// 							"valor":              categoria.(map[string]interface{})["valor"],
	// 							"quantidade":         categoria.(map[string]interface{})["quantidade"],
	// 							"quantidade_vendida": categoria.(map[string]interface{})["quantidade_vendida"],
	// 							"lote":               categoria.(map[string]interface{})["lote"],
	// 						})
	// 					}
	// 				}
	// 				periodoAux = append(periodoAux, map[string]interface{}{
	// 					"atracao":   periodo.(map[string]interface{})["atracao"],
	// 					"data":      periodo.(map[string]interface{})["data_periodo"],
	// 					"categoria": categoriaAux,
	// 				})
	// 			}
	// 		}

	// 		galeriaAux = nil
	// 		for _, galeria := range galerias {
	// 			if galeria.(map[string]interface{})["idEvento"] == idEvento {
	// 				// fmt.Println("Galeria", galeria)
	// 				galeriaAux = append(galeriaAux, map[string]interface{}{
	// 					"imagem": galeria.(map[string]interface{})["imagem"],
	// 				})
	// 			}
	// 		}

	// 		eventoAux = map[string]interface{}{
	// 			"titulo":    evento.(map[string]interface{})["titulo"],
	// 			"imagem":    evento.(map[string]interface{})["imagem"],
	// 			"cidade":    evento.(map[string]interface{})["cidade"],
	// 			"estado":    evento.(map[string]interface{})["uf"],
	// 			"local":     evento.(map[string]interface{})["localidade"],
	// 			"taxa":      evento.(map[string]interface{})["taxa"],
	// 			"mapa":      evento.(map[string]interface{})["mapa"],
	// 			"descricao": evento.(map[string]interface{})["descricao"],
	// 			"link":      evento.(map[string]interface{})["link"],
	// 			"periodo":   periodoAux,
	// 			"galeria":   galeriaAux,
	// 		}
	// 	}

	// 	jsonEvento = eventoAux

	// 	if jsonEvento == nil {
	// 		return nil, http.StatusNotFound, errors.New("Not Found: Id não encontrado")
	// 	}

	// 	byteEvento, err := json.Marshal(jsonEvento)
	// 	if err != nil {
	// 		return nil, http.StatusBadRequest, errors.New("Not Found: Id não encontrado")
	// 	}

	return nil, http.StatusOK, nil

}