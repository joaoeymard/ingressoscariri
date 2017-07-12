package main

import (
	"runtime"

	"github.com/JoaoEymard/ingressoscariri/api"
	"github.com/JoaoEymard/ingressoscariri/api/utils"
	//sql "github.com/JoaoEymard/ingressoscariri/service/core/database"

	"fmt"

	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
)

var (
	app *iris.Application
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	app = iris.New()

	/*	if err := sql.Open(); err != nil {
		if GoDetails, _ := strconv.ParseBool(os.Getenv("GO_DETAILS")); GoDetails {
			fmt.Println("[Erro] Conexão Postgres", err.Error())
		} else {
			fmt.Println("[Erro] Open Postgres")
		}
	}*/

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

	err := app.Run(iris.Addr(utils.GetSettings().Listen), iris.WithCharset("UTF-8"))
	if err != nil {
		if err != iris.ErrServerClosed {
			app.Logger().Warnf("Shutdown with error: %v\n", err.Error())
		} else {
			fmt.Print("\n")
		}
	}

}
