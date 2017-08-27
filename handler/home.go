package handler

import (
	log "github.com/Sirupsen/logrus"
	"github.com/xiechuanj/blog/module"
	"gopkg.in/macaron.v1"
)

func Home(ctx *macaron.Context) {
	ctx.Data["IsHome"] = true

	ctx.Data["IsLogin"] = checkAccount(ctx)
	var err error
	ctx.Data["Topics"], err = module.GetTopics()
	if err != nil {
		log.Error(err.Error())
	}
	ctx.HTML(200, "home")
}
