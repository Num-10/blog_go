package controller

import (
	"blog_go/model"
	"blog_go/util/e"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

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
		e.Json(c, &e.Return{Code:e.TAG_TITLE_IS_EXISTS})
		return
	}
	tag.Sort = sort

	var err error
	if id > 0 {
		err = tag.Update(map[string]interface{}{"id": id}, map[string]interface{}{
			"updated_at": time.Now().Unix(),
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
