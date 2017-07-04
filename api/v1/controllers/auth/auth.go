package auth

import (
	"github.com/kataras/iris/context"
)

func Check(ctx context.Context) bool {
	session, _ := ctx.Session().GetBoolean("iddoido")
	if session {
		return true
	}
	return false

}

func Login(ctx context.Context) {
	if ctx.Params().Get("user") == "c019" && ctx.Params().Get("passw") == "admin" {
		ctx.Session().Set("iddoido", true)
	}
}
func Logoff(ctx context.Context) {
	ctx.Session().Delete("iddoido")
}
