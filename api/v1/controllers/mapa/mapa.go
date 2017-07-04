package mapa

import (
	"github.com/JoaoEymard/ingressoscariri/api/v1/model/mapa"
	"github.com/kataras/iris/context"
)

// Find Retorna o mapa via json
func Find(ctx context.Context) {

	jsonMap, err := mapa.Find()
	if err != nil {
		ctx.JSON(map[string]interface{}{"Erro": err.Error()})
		return
	}

	ctx.JSON(jsonMap)

}
