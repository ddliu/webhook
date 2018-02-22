package app

import (
	"github.com/kataras/iris"
)

// Start web server
func startServer() {
	app := iris.New()
	iris.WithoutVersionChecker(app)

	// Register custom handler for specific http errors.
	app.OnErrorCode(iris.StatusInternalServerError, func(ctx iris.Context) {
		// .Values are used to communicate between handlers, middleware.
		errMessage := ctx.Values().GetString("error")
		if errMessage != "" {
			ctx.Writef("Internal server error: %s", errMessage)
			return
		}

		ctx.Writef("(Unexpected) internal server error")
	})

	app.Any("/receive/{id}", receive)

	// Listen for incoming HTTP/1.x & HTTP/2 clients on localhost port 8080.
	app.Run(iris.Addr(GetApp().config.Listen), iris.WithCharset("UTF-8"))
}

func receive(ctx iris.Context) {
	webhookId := ctx.Params().Get("id")

	if err := GetApp().runHook(webhookId, ctx.Request()); err != nil {
		ctx.Writef("Error: " + err.Error())
		return
	}
	ctx.Writef("Received " + webhookId)
}
