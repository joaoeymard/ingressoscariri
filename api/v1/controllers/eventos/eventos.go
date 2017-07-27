package eventos

import (
	"github.com/JoaoEymard/ingressoscariri/api/v1/model/eventos"
	"github.com/kataras/iris/context"
)

// Insert Responsavel por Inserir o registro
func Insert(ctx context.Context) {

	jsonEventos, err := eventos.Insert()
	if err != nil {
		ctx.JSON(map[string]interface{}{"Erro": err.Error()})
		return
	}

	ctx.JSON(jsonEventos)
}

// FindAll Retorna os eventos via json
func FindAll(ctx context.Context) {

	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Access-Control-Allow-Methods", "GET")
	jsonEventos, err := eventos.FindAll()
	if err != nil {
		ctx.JSON(map[string]interface{}{"Erro": err.Error()})
		return
	}

	ctx.JSON(jsonEventos)
}

// FindByID Retorna o evento correspondente ao id via json
func FindByID(ctx context.Context) {

	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Access-Control-Allow-Methods", "GET")
	jsonEvento, err := eventos.FindByID(ctx.Params().Get("id"))
	if err != nil {
		ctx.JSON(map[string]interface{}{"Erro": err.Error()})
		return
	}

	ctx.JSON(jsonEvento)

}
