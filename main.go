package main

import(
  "github.com/kataras/iris"
  "github.com/kataras/iris/utils"
  "github.com/iris-contrib/template/markdown"
  "github.com/iris-contrib/middleware/recovery"
)

func main(){

  // use template engine
  iris.UseTemplate(markdown.New()).Directory("../iris-gitbook", ".md")

  // use recovery middleware
  iris.Use(recovery.New(iris.Logger))

  iris.Get("/", func(c *iris.Context){
    c.Render("README.md", struct{}{}, iris.RenderOptions{"gzip":true})

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
       c.Render(file, struct{}{}, iris.RenderOptions{"gzip":true})
    }

  })
 
  iris.Listen(":8080")
}
