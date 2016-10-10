package app

import (
	"github.com/kataras/iris"
	"html/template"
	"github.com/kataras/iris/utils"
)

func (app *App) SetRoute(){

	app.subdomain()

	// set website routes
	app.website()

	app.Framework.Get("/gitbook", func(c *iris.Context){

		tableContent := c.TemplateString("README.md", nil, iris.RenderOptions{"gzip":false})
		c.Render("content.html", MDContent{Content:template.HTML(tableContent),TableContent:template.HTML("")}, iris.RenderOptions{"gzip":false})

	})


	app.Framework.Get("/gitbook/:markdown", func(c *iris.Context){
		file := c.Param("markdown")

		path := app.DefaultConfig.Markdown +"/"+ file

		if !utils.Exists(path){
			c.RedirectTo("/gitbook")
			return
		}

		// serve PNG images
		if file[len(file)-4:] == ".png"{
			c.ServeFile(path, true)
		} else {
			tableContent := c.TemplateString("README.md", nil, iris.RenderOptions{"gzip":false})
			content := c.TemplateString(file, nil, iris.RenderOptions{"gzip":false})
			c.Render("content.html",
				MDContent{Content:template.HTML(content),
					TableContent:template.HTML(tableContent),
				},
				iris.RenderOptions{"gzip":false})
		}

	})
}

/**
 * set website pages
 */
func (app *App) website(){

	app.Framework.StaticServe("./static", "/static")

	app.Framework.Get("/", func(c *iris.Context){
		c.Render("site/home.html", nil, iris.RenderOptions{})
	})

	app.Framework.Get("/about", func(c *iris.Context){
		c.Render("site/about.html", nil, iris.RenderOptions{})
	})
}

/**
 * remove subdomain
 */
func (app *App) subdomain(){

	subdomain := app.Framework.Party("*.")
	{
		subdomain.Get("/:path", func(c *iris.Context) {
			c.RedirectTo("/")
			return
		})
	}
}