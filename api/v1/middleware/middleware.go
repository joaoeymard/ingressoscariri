package middleware

import (
	"github.com/kataras/iris/context"
)

// func Check(ctx context.Context) {
// 	if ctrlAuth.Check(ctx) {
// 		ctx.Next()
// 	}
// 	ctx.StatusCode(iris.StatusForbidden)
// 	return
// }

func Cors(ctx context.Context) {

	ctx.Header("Access-Control-Allow-Origin", "*")

	ctx.Next()
}
