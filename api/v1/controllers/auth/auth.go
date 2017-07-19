package auth

import (
	"github.com/JoaoEymard/ingressoscariri/api/utils/session"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

func Check(ctx context.Context) {
	sess := session.GetSession().Start(ctx)
	value, _ := sess.GetBoolean("sessionCript")
	if value {
		ctx.Next()
	}
	ctx.StatusCode(iris.StatusForbidden)
}

func Login(ctx context.Context) {
	if ctx.Params().Get("user") == "c019" && ctx.Params().Get("passw") == "admin" {
		sess := session.GetSession().Start(ctx)
		sess.Set("sessionCript", true)
		ctx.JSON(map[string]interface{}{"Login": "OK"})
		return
	}
	ctx.JSON(map[string]interface{}{"Login": "Incorreto"})

}

func Logout(ctx context.Context) {
	sess := session.GetSession().Start(ctx)
	if sess.Get("sessionCript") != nil {
		sess.Delete("sessionCript")
		ctx.JSON(map[string]interface{}{"Logout": "OK"})
		return
	}
	ctx.JSON(map[string]interface{}{"Logout": "Nil"})
}
