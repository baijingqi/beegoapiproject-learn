package util

import (
    "beegoapiproject-learn/response"
    "github.com/fatih/structs"
    "unicode"
)

/**
 * @Description: 接口标准响应
 * @param code
 * @param message
 * @param data
 * @return response.StandardRes
 */
func MakeStdJsonRes(code int, message string, data interface{}) response.StandardRes {
    res := response.StandardRes{Code: code, Message: message, Data: data}
    return res
}

/**
 * @Description: 生成空map对象
 * @return map[string]interface{}
 */
func EmptyMap() map[string]interface{} {
    return make(map[string]interface{})
}

/**
 * @Description: 将结构体中的值赋值给响应对象
 * @param sourceObj 结构体数据
 * @param formatObj 响应对象
 * @return map[string]interface{}
 */
func StructAssign(sourceObj interface{}, formatObj map[string]interface{}) map[string]interface{} {
    sourceObjCopy := structs.Map(sourceObj)

    for key, _ := range formatObj {
        UcFirstKey := UcFirst(key)
        if _, ok := sourceObjCopy[UcFirstKey]; ok {
            formatObj[key] = sourceObjCopy[UcFirstKey]
        }
    }
    return formatObj
}

/**
 * @Description: 首字母转大写
 * @param str
 * @return string
 */
func UcFirst(str string) string {
    for i, v := range str {
        return string(unicode.ToUpper(v)) + str[i+1:]
    }
    return ""
}

/**
 * @Description: 首字母转小写
 * @param str
 * @return string
 */
func LcFirst(str string) string {
    for i, v := range str {
        return string(unicode.ToLower(v)) + str[i+1:]
    }
    return ""
}
