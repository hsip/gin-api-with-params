package api

import (
	"fmt"
	"net/http"

	gingerCommon "code.byted.org/dfic/ginger/common"
	"code.byted.org/hsipeng/ginapi/common"
	"github.com/gin-gonic/gin"
)

/// 冒号:加上一个参数名组成路由参数。可以使用c.Params的方法读取其值
/// /user/format/int
func GetFormatInt(c *gin.Context) {
	id := common.GetInt(c, "id")

	c.JSON(http.StatusOK, gin.H{
		"uid": id,
	})
}

func GetFormatString(c *gin.Context) {
	nick := common.GetString(c, "nick")

	c.JSON(http.StatusOK, gin.H{
		"nick": nick,
	})
}

func GetFormatInt64(c *gin.Context) {
	id := common.GetInt64(c, "id")

	c.JSON(http.StatusOK, gin.H{
		"uid": id,
	})
}

func GetFormatInt64Again(c *gin.Context) {
	id := common.GetInt64(c, "id")
	// again c.bindJSON
	jsonData := make(map[string]interface{})
	c.BindJSON(&jsonData)
	fmt.Println("again bindJSON", jsonData)

	c.JSON(http.StatusOK, gin.H{
		"uid": id,
	})
}

type BindingParamsInt struct {
	ID string `form:"id" json:"id"`
}

func GetBindInt64(c *gin.Context) {
	var params BindingParamsInt
	err := gingerCommon.Bind(c, &params)

	if err != nil {
		fmt.Println("err:", err)
	}

	fmt.Println("params", params, c.PostForm("id"), c.Query("id"), c.Param("id"))
	c.JSON(http.StatusOK, gin.H{
		"uid": params.ID,
	})
}

func GetBindInt64JSON(c *gin.Context) {
	var params BindingParamsInt
	err := c.BindJSON(&params)

	if err != nil {
		fmt.Println("err:", err)
	}

	fmt.Println("params", params, c.PostForm("id"), c.Query("id"), c.Param("id"))
	c.JSON(http.StatusOK, gin.H{
		"uid": params.ID,
	})
}
