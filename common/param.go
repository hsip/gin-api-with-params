package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetString 获取用户请求参数并转换成string
func GetString(c *gin.Context, key string, defValue ...string) (value string) {
	value = c.Param(key) // URL参数
	if value == "" {
		value = c.Query(key) // GET参数
	}
	if value == "" {
		value = c.PostForm(key) // POST参数
	}
	if value == "" { // POST application/json
		data, err := c.GetRawData()
		if err == nil {
			value = ""
		}
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data)) // fix Gin框架，body参数只能读取一次, 这边直接设置一次
		jsonData := make(map[string]interface{})                 // 二进制转换json
		err = json.Unmarshal(data, &jsonData)
		if err == nil && jsonData != nil && jsonData[key] != nil {
			// map 取出转换到string 注意json unmarshal 转出的类型
			switch jsonData[key].(type) {
			case int:
				value = fmt.Sprint(jsonData[key])
			case string:
				value = jsonData[key].(string)
			case float64:
				value = fmt.Sprint(jsonData[key])
			default:
				value = ""
			}
		}
	}
	if value == "" && len(defValue) > 0 {
		value = defValue[0]
	}
	return
}

// GetInt 获取用户请求参数并转换成int
func GetInt(c *gin.Context, key string, defValue ...int) (value int) {
	str := GetString(c, key)
	value, err := strconv.Atoi(str)
	if err != nil && len(defValue) > 0 {
		value = defValue[0]
	}
	return
}

// GetInt64 获取用户请求参数并转换成int64
func GetInt64(c *gin.Context, key string, defValue ...int64) (value int64) {
	str := GetString(c, key)
	value, err := strconv.ParseInt(str, 10, 64)
	if err != nil && len(defValue) > 0 {
		value = defValue[0]
	}
	return
}
