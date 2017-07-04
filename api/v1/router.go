package v1

import (
	"github.com/JoaoEymard/ingressoscariri/api/v1/controllers/auth"
	ctrlEventos "github.com/JoaoEymard/ingressoscariri/api/v1/controllers/eventos"
	ctrlMapa "github.com/JoaoEymard/ingressoscariri/api/v1/controllers/mapa"
	"github.com/JoaoEymard/ingressoscariri/api/v1/middleware"
	"github.com/kataras/iris/core/router"
)

// ConfigRoutes Tratamento das Rotas publicas
func ConfigRoutes(router router.Party) {

	//Login
	router.Get("/login/{user:string}/{passw:string}", auth.Login)

	//Logoff
	router.Get("/logoff", auth.Logoff)

	// Mapa
	router.Get("/map", middleware.Check, ctrlMapa.Find)

	//Eventos
	router.Get("/eventos", ctrlEventos.FindAll)
	router.Get("/eventos/simples")
	router.Get("/eventos/{id:int}")

	// CEP
	//router.Get("/cep/{cep:"+utils.Regex["integer"]+"}", ctrlCep.Find)

	//router.Get("/cidades", ctrlCidades.FindAll)

}
