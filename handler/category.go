package handler

import (
	// "encoding/json"
	// "fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/xiechuanj/blog/module"
	"gopkg.in/macaron.v1"
	// "net/http"
)

func GetCategories(ctx *macaron.Context) {
	op := ctx.Query("op")
	switch op {
	case "add":
		name := ctx.Query("name")
		if len(name) == 0 {
			break
		}
		err := module.AddCategory(name)
		if err != nil {
			log.Error(err.Error())
		}

		ctx.Redirect("/category", 302)
		return
	case "del":
		id := ctx.QueryInt64("id")
		if id == 0 {
			break
		}

		err := module.DeleteCategoryById(id)
		if err != nil {
			log.Error(err.Error())
		}

		ctx.Redirect("/category", 302)
	}
	ctx.Data["IsCategory"] = true
	ctx.Data["IsLogin"] = checkAccount(ctx)
	var err error
	ctx.Data["Categories"], err = module.GetAllCategories()

	if err != nil {
		log.Error(err.Error())
	}
	ctx.HTML(200, "category")
}
