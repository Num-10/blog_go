package controller

import (
	"blog_go/model"
	"blog_go/util/e"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
	"unicode/utf8"
)

type link_list struct {
	LinkID int `json:"link_id"`
	Title string `json:"title"`
	Desc string `json:"desc"`
	Link string `json:"link"`
	Sort int `json:"sort"`
}

func LinkList(c *gin.Context)  {
	page, _ := strconv.Atoi(c.Query("page"))
	page_size, _ := strconv.Atoi(c.Query("page_size"))
	params := make(map[string]interface{})

	link := &model.Link{}
	var linkList []model.Link
	result := make(map[string]interface{})
	var link_lists []link_list
	var count int
	extra := make(map[string]interface{})
	extra["order"] = "sort desc,id desc"
	extra["field"] = "id,title,`desc`,link,sort"
	if page > 0 && page_size > 0 {
		extra["page"] = page
		extra["page_size"] = page_size
	}

	link.GetList(params, extra, &linkList, &count)

	var link_key link_list
	for _, value := range linkList {
		link_key.LinkID = value.ID
		link_key.Link = value.Link
		link_key.Title = value.Title
		link_key.Desc = value.Desc
		link_key.Sort = value.Sort
		link_lists = append(link_lists, link_key)
	}

	result["list"] = link_lists
	e.Json(c, &e.Return{Code:e.SERVICE_SUCCESS, Data: result})
}

func LinkCreate(c *gin.Context)  {
	id, _ := strconv.Atoi(c.Param("id"))
	user := c.GetStringMap("login_user")

	title := strings.TrimSpace(c.PostForm("title"))
	desc := strings.TrimSpace(c.PostForm("desc"))
	site_link := strings.TrimSpace(c.PostForm("link"))
	sort, _ := strconv.Atoi(c.PostForm("sort"))

	if title == "" || utf8.RuneCountInString(title) > 10 {
		e.Json(c, &e.Return{Code:e.PRRAMS_ERROR})
		return
	}

	link := &(model.Link{
		UserID: (user["user_id"]).(int),
		Title: title,
	})

	link.Find(link, "")
	if link.ID > 0 && link.ID != id {
		e.Json(c, &e.Return{Code:e.TITLE_IS_EXISTS})
		return
	}
	link.Sort = sort

	var err error
	if id > 0 {
		err = link.Update(map[string]interface{}{"id": id}, map[string]interface{}{
			"title": title,
			"desc": desc,
			"link": site_link,
			"sort": sort,
		})
	} else {
		err = link.Create()
	}
	if err != nil {
		e.Json(c, &e.Return{Code:e.SERVICE_FIAL})
		return
	}
	e.Json(c, &e.Return{Code:e.SERVICE_SUCCESS})
}

func LinkDelete(c *gin.Context)  {
	id, _ := strconv.Atoi(c.Param("id"))
	user := c.GetStringMap("login_user")

	if id < 0 {
		e.Json(c, &e.Return{Code: e.PRRAMS_ERROR})
		return
	}

	link := &(model.Link{
		ID: id,
		UserID: (user["user_id"]).(int),
	})
	err := link.Delete()
	if err != nil {
		e.Json(c, &e.Return{Code: e.SERVICE_FIAL})
		return
	}
	e.Json(c, &e.Return{Code: e.SERVICE_SUCCESS})
}

func LinkFind(c *gin.Context)  {
	link_id, _ := strconv.Atoi(c.Param("id"))
	params := make(map[string]interface{})

	link := &model.Link{}
	var tagList []model.Link
	var list_lists []link_list
	var count int
	params["id"] = link_id
	extra := make(map[string]interface{})
	extra["field"] = "id,title,sort,`desc`,link"

	link.GetList(params, extra, &tagList, &count)

	var link_key link_list

	for _, value := range tagList {
		link_key.LinkID = value.ID
		link_key.Title = value.Title
		link_key.Desc = value.Desc
		link_key.Link = value.Link
		link_key.Sort = value.Sort

		list_lists = append(list_lists, link_key)
	}

	if len(list_lists) > 0 {
		e.Json(c, &e.Return{Code:e.SERVICE_SUCCESS, Data: list_lists[0]})
	} else {
		e.Json(c, &e.Return{Code:e.SERVICE_SUCCESS})
	}
}
