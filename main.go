package main

import(
  "github.com/kataras/iris"
  "github.com/iris-contrib/template/markdown"
)

func main(){

  iris.UseTemplate(markdown.New()).Directory("../iris-gitbook", ".md")

  iris.Get("/", func(c *iris.Context){
    c.Render("README.md", struct{}{}, iris.RenderOptions{})

  })

  iris.Get("/:file", func(c *iris.Context){
    file := c.Param("file")

    c.Render(file, struct{}{}, iris.RenderOptions{})
  })

  println("Hello Iris!")

  iris.Listen(":8080")
}
