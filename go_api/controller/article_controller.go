package controller

import (
	"github.com/gin-gonic/gin"
)

type Article struct {
	Router string `json:"Router"`
}

func (a Article) GetAll(c *gin.Context) {
	c.JSON(200, gin.H{
		"messages": "GetAll",
	})
}

func (a Article) GetById(c *gin.Context) {
	c.JSON(200, gin.H{
		"messages": "GetById",
	})
}

func (a Article) Create(c *gin.Context) {
	c.JSON(200, gin.H{
		"messages": "Create",
	})
}

func (a Article) Update(c *gin.Context) {
	c.JSON(200, gin.H{
		"messages": "Update",
	})
}

func (a Article) Delete(c *gin.Context) {
	c.JSON(200, gin.H{
		"messages": "Delete",
	})
}

func (a Article) Query(c *gin.Context) {

}

func NewArticleCtl() Article {
	return Article{Router: "article"}
}