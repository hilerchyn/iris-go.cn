package main

import(
  "github.com/hilerchyn/iris-go.cn/app"
  "flag"
)


var configFile = flag.String("config", "./config.json", "app configuration")

func main(){

  flag.Parse()

  server := app.New(*configFile)

  server.Run()

}
