package v1

import (
	ctrlEvento "github.com/JoaoEymard/ingressoscariri/api/v1/controllers/evento"
	ctrlUser "github.com/JoaoEymard/ingressoscariri/api/v1/controllers/usuario"
	"github.com/gorilla/mux"
)

// ConfigRoutes Tratamento das Rotas publicas
func ConfigRoutes(route *mux.Router) {

	// Eventos
	route.HandleFunc("/evento", ctrlEvento.FindAll).Methods("GET")
	route.HandleFunc("/evento/{link:[a-zA-Z0-9_]+?}", ctrlEvento.FindByID).Methods("GET")

	// Usuarios
	route.HandleFunc("/usuario", ctrlUser.Insert).Methods("POST")
	route.HandleFunc("/usuario", ctrlUser.FindAll).Methods("GET")
	route.HandleFunc("/usuario/{id:[0-9]}", ctrlUser.FindByID).Methods("GET")
	route.HandleFunc("/usuario/{id:[0-9]}", ctrlUser.Update).Methods("PUT")
	route.HandleFunc("/usuario/{id:[0-9]}", ctrlUser.Delete).Methods("DELETE")

}
