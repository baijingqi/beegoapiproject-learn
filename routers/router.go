// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
    "beegoapiproject-learn/controllers"
    "github.com/astaxie/beego"
)

type ObjectController struct {
}

func init() {
    ns := beego.NewNamespace("/v1",
        beego.NSNamespace("/object",
            beego.NSInclude(
                &controllers.ObjectController{},
            ),
        ),
        beego.NSNamespace("/user",
            beego.NSInclude(
                &controllers.UserController{},
            ),
        ),
    )
    beego.AddNamespace(ns)
    beego.Router("/GetAdminUserInfo/?:id", &controllers.AdminUserController{}, "get:GetInfo")
    beego.Router("/AddAdminUser/?:username/?:password", &controllers.AdminUserController{}, "get:AddAdminUser")
}
