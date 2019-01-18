package handler

import (
	"gameserver/pay"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Pings(c *gin.Context) {
	pay.UnifiedOrder()
	c.JSON(200,gin.H{"status":"ok"})
}

func Html(c *gin.Context)  {
	c.HTML(http.StatusOK,"fudai.html", gin.H{
		"title": "布局页面",
	})
}



