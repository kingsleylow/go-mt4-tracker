package server

import (
	"github.com/gin-gonic/gin"
	"github.com/hsukvn/go-mt4-tracker/controller"
)

func newRouter() *gin.Engine {
	r := gin.Default()

	ping := new(controller.PingController)
	r.GET("/ping", ping.GetController)

	orders := new(controller.OrdersController)
	r.POST("/orders", orders.PostController)

	return r
}
