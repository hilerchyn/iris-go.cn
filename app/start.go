package app

import (
	"github.com/kataras/iris"
)

type App struct {
	Framework *iris.Framework
	DefaultConfig *Config
}

func (app *App)Run(){

	app.Framework.Config.IsDevelopment = true

	// set template engine
	app.SetTemplateEngine()
	// set middleware
	app.SetMiddleware()
	// set route
	app.SetRoute()
	// set error page
	app.SetErrPage()


	app.Framework.Listen(app.DefaultConfig.Host + ":" + app.DefaultConfig.Port)
}