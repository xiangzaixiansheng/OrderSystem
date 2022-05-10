package controllers

import (
	"github.com/beego/beego/v2/server/web"
)

type UserController struct {
	web.Controller
}

func (u *UserController) GetUserById() {
	u.Ctx.WriteString("GetUserById")
}

func (u *UserController) UpdateUser() {
	u.Ctx.WriteString("UpdateUser")
}

func (u *UserController) UserHome() {
	u.Ctx.WriteString("UserHome")
}

func (u *UserController) DeleteUser() {
	u.Ctx.WriteString("DeleteUser")
}
