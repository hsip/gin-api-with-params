package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/// 冒号:加上一个参数名组成路由参数。可以使用c.Params的方法读取其值
/// /user/:id
func GetUserByID(c *gin.Context) {
	id := c.Param("id")

	c.JSON(http.StatusOK, gin.H{
		"uid": id,
	})
}

/// 查询字符串query string  即路由用，用?以后连接的key1=value2&key2=value2的形式的参数。
/// 注意默认值是指参数没有出现在url中，如果出现但没有值，则为空字符串。
func GetUserByQuery(c *gin.Context) {
	firstname := c.DefaultQuery("firstname", "Guest")
	lastname := c.Query("lastname")

	c.JSON(http.StatusOK, gin.H{
		"firstname": firstname,
		"lastname":  lastname,
	})
}

/// 报文体body参数
// body http的报文体传输数据就比query string稍微复杂一点，常见的格式就有四种。
// 例如application/json，application/x-www-form-urlencoded, application/xml和multipart/form-data。
// c.PostFROM解析的是x-www-form-urlencoded或 from-data的参数。

func GetUserInBody(c *gin.Context) {
	firstname := c.DefaultPostForm("firstname", "Guest")
	lastname := c.PostForm("lastname")

	c.JSON(http.StatusOK, gin.H{
		"firstname": firstname,
		"lastname":  lastname,
	})
}
