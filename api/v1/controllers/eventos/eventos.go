package eventos

import (
	"github.com/JoaoEymard/ingressoscariri/api/v1/model/eventos"
	"github.com/kataras/iris/context"
)

func FindAll(ctx context.Context) {
	jsonEventos, err := eventos.FindAll()
	if err != nil {
		ctx.JSON(map[string]interface{}{"Erro": err.Error()})
		return
	}

	ctx.JSON(jsonEventos)
}
