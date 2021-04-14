package api

import (
	"fmt"
	"log"
	"net/http"

	"loveraw.club/hsipeng/ginapi/model"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

/// post 获取json参数 application/json
func GetUser(c *gin.Context) {
	var user model.User
	err := c.BindJSON(&user)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"uid":      user.Uid,
		"nickName": user.NickName,
	})
}

func GetFormUser(c *gin.Context) {
	var user model.User
	err := c.BindWith(&user, binding.Form)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"uid":      user.Uid,
		"nickName": user.NickName,
	})
}

func GetUserCommon(c *gin.Context) {
	var user model.User
	err := c.Bind(&user)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"uid":      user.Uid,
		"nickName": user.NickName,
	})
}
