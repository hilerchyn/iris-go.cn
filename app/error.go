package app

import "github.com/kataras/iris"

func (app *App)SetErrPage(){
	app.Framework.OnError(iris.StatusNotFound, func(c *iris.Context){
		c.RedirectTo("/")
	})
}