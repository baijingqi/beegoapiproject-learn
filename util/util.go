package util

import (
    "github.com/astaxie/beego"
    _ "github.com/astaxie/beego"
    "github.com/astaxie/beego/orm"
    _ "github.com/go-sql-driver/mysql"
)
//此处定义常用全局常量
//go时间格式化
const YMDHIS = "2006-01-02 15:04:05"

func InitMysql() {
    // 注册驱动
    _ = orm.RegisterDriver("mysql", orm.DRMySQL)
    sqlCoon := beego.AppConfig.String("sqlConn")
    // 设置默认数据库
    _ = orm.RegisterDataBase("default", "mysql", sqlCoon, 30)
    // 创建 table
    _ = orm.RunSyncdb("default", false, true)
}
