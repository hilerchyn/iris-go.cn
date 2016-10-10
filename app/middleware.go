package app

import (
	"github.com/iris-contrib/middleware/recovery"
	//"github.com/kataras/iris"
	customHeader "github.com/hilerchyn/iris-go.cn/middleware/header"

)

func (app *App) SetMiddleware(){
	// use recovery middleware
	app.Framework.Use(recovery.New())

	// change HEADER Server
	app.Framework.Use(&customHeader.CustomHeader{Name:app.DefaultConfig.Server})
}
