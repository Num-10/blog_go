package controller

import (
	"blog_go/model"
	"blog_go/util/e"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
	"unicode/utf8"
)

type tag_list struct {
	TagId int  `json:"tag_id"`
	Title string `json:"title"`
	ArticleCount int `json:"article_count"`
	Sort int `json:"sort"`
}

func TagList(c *gin.Context)  {
	page, _ := strconv.Atoi(c.Query("page"))
	page_size, _ := strconv.Atoi(c.Query("page_size"))
	params := make(map[string]interface{})

	tag := &model.Tag{}
	var tagList []model.Tag
	result := make(map[string]interface{})
	var tag_lists []tag_list
	var count int
	extra := make(map[string]interface{})
	extra["order"] = "sort desc,id desc"
	extra["field"] = "id,title,sort"
	if page > 0 && page_size > 0 {
		extra["page"] = page
		extra["page_size"] = page_size
	}

	tag.GetList(params, map[string]interface{}{"count": true}, &tagList, &count)
	result["count"] = count

	if (count > 0) {
		tag.GetList(params, extra, &tagList, &count)

		article := &model.Article{}
		var articleList []string
		var tag_key tag_list

		for _, value := range tagList {
			tag_key.TagId = value.ID
			tag_key.Title = value.Title
			tag_key.Sort = value.Sort
			article.GetList(map[string]interface{}{"tag_id": value.ID}, map[string]interface{}{"count": true}, &articleList, &count)
			tag_key.ArticleCount = count
			tag_lists = append(tag_lists, tag_key)
		}
	}


	result["list"] = tag_lists
	e.Json(c, &e.Return{Code:e.SERVICE_SUCCESS, Data: result})
}

func TagCreate(c *gin.Context)  {
	id, _ := strconv.Atoi(c.Param("id"))
	user := c.GetStringMap("login_user")

	title := strings.TrimSpace(c.PostForm("title"))
	sort, _ := strconv.Atoi(c.PostForm("sort"))

	if title == "" || utf8.RuneCountInString(title) > 10 {
		e.Json(c, &e.Return{Code:e.PRRAMS_ERROR})
		return
	}

	tag := &(model.Tag{
		UserID: (user["user_id"]).(int),
		Title: title,
	})

	tag.Find(tag, "")
	if tag.ID > 0 && tag.ID != id {
		e.Json(c, &e.Return{Code:e.TITLE_IS_EXISTS})
		return
	}
	tag.Sort = sort

	var err error
	if id > 0 {
		err = tag.Update(map[string]interface{}{"id": id}, map[string]interface{}{
			"title": title,
			"sort": sort,
		})
	} else {
		err = tag.Create()
	}
	if err != nil {
		e.Json(c, &e.Return{Code:e.SERVICE_FIAL})
		return
	}
	e.Json(c, &e.Return{Code:e.SERVICE_SUCCESS})
}

func TagDelete(c *gin.Context)  {
	id, _ := strconv.Atoi(c.Param("id"))
	user := c.GetStringMap("login_user")

	if id < 0 {
		e.Json(c, &e.Return{Code: e.PRRAMS_ERROR})
		return
	}

	tag := &(model.Tag{
		ID: id,
		UserID: (user["user_id"]).(int),
	})
	err := tag.Delete()
	if err != nil {
		e.Json(c, &e.Return{Code: e.SERVICE_FIAL})
		return
	}
	e.Json(c, &e.Return{Code: e.SERVICE_SUCCESS})
}

func TagFind(c *gin.Context)  {
	tag_id, _ := strconv.Atoi(c.Param("id"))
	params := make(map[string]interface{})

	tag := &model.Tag{}
	var tagList []model.Tag
	var tag_lists []tag_list
	var count int
	params["id"] = tag_id
	extra := make(map[string]interface{})
	extra["field"] = "id,title,sort"

	tag.GetList(params, extra, &tagList, &count)

	var tag_key tag_list

	for _, value := range tagList {
		tag_key.TagId = value.ID
		tag_key.Title = value.Title
		tag_key.Sort = value.Sort

		tag_lists = append(tag_lists, tag_key)
	}

	if len(tag_lists) > 0 {
		e.Json(c, &e.Return{Code:e.SERVICE_SUCCESS, Data: tag_lists[0]})
	} else {
		e.Json(c, &e.Return{Code:e.SERVICE_SUCCESS})
	}
}
