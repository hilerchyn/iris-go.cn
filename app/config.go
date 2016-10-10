package app

import (
	"github.com/kataras/iris"
	"encoding/json"
	"io/ioutil"
	//"github.com/kataras/iris/config"
)

type Config struct {
	Server string
	Version string
	Host string
	Port string
	Markdown string
	Html string
	HtmlLayout string
}

func New(configFile string)*App{


	// get config
	byteContent, err := ioutil.ReadFile(configFile)
	if err != nil {
		panic(err.Error())
	}

	// create app
	app := &App{}


	// parse config
	err = json.Unmarshal(byteContent, &app.DefaultConfig)
	if err != nil {
		panic(err.Error())
	}

	// disable banner
	//config := config.Iris{DisableBanner:true}
	config := iris.Configuration{DisableBanner:false}

	// set Iris Framework
	app.Framework = iris.New(config)

	return app
}