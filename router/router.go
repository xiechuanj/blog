package router

import (
	// "github.com/go-macaron/binding"
	"gopkg.in/macaron.v1"

	"github.com/xiechuanj/blog/handler"
)

//SetRouters is pilotage router's definition function.
func SetRouters(m *macaron.Macaron) {
	m.Group("/v1", func() {
		m.Get("/", handler.IndexV1Handler)
		m.Group("/topic", func() {
			m.Get("", handler.GetTopics)
			m.Post("", handler.PostTopic)

			m.Get("/:topic", handler.GetTopic)
			m.Put("/:topic", handler.PutTopic)
			m.Delete("/:topic", handler.DeleteTopic)
		})
		m.Group("/categories", func() {
			m.Get("", handler.GetCategories)
			m.Post("", handler.PostCategory)

			m.Get("/:category", handler.GetCategory)
			m.Put("/:category", handler.PutCategory)
			m.Delete("/:category", handler.DeleteCategory)
		})
	})
}
