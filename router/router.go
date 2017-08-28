package router

import (
	// "fmt"
	"github.com/go-macaron/binding"
	"github.com/xiechuanj/blog/handler"
	"gopkg.in/macaron.v1"
)

//SetRouters is pilotage router's definition function.
func SetRouters(m *macaron.Macaron) {
	m.Get("/", handler.Home)
	m.Group("/login", func() {
		m.Get("", handler.Login)
		m.Post("", binding.Bind(handler.SignIn{}), handler.LoginPost)
	})
	m.Group("/category", func() {
		m.Get("", handler.GetCategories)
	})
	m.Group("/topic", func() {
		m.Get("", handler.GetTopics)
		m.Get("/add", handler.TopicAdd)
		m.Post("", binding.Bind(handler.TopicAddForm{}), handler.TopicPost)

		m.Get("/:topicId", handler.TopicView)
		m.Get("/modify", handler.ModifyTopic)
		m.Put("", binding.Bind(handler.TopicModifyForm{}), handler.TopicPut)
		// m.Delete("/:topicId", handler.DeleteTopic)
	})
	// m.Group("/v1", func() {
	// 	m.Get("/", handler.IndexV1Handler)
	// 	m.Group("/topic", func() {
	// 		m.Get("", handler.GetTopics)
	// 		m.Post("", handler.PostTopic)

	// 		m.Get("/:topic", handler.GetTopic)
	// 		m.Put("/:topic", handler.PutTopic)
	// 		m.Delete("/:topic", handler.DeleteTopic)
	// 	})
	// 	m.Group("/categories", func() {
	// 		m.Get("", handler.GetCategories)
	// 		m.Post("", handler.PostCategory)

	// 		m.Get("/:category", handler.GetCategory)
	// 		m.Put("/:category", handler.PutCategory)
	// 		m.Delete("/:category", handler.DeleteCategory)
	// 	})
	// })
}
