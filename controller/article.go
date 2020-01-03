package controller

import (
	"blog_go/model"
	"blog_go/util/e"
	"blog_go/util/upload"
	"github.com/gin-gonic/gin"
	"strconv"
	"unicode/utf8"
)

type article_form struct {
	TagID         int    `form:"tag_id" binding:"required"`
	Title         string `form:"title"  binding:"required"`
	Content       string `form:"content" binding:"required"`
	CoverImageURL string `form:"cover_image_url"`
	Desc          string `form:"desc" binding:"required"`
	IsMarrow      int    `form:"is_marrow"`
	IsTop         int    `form:"is_top"`
	Sort          int    `form:"sort"`
	Status		  int    `form:"status"`
}

func Index(c *gin.Context)  {
	e.Json(c, &e.Return{Code:e.SERVICE_SUCCESS})
}

func ArticleSave(c *gin.Context)  {
	var json article_form
	var err error
	id, _ := strconv.Atoi(c.Param("id"))
	user := c.GetStringMap("login_user")
	file, _ := c.FormFile("cover_image")

	if err := c.ShouldBind(&json); err != nil {
		e.Json(c, &e.Return{Code:e.PRRAMS_ERROR})
		return
	}
	if utf8.RuneCountInString(json.Title) > 50 {
		e.Json(c, &e.Return{Code:e.PRRAMS_ERROR})
		return
	}
	if json.Status == 0 {
		json.Status = 1
	}
	article := &(model.Article{
		UserID:        (user["user_id"]).(int),
		TagID:         json.TagID,
		Title:         json.Title,
		Content:       json.Content,
		CoverImageURL: json.CoverImageURL,
		Desc:          json.Desc,
		IsMarrow:      json.IsMarrow,
		IsTop:         json.IsTop,
		Sort:          json.Sort,
		Status:        json.Status,
	})
	article.Find(map[string]interface{}{"user_id": (user["user_id"]).(int), "title": json.Title}, "")
	if article.ID > 0 && article.ID != id {
		e.Json(c, &e.Return{Code:e.TITLE_IS_EXISTS})
		return
	}

	if file != nil {
		if ok, code := upload.CheckImage(file); !ok {
			e.Json(c, &e.Return{Code:code})
			return
		}

		if ok, image_url := upload.SaveImage(file, c); !ok {
			e.Json(c, &e.Return{Code:e.IMAGE_SAVE_FIAL})
			return
		} else {
			json.CoverImageURL = image_url
		}
	}

	if id > 0 {
		updateData := map[string]interface{}{
			"tag_id":          json.TagID,
			"title":           json.Title,
			"content":         json.Content,
			"desc":            json.Desc,
			"is_marrow":       json.IsMarrow,
			"is_top":          json.IsTop,
			"sort":            json.Sort,
			"status":          json.Status,
		}
		if json.CoverImageURL != "" {
			updateData["cover_image_url"] = json.CoverImageURL
		}
		err = article.Update(map[string]interface{}{"id": id}, updateData)
	} else {
		err = article.Create()
	}
	if err != nil {
		e.Json(c, &e.Return{Code:e.SERVICE_FIAL})
		return
	}
	e.Json(c, &e.Return{Code:e.SERVICE_SUCCESS})
}
