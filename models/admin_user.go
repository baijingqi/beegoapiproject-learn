package models

import (
    "errors"
    "fmt"
    "github.com/astaxie/beego/orm"
    "time"
)

var (
    AdminUserList map[uint]*AdminUser
)

const TableName = "admin_user"
const StatusNormal = 1 //正常
const StatusDel = 2    //删除

type AdminUser struct {
    Id        uint   `orm:"pk;column(uid);auto" json:"id"`
    Username  string `orm:"size(30)" json:"username"`
    Password  string `orm:"size(100);description(密码)" json:"password"`
    Status    uint8  `orm:"default(1);description(状态1 正常 2 删除)" json:"status"`
    CreatedAt uint   `orm:"default(0)" json:"createdAt"`
    UpdatedAt uint   `orm:"default(0)" json:"updatedAt"`
}

/**
 * @Description:  添加用户
 * @param username
 * @param password
 * @return int64
 */
func AddAdminUser(username string, password string) int64 {
    o := orm.NewOrm()
    u := AdminUser{Username: username, Password: password}
    u.CreatedAt = uint(time.Now().Unix())
    u.Status = StatusNormal
    id, err := o.Insert(&u)
    if err != nil {
        return 0
    } else {
        return id
    }
}

/**
 * @Description: 根据用户名获取用户信息
 * @param username
 * @param limit
 * @return []AdminUser
 */
func GetUsersByUsername(username string, limit int) []AdminUser {
    o := orm.NewOrm()
    var AdminUserList []AdminUser
    qs := o.QueryTable(TableName)
    _, _ = qs.Filter("username", username).Filter("status", StatusNormal).Limit(limit).All(&AdminUserList)
    return AdminUserList
}

/**
 * @Description: 获取用户信息
 * @param id
 * @return u
 * @return err
 */
func GetAdminUser(id uint) (u AdminUser, err error) {
    o := orm.NewOrm()
    adminUser := AdminUser{Id: id}
    err = o.Read(&adminUser)
    if err == nil {
        AdminUserList[id] = &adminUser
    } else {
        fmt.Println("GetAdminUser 错误信息", err)
        adminUser.Id = 0
    }
    return adminUser, err
}

func GetAllAdminUsers() map[uint]*AdminUser {
    return AdminUserList
}

func UpdateAdminUser(uId uint, uu *AdminUser) (a *AdminUser, err error) {
    if u, ok := AdminUserList[uId]; ok {
        return u, nil
    }
    return nil, errors.New("AdminUser Not Exist")
}

func DeleteAdminUser(uId uint) {
    delete(AdminUserList, uId)
}
