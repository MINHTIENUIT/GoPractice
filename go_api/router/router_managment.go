package router

import (
	"example.com/controller"
	"github.com/gin-gonic/gin"
)

func RouterConfig() *gin.Engine {
	r := gin.Default()
	return r
}

func CreateRouter(r *gin.Engine) {
	apiv1 := r.Group("/app/v1")
	createCustomerRouter(apiv1)
}

func createCustomerRouter(g *gin.RouterGroup) {
	//Customer Controller
	c := controller.NewCustomerCtl()
	customerCtl := g.Group(c.Router)
	{
		customerCtl.GET("", c.GetAll)
		customerCtl.GET("/:id", c.GetById)
		customerCtl.POST("", c.Create)
		customerCtl.PATCH("/:id", c.Update)
		customerCtl.DELETE("/:id", c.Delete)
	}
}
