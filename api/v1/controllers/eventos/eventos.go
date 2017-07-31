package eventos

import (
	"fmt"

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
	jsonEventos, err := eventos.FindAll()
	if err != nil {
		ctx.JSON(map[string]interface{}{"Erro": err.Error()})
		return
	}

	ctx.JSON(jsonEventos)
}

// FindByID Retorna o evento correspondente ao id via json
func FindByID(ctx context.Context) {
	fmt.Println("id", ctx.Params().Get("id"))
	jsonEvento, err := eventos.FindByID(ctx.Params().Get("id"))
	if err != nil {
		ctx.JSON(map[string]interface{}{"Erro": err.Error()})
		return
	}

	ctx.JSON(jsonEvento)
}
