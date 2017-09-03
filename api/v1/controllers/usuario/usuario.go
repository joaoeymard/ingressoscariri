package usuario

import (
	"fmt"
	"net/http"
	"time"

	"github.com/JoaoEymard/ingressoscariri/api/utils/logger"
	"github.com/JoaoEymard/ingressoscariri/api/v1/model/usuario"
)

// Insert Responsavel por Inserir um registro
func Insert(res http.ResponseWriter, req *http.Request) {

	begin := time.Now().UTC()

	jsonEventos, statusCode, err := usuario.Insert(req.Body)

	res.WriteHeader(statusCode)
	if err != nil {
		res.Write([]byte(err.Error()))
		logger.Warnln(logger.Status(fmt.Sprintf("%+v\n", res)), req.RemoteAddr, req.Method, req.URL, time.Now().UTC().Sub(begin))
		return
	}

	res.Write(jsonEventos)

	logger.Infoln(logger.Status(fmt.Sprintf("%+v\n", res)), req.RemoteAddr, req.Method, req.URL, time.Now().UTC().Sub(begin))

}

// Find Retorna o(s) registro(s) via json
func Find(res http.ResponseWriter, req *http.Request) {

	begin := time.Now().UTC()

	jsonEventos, statusCode, err := usuario.Find(req.URL.Query())

	res.WriteHeader(statusCode)
	if err != nil {
		res.Write([]byte(err.Error()))
		logger.Warnln(logger.Status(fmt.Sprintf("%+v\n", res)), req.RemoteAddr, req.Method, req.URL, time.Now().UTC().Sub(begin))
		return
	}

	res.Write(jsonEventos)

	logger.Infoln(logger.Status(fmt.Sprintf("%+v\n", res)), req.RemoteAddr, req.Method, req.URL, time.Now().UTC().Sub(begin))
}

// Update Responsavel por Atualizar um registro
func Update(res http.ResponseWriter, req *http.Request) {

}

// Delete Responsavel por Deletar um registro
func Delete(res http.ResponseWriter, req *http.Request) {

}
