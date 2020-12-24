package controllers

import (
    "beegoapiproject-learn/logic"
    "beegoapiproject-learn/util"
    "github.com/astaxie/beego"
)

// Operations about object
type AdminUserController struct {
    beego.Controller
}

func (o *AdminUserController) GetInfo() {
    userId, _ := o.GetUint64(":id")
    user := logic.GetUserDetail(uint(userId))
    data := util.EmptyMap()
    data["user"] = user
    o.Data["json"] = util.MakeStdJsonRes(1, "", data)
    o.ServeJSON()
}

func (o *AdminUserController) AddAdminUser() {
    username := o.GetString(":username")
    password := o.GetString(":password")
    id, err := logic.AddAdminUser(username, password)
    res := util.EmptyMap()
    code := 1
    if id <= 0 {
        code = -1
    }
    o.Data["json"] = util.MakeStdJsonRes(code, err, res)
    o.ServeJSON()
}
