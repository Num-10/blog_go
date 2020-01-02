package controller

import (
	"blog_go/util/e"
	"fmt"
	"github.com/gin-gonic/gin"
)

type article_form struct {
	TagID         int    `form:"column:tag_id" binding:"required"`
	Title         string `form:"column:title"  binding:"required"`
	Content       string `form:"column:content" binding:"required"`
	CoverImageURL string `form:"column:cover_image_url"`
	Desc          string `form:"column:desc" binding:"required"`
	IsMarrow      int    `form:"column:is_marrow"`
	IsTop         int    `form:"column:is_top"`
	Sort          int    `form:"column:sort"`
}

func Index(c *gin.Context)  {
	e.Json(c, &e.Return{Code:e.SERVICE_SUCCESS})
}

func ArticleSave(c *gin.Context)  {
	var json article_form
	if err := c.ShouldBind(&json); err != nil {
		fmt.Println(err)
		e.Json(c, &e.Return{Code:e.PRRAMS_ERROR})
		return
	}
	fmt.Println(json)
	e.Json(c, &e.Return{Code:e.SERVICE_SUCCESS})
}
