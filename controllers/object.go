package controllers

import (
    "beegoapiproject-learn/models"
    "github.com/astaxie/beego"
)

// Operations about object
type ObjectController struct {
    beego.Controller
}



// @Title Delete
// @Description delete the object
// @Param	objectId		path 	string	true		"The objectId you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 objectId is empty
// @router /:objectId [delete]
func (o *ObjectController) Delete() {
    objectId := o.Ctx.Input.Param(":objectId")
    models.Delete(objectId)
    o.Data["json"] = "delete success!"
    o.ServeJSON()
}
