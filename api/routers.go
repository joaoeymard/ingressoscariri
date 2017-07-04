package api

import (
	"github.com/JoaoEymard/ingressoscariri/api/v1"
	"github.com/kataras/iris"
)

// Routes pacotes das rotas
func Routes(app *iris.Application) {

	apiParty := app.Party("/api")
	{

		v1Party := apiParty.Party("/v1")
		v1.ConfigRoutes(v1Party)

	}

}
