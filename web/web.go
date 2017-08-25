package web

import (
	"gopkg.in/macaron.v1"

	"github.com/xiechuanj/blog/middleware"
	"github.com/xiechuanj/blog/router"
)

// SetPilotagedMacaron is
func SetBlogdMacaron(m *macaron.Macaron) {
	//Setting Middleware
	middleware.SetMiddlewares(m)

	//Setting Router
	router.SetRouters(m)
}
