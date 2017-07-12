package middleware

import (
	ctrlAuth "github.com/JoaoEymard/ingressoscariri/api/v1/controllers/auth"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

func Check(ctx context.Context) {
	if ctrlAuth.Check(ctx) {
		ctx.Next()
	}
	ctx.StatusCode(iris.StatusForbidden)
	return
}
