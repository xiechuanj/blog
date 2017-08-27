package handler

import (
	"fmt"
	"gopkg.in/macaron.v1"
)

type SignIn struct {
	Name      string `form:"uname" binding:"Required"`
	Pwd       string `form:"pwd" binding:"Required"`
	AutoLogin string `form:"autologin" binding:""`
}

func LoginPost(ctx *macaron.Context, lg SignIn) {
	uname := lg.Name
	pwd := lg.Pwd
	autoLogin := lg.AutoLogin == "on"
	if uname == "admin" && pwd == "admin12345" {
		maxAge := 0
		if autoLogin {
			maxAge = 1<<31 - 1
		}
		ctx.SetCookie("uname", uname, maxAge, "/")
		ctx.SetCookie("pwd", pwd, maxAge, "/")
	}

	ctx.Redirect("/", 301)
	return

}

func Login(ctx *macaron.Context) {
	isExit := ctx.QueryInt64("exit")
	fmt.Println(isExit)
	if isExit == 1 {
		ctx.SetCookie("uname", "", -1, "/")
		ctx.SetCookie("pwd", "", -1, "/")
		ctx.Redirect("/", 301)
		return
	}
	ctx.HTML(200, "login")
}

func checkAccount(ctx *macaron.Context) bool {
	ck, err := ctx.Req.Cookie("uname")
	if err != nil {
		return false
	}
	uname := ck.Value

	ck, err = ctx.Req.Cookie("pwd")
	if err != nil {
		return false
	}
	pwd := ck.Value

	return uname == "admin" && pwd == "admin12345"
}
