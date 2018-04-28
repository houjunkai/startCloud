package controllers

import (
	"github.com/astaxie/beego"
)

type UserPermission struct {
	beego.Controller
}

func (b *UserPermission) Login() {
	b.Data["Name"] = "罗芳"
	b.TplName = "login.tpl"
}
