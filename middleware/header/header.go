package header

import "github.com/kataras/iris"

type CustomHeader struct {
	Name string
}

func (ch *CustomHeader) Serve(c *iris.Context){

	if ch.Name == "" {
		ch.Name = "HilerChen"
	}
	c.Response.Header.Set("Server", ch.Name)
	c.Next()
}
