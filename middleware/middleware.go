package middleware

import (
	"gopkg.in/macaron.v1"
)

var (
	checker Checker
)

// Checker is
type Checker interface {
	permissionCheck(ctx *macaron.Context)
}

// DefaultChecker is
type DefaultChecker struct {
}

func (d *DefaultChecker) permissionCheck(ctx *macaron.Context) {
}

func init() {
	if checker == nil {
		checker = new(DefaultChecker)
	}
}

// SetMiddlewares is
func SetMiddlewares(m *macaron.Macaron) {
	//Set static file directory,static file access without log output
	m.Use(macaron.Static("public", macaron.StaticOptions{
		Expires: func() string { return "max-age=0" },
	}))
	m.Use(macaron.Renderer(macaron.RenderOptions{
		Directory:  "templates",
		Charset:    "UTF-8",
		IndentJSON: true,
	}))
	//Set recovery handler to returns a middleware that recovers from any panics
	m.Use(macaron.Recovery())

	m.Use(func(ctx *macaron.Context) {
		ctx.Resp.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Resp.Header().Set("Access-Control-Allow-Methods", "POST,PUT,DELETE")
	})

	m.Use(func(ctx *macaron.Context) {
		if ctx.Req.Method == "OPTIONS" {
			ctx.Resp.Write([]byte("success"))
			ctx.Resp.Flush()
		}
	})

	m.Use(checker.permissionCheck)
}
