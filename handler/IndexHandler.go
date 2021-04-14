package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"Title": "作品欣赏",
		"Msg":   "欢迎进入首页",
	})
}

func AddHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "add.html", gin.H{
		"Title": "作品欣赏 - add",
		"Msg":   "欢迎进入添加",
	})
}
