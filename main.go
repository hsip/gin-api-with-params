package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"code.byted.org/hsipeng/ginapi/api"
	"code.byted.org/hsipeng/ginapi/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/tpl/*/*")
	router.GET("/", handler.IndexHandler)
	v := router.Group("/")
	{
		v.GET("/index.html", handler.IndexHandler)
		v.GET("/add.html", handler.AddHandler)
	}

	v1 := router.Group("/v1")
	{
		v1.GET("/user/:id", api.GetUserByID)
		v1.GET("/user/query", api.GetUserByQuery)
		v1.POST("/user/post", api.GetUserInBody)
		v1.POST("/user/post/json", api.GetUser)
		v1.POST("/user/post/form", api.GetFormUser)
		v1.POST("/user/post/common", api.GetUserCommon)
	}

	v2 := router.Group("/v2")
	{
		v2.GET("/user/format/i/:id", api.GetFormatInt)
		v2.GET("/user/format/int", api.GetFormatInt)
		v2.POST("/user/format/int", api.GetFormatInt)
		v2.GET("/user/format/i64/:id", api.GetFormatInt64)
		v2.GET("/user/format/int64", api.GetFormatInt64)
		v2.POST("/user/format/int64", api.GetFormatInt64)

		v2.POST("/user/format/int642", api.GetFormatInt64Again)

		v2.GET("/user/format/s/:nick", api.GetFormatString)
		v2.GET("/user/format/string", api.GetFormatString)
		v2.POST("/user/format/string", api.GetFormatString)

		v2.GET("/user/format/bind", api.GetBindInt64)
		v2.POST("/user/format/bind", api.GetBindInt64)
		v2.POST("/user/format/json", api.GetBindInt64JSON)
	}
	srv := &http.Server{
		Addr:         ":80",
		Handler:      router,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)

	signal.Notify(quit, os.Interrupt)

	<-quit

	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	log.Println("Server exit.")
}
