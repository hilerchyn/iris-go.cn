package app

import (
	"github.com/kataras/iris"
	"html/template"
	"github.com/kataras/iris/utils"
)

func (app *App) SetRoute(){

	app.Framework.Get("/", func(c *iris.Context){

		content := c.TemplateString("README.md", nil, iris.RenderOptions{"gzip":false})
		c.Render("content.html", struct{Content interface{}}{Content:template.HTML(content)}, iris.RenderOptions{"gzip":false})

	})


	app.Framework.Get("/:markdown", func(c *iris.Context){
		file := c.Param("markdown")

		path := "../iris-gitbook/" + file
		if !utils.DirectoryExists(path){
			c.RedirectTo("/")
			return
		}

		// serve PNG images
		if file[len(file)-4:] == ".png"{
			c.ServeFile(path, true)
		} else {
			content := c.TemplateString(file, nil, iris.RenderOptions{"gzip":false})
			c.Render("content.html", struct{Content interface{}}{Content:template.HTML(content)}, iris.RenderOptions{"gzip":false})
		}

	})
}
