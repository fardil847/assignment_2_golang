package routers

import (
	"assignment_2_golang/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer(c *controllers.Controller) *gin.Engine {
	router := gin.Default()

	router.POST("/orders", c.CreateOrder)
	router.GET("/orders", c.GetOrders)
	router.PUT("/orders/:orderId", c.UpdateOrder)
	router.DELETE("/orders/:orderId", c.DeleteOrder)
	return router
}
