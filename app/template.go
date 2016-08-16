package app

import (
	"github.com/iris-contrib/template/html"
	"github.com/iris-contrib/template/markdown"
)

func (app *App) SetTemplateEngine(){

	app.Framework.UseTemplate(html.New(html.Config{
		Layout: app.DefaultConfig.HtmlLayout,
	})).Directory(app.DefaultConfig.Html, ".html")

	// use template engine
	app.Framework.UseTemplate(markdown.New(markdown.Config{Sanitize:false,
	})).Directory(app.DefaultConfig.Markdown, ".md")


}