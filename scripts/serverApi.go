package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"

	"github.com/JoaoEymard/ingressoscariri/api"
	"github.com/JoaoEymard/ingressoscariri/api/utils/database/postgres"
	"github.com/JoaoEymard/ingressoscariri/api/utils/settings"

	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
)

var (
	app *iris.Application
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	app = iris.New()

	if err := postgres.Open(); err != nil {
		if GoDetails, _ := strconv.ParseBool(os.Getenv("GO_DETAILS")); GoDetails {
			fmt.Println("[Erro] Conexão Postgres", err.Error())
		} else {
			fmt.Println("[Erro] Open Postgres")
		}
	}

	infoLogger := logger.New(logger.Config{
		Status: true,
		IP:     true,
		Method: true,
		Path:   true,
	})

	app.Use(infoLogger)
}

func main() {

	api.Routes(app)

	err := app.Run(iris.Addr(settings.GetSettings().Listen), iris.WithCharset("UTF-8"), iris.WithoutServerError(iris.ErrServerClosed))
	if err != nil {
		app.Logger().Errorln("Exiting the server, with error:", err.Error())
		return
	}
	app.Logger().Infoln("Exiting the server...")

}
