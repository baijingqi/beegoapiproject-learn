package main

import (
    _ "beegoapiproject-learn/routers"
    "beegoapiproject-learn/util"
    "github.com/astaxie/beego"
)

func main() {
    util.InitMysql()
    if beego.BConfig.RunMode == "dev" {
        beego.BConfig.WebConfig.DirectoryIndex = true
        beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
    }
    beego.Run()
}
