package logic

import (
	"beegoapiproject-learn/models"
	"beegoapiproject-learn/response"
	"beegoapiproject-learn/util"
	"time"
)

/**
 * @Description:  获取用户详情
 * @param userId
 * @return interface
 */
func GetUserDetail(userId uint) map[string]interface{} {
	user, _ := models.GetAdminUser(userId)

	formatUser := response.AdminUser()
	if user.Id <= 0 {
		return formatUser
	}

	formatUser = util.StructAssign(user, formatUser)
	formatUser["createdAt"] = time.Unix(int64(user.CreatedAt), 0).Format(util.YMDHIS)
	return formatUser
}

/**
 * @Description: 添加用户
 * @param username
 * @param password
 * @return int64
 */
func AddAdminUser(username string, password string) (int64, string) {
	user := models.GetUsersByUsername(username, 1)
	if len(user) > 0 {
		return 0, "用户名已存在"
	}
	id := models.AddAdminUser(username, password)
	return id, ""
}
