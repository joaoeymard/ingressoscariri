package main

import (
	"fmt"
	"runtime"

	routers "github.com/JoaoEymard/ingressoscariri/api"
	"github.com/gorilla/securecookie"
	//sql "github.com/JoaoEymard/ingressoscariri/service/core/database"

	"os"

	utils "github.com/JoaoEymard/ingressoscariri/api/utils"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/sessions"
)

var (
	app *iris.Application
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	app = iris.New()

	errorLogger := logger.New(logger.Config{
		Status: true,
		IP:     true,
		Method: true,
		Path:   true,
	})

	app.Use(errorLogger)

	cookieName := utils.Get().CookieName
	// AES only supports key sizes of 16, 24 or 32 bytes.
	// You either need to provide exactly that amount or you derive the key from what you type in.
	hashKey := []byte(utils.Get().HashKey)
	blockKey := []byte(utils.Get().BlockKey)
	secureCookie := securecookie.New(hashKey, blockKey)

	mySessions := sessions.New(sessions.Config{
		Cookie: cookieName,
		Encode: secureCookie.Encode,
		Decode: secureCookie.Decode,
	})

	app.AttachSessionManager(mySessions)

	/*	if err := sql.Open(); err != nil {
		if GoDetails, _ := strconv.ParseBool(os.Getenv("GO_DETAILS")); GoDetails {
			fmt.Println("[Erro] Conexão Postgres", err.Error())
		} else {
			fmt.Println("[Erro] Open Postgres")
		}
	}*/
}

func main() {

	routers.Routes(app)

	err := app.Run(iris.Addr(utils.Get().Listen))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
