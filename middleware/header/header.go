package header

import "github.com/kataras/iris"

type CustomHeader struct {

}

func (ch *CustomHeader) Serve(c *iris.Context){
	c.Response.Header.Set("Server", "iris-go.cn")
	c.Next()
}
