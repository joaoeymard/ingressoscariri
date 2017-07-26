package v1

import (
	ctrlAuth "github.com/JoaoEymard/ingressoscariri/api/v1/controllers/auth"
	ctrlEventos "github.com/JoaoEymard/ingressoscariri/api/v1/controllers/eventos"
	ctrlMapa "github.com/JoaoEymard/ingressoscariri/api/v1/controllers/mapa"
	// ctrlMiddleware "github.com/JoaoEymard/ingressoscariri/api/v1/middleware"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/core/router"
)

// ConfigRoutes Tratamento das Rotas publicas
func ConfigRoutes(route router.Party) {

	//Login
	route.Get("/login/{user:string}/{passw:string}", ctrlAuth.Login)

	//Logoff
	route.Get("/logout", ctrlAuth.Logout)

	// Mapa
	route.Get("/map", ctrlAuth.Check, ctrlMapa.Find)

	//Eventos
	route.Post("/eventos", ctrlEventos.Insert)
	route.Get("/eventos/", ctrlEventos.FindAll)
	route.Get("/evento/{id:int min(1)}", ctrlEventos.FindByID)
	// route.Get("/eventos/simples")
	// route.Get("/eventos/{id:int}")

	route.Get("/", func(ctx context.Context) {
		ctx.JSON(map[string]string{"api": "testeAPI"})
	})

	// CEP
	//route.Get("/cep/{cep:"+utils.Regex["integer"]+"}", ctrlCep.Find)

	//route.Get("/cidades", ctrlCidades.FindAll)

}
