package eventos

import (
	"github.com/JoaoEymard/ingressoscariri/api/v1/model/eventos"
	"github.com/kataras/iris/context"
)

// Insert Responsavel por Inserir o registro
func Insert(ctx context.Context) {

	// jsonEventos, statusCode, err := eventos.Insert()
	// ctx.StatusCode(statusCode)

	// if err != nil {
	// 	ctx.JSON(err.Error())
	// 	return
	// }

	// ctx.JSON(jsonEventos)
	ctx.JSON("1")
}

// FindAll Retorna os eventos via json
func FindAll(ctx context.Context) {

	jsonEventos, statusCode, err := eventos.FindAll()
	ctx.StatusCode(statusCode)

	if err != nil {
		ctx.JSON(err.Error())
		return
	}

	ctx.JSON(jsonEventos)
}

// FindByID Retorna o evento correspondente ao id via json
func FindByID(ctx context.Context) {

	jsonEvento, statusCode, err := eventos.FindByID(ctx.Params().Get("link"))
	ctx.StatusCode(statusCode)

	if err != nil {
		ctx.JSON(err.Error())
		return
	}

	ctx.JSON(jsonEvento)

}
