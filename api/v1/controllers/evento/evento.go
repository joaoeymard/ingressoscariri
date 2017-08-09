package evento

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/JoaoEymard/ingressoscariri/api/utils/logger"
	"github.com/JoaoEymard/ingressoscariri/api/v1/model/evento"
)

// Insert Responsavel por Inserir o registro
func Insert(res http.ResponseWriter, req *http.Request) {

	var c struct {
		Evento string
	}

	// if err := ctx.ReadJSON(&c); err != nil {
	// ctx.StatusCode(iris.StatusBadRequest)
	// ctx.WriteString(err.Error())
	// 	return
	// }

	fmt.Printf("Received: %#v\n", c)

	// ctx.Writef("Received: %#v\n", c)

	// jsonEventos, statusCode, err := eventos.Insert()
	// ctx.StatusCode(statusCode)

	// if err != nil {
	// 	ctx.JSON(err.Error())
	// 	return
	// }

	// ctx.JSON(jsonEventos)
	// ctx.JSON("1")
}

// FindAll Retorna os eventos via json
func FindAll(res http.ResponseWriter, req *http.Request) {

	begin := time.Now().UTC()

	jsonEventos, statusCode, err := evento.FindAll()

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

	jsonEvento, statusCode, err := evento.FindByID(mux.Vars(req)["link"])

	res.WriteHeader(statusCode)
	if err != nil {
		res.Write([]byte(err.Error()))
		return
	}

	res.Write(jsonEvento)

	logger.Infoln(logger.Status(fmt.Sprintf("%+v\n", res)), req.RemoteAddr, req.Method, req.URL, time.Now().UTC().Sub(begin))

}
