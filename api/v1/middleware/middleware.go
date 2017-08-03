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
	ctx.Header("Access-Control-Allow-Headers", "Content-Type, X-Requested-With")
	ctx.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

	ctx.Next()
}
