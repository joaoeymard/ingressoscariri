package usuario

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/JoaoEymard/ingressoscariri/api/utils/logger"
	"github.com/JoaoEymard/ingressoscariri/api/v1/model/usuario"
)

// Insert Responsavel por Inserir um registro
func Insert(res http.ResponseWriter, req *http.Request) {

}

// FindAll Retorna os eventos via json
func FindAll(res http.ResponseWriter, req *http.Request) {

	begin := time.Now().UTC()

	jsonEventos, statusCode, err := usuario.FindAll()

	res.WriteHeader(statusCode)
	if err != nil {
		res.Write([]byte(err.Error()))
		return
	}

	res.Write(jsonEventos)

	logger.Infoln(logger.Status(fmt.Sprintf("%+v\n", res)), req.RemoteAddr, req.Method, req.URL, time.Now().UTC().Sub(begin))
}

// FindByID Retorna o evento correspondente ao id via json
func FindByID(res http.ResponseWriter, req *http.Request) {

	begin := time.Now().UTC()

	jsonEvento, statusCode, err := usuario.FindByID(mux.Vars(req)["link"])

	res.WriteHeader(statusCode)
	if err != nil {
		res.Write([]byte(err.Error()))
		return
	}

	res.Write(jsonEvento)

	logger.Infoln(logger.Status(fmt.Sprintf("%+v\n", res)), req.RemoteAddr, req.Method, req.URL, time.Now().UTC().Sub(begin))

}

// Update Responsavel por Atualizar um registro
func Update(res http.ResponseWriter, req *http.Request) {

}

// Delete Responsavel por Deletar um registro
func Delete(res http.ResponseWriter, req *http.Request) {

}
