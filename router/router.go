package router

import (
	"gameserver/handler"
	"github.com/gin-gonic/gin"
)

func InitRouter()  *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("views/*")
	router.Static("/static","./static")
	router.GET("/ping", handler.Pings)
	router.GET("/html",handler.Html)
	// router.POST("/createOrder",handler.CreateOrder)
	return  router
}


