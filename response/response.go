package response

//接口标准响应
type StandardRes struct {
    Code    int         `json:"code"`
    Message string      `json:"message"`
    Data    interface{} `json:"data"`
}

//后台用户响应
func AdminUser() map[string]interface{} {
    return map[string]interface{}{
        "id":        0,
        "username":  "",
        "status":    "",
        "createdAt": "",
    }
}
