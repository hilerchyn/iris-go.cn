package main

import(
  "github.com/kataras/iris"
  "github.com/kataras/iris/utils"
  "github.com/iris-contrib/template/markdown"
  "github.com/iris-contrib/template/html"
  "github.com/iris-contrib/middleware/recovery"

  customHeader "github.com/hilerchyn/iris-go.cn/middleware/header"

  "html/template"
)

func main(){

  iris.UseTemplate(html.New(html.Config{
    Layout: "layout.html",
  })).Directory("./templates", ".html")

  // use template engine
  iris.UseTemplate(markdown.New(markdown.Config{Sanitize:false})).Directory("../iris-gitbook", ".md")

  // use recovery middleware
  iris.Use(recovery.New(iris.Logger))

  // change HEADER Server
  iris.Use(&customHeader.CustomHeader{})

  iris.Get("/", func(c *iris.Context){

    content := c.TemplateString("README.md", nil, iris.RenderOptions{"gzip":false})
    c.Render("content.html", struct{Content interface{}}{Content:template.HTML(content)}, iris.RenderOptions{"gzip":false})

  })

  iris.Get("/:markdown", func(c *iris.Context){
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
 
  iris.Listen(":8080")
}
