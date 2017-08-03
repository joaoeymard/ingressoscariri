package main

import (
	"net/http"
	"os"
	"runtime"
	"strconv"
	"time"

	"github.com/JoaoEymard/ingressoscariri/api"
	"github.com/JoaoEymard/ingressoscariri/api/utils/database/postgres"
	"github.com/JoaoEymard/ingressoscariri/api/utils/logger"
	"github.com/JoaoEymard/ingressoscariri/api/utils/settings"
	"github.com/gorilla/mux"
	//sql "github.com/JoaoEymard/ingressoscariri/service/core/database"
)

var (
	app *mux.Router
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	app = mux.NewRouter()

	if err := postgres.Open(); err != nil {
		if GoDetails, _ := strconv.ParseBool(os.Getenv("GO_DETAILS")); GoDetails {
			logger.Errorf("Conexão Postgres %v", err.Error())
		} else {
			logger.Errorln("Open Postgres")
		}
	}
}

func main() {

	api.Routes(app)

	logger.Infof("Inciando a aplicação, acesse: %v", "http://"+settings.GetSettings().Listen)

	srv := &http.Server{
		Handler:      app,
		Addr:         settings.GetSettings().Listen,
		ReadTimeout:  2 * time.Minute,
		WriteTimeout: 2 * time.Minute,
	}

	err := srv.ListenAndServe()
	if err != nil {
		logger.Errorln("Fechando aplicação com erro:", err.Error())
	}
}
