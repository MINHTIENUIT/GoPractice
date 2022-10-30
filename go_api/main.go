package main

import (
	"example.com/model"
	"example.com/controller"
	"github.com/gin-gonic/gin"
)

var Articles []model.Article

func config() *gin.Engine {
	r := gin.Default()
	return r
}

func handleRequest(r *gin.Engine) {
	//Simple API:
	apiv1 := r.Group("/v1") 
	{
		//Article Controller
		a := controller.NewArticleCtl()				
		articlesCtl := apiv1.Group(a.Router) 
		{
			articlesCtl.GET("/", a.GetAll)
			articlesCtl.GET("/:id", a.GetById)
			articlesCtl.POST("/:id", a.Create)
			articlesCtl.PUT("/:id", a.Update)
			articlesCtl.PATCH("/:id", a.Update)
			articlesCtl.DELETE("/:id", a.Delete)
		}

		//Customer Controller
		c := controller.NewCustomerCtl()
		customerCtl := apiv1.Group(c.Router)
		{
			customerCtl.GET("/", c.GetAll)
			customerCtl.GET("/:id", c.GetById)
		}
	}
}

func main() {
	Articles = []model.Article{
		model.Article{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
        model.Article{Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}
	r := config();	
	handleRequest(r);
	r.Run(":8080")
}