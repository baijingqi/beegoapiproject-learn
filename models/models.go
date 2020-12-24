package models

import (
	"fmt"
    "github.com/astaxie/beego/orm"
)


func init(){
    fmt.Println("zhuce model......")
    orm.RegisterModel(new(AdminUser))
}
