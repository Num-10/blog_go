package controller

import (
	"blog_go/conf"
	"blog_go/model"
	"blog_go/pkg"
	"blog_go/util/e"
	"blog_go/util/upload"
	"crypto/md5"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"strconv"
	"strings"
	"time"
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

type article_list struct {
	ID            int    `gorm:"column:id;primary_key" json:"article_id"`
	TagID         int    `gorm:"column:tag_id" json:"-"`
	Title         string `gorm:"column:title" json:"title"`
	Content       string `gorm:"column:content" json:"content"`
	CoverImageURL string `gorm:"column:cover_image_url" json:"cover_image_url"`
	Created       int    `gorm:"column:created" json:"created"`
	Desc          string `gorm:"column:desc" json:"desc"`
	IsMarrow      int    `gorm:"column:is_marrow" json:"is_marrow"`
	IsTop         int    `gorm:"column:is_top" json:"is_top"`
	Sort          int    `gorm:"column:sort" json:"sort"`
	Status        int    `gorm:"column:status" json:"-"`
	ViewCount	  int	 `gorm:"column:view_count" json:"view_count"`
	Updated       int    `gorm:"column:updated" json:"updated"`
	CreatedFormat string `json:"created_format"`
	UpdatedFormat string `json:"updated_format"`
	TagName		  string `json:"tag_name"`
	Test		  string `json:"test"`
	Test1		  string `json:"test1"`
	Test2		  string `json:"test2"`
	Test3		  string `json:"test3"`
	Test4		  string `json:"test4"`
	Test5		  string `json:"test5"`
	Test6		  string `json:"test6"`
}

func Index(c *gin.Context)  {
	search := c.Query("search")
	tag_id, _ := strconv.Atoi(c.Query("tag_id"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	page_size, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	params := make(map[string]interface{})
	if tag_id > 0 {
		params["tag_id"] = tag_id
	}

	article := &model.Article{}
	var articleList []article_list
	article.GetList(params, map[string]interface{}{
		"page": page,
		"page_size": page_size,
		"multi_like_search": search,
		"order": "is_top desc,sort desc,created desc,id desc",
	}, &articleList, 0)

	for key, value := range articleList {
		if value.Created > 0 {
			articleList[key].CreatedFormat = time.Unix(int64(value.Created), 0).Format("2006-01-02 15:04:05")
		}
		if value.Updated > 0 {
			articleList[key].UpdatedFormat = time.Unix(int64(value.Updated), 0).Format("2006-01-02 15:04:05")
		}
		tag := &model.Tag{}
		tag.Find(map[string]interface{}{"id": value.TagID}, "")
		articleList[key].TagName = tag.Title
		view_count, _ := pkg.Redis.Get(model.ARTICLE_VIEW_COUNT_PREFIX + "id:" + strconv.Itoa(value.ID)).Int()
		articleList[key].ViewCount += view_count
		if value.CoverImageURL != "" {
			articleList[key].CoverImageURL = conf.AppIni.DomainUrl + conf.AppIni.ImageUrl + value.CoverImageURL
		}
	}

	e.Json(c, &e.Return{Code:e.SERVICE_SUCCESS, Data: articleList})
}

func SingleArticle(c *gin.Context)  {
	id, _ := strconv.Atoi(c.Param("id"))

	article := &model.Article{}
	var articleList []article_list
	article.GetList(map[string]interface{}{"id": id}, map[string]interface{}{}, &articleList, 0)

	//获取cookie
	cookie, _ := c.Cookie("view_article")

	var has_view string
	for key, value := range articleList {
		if cookie != "" {
			has_view = pkg.Redis.Get(model.ARTICLE_VIEW_COUNT_PREFIX + "cookie:" + cookie + "|id:" + strconv.Itoa(value.ID)).Val()
		} else {
			//设置cookie标记
			str := []byte(c.Request.Header.Get("User-Agent") + strconv.Itoa(int(time.Now().Unix())))
			cookie = fmt.Sprintf("%x", md5.Sum(str))
			c.SetCookie("view_article", cookie, 86400, "", "", false, false)
		}
		if has_view == "" {
			//记录文章访问人数
			now := time.Now()
			now_str := time.Now().Format("2006-01-02")
			tomorrow, _ := time.ParseInLocation("2006-01-02 15:04:05", now_str + " 23:59:59", time.Local)
			pkg.Redis.Set(model.ARTICLE_VIEW_COUNT_PREFIX + "cookie:" + cookie + "|id:" + strconv.Itoa(value.ID), 1, tomorrow.Sub(now))
			pkg.Redis.Incr(model.ARTICLE_VIEW_COUNT_PREFIX + "id:" + strconv.Itoa(value.ID))
		}

		if value.Created > 0 {
			articleList[key].CreatedFormat = time.Unix(int64(value.Created), 0).Format("2006-01-02 15:04:05")
		}
		if value.Updated > 0 {
			articleList[key].UpdatedFormat = time.Unix(int64(value.Updated), 0).Format("2006-01-02 15:04:05")
		}
		tag := &model.Tag{}
		tag.Find(map[string]interface{}{"id": value.TagID}, "")
		articleList[key].TagName = tag.Title
		view_count, _ := pkg.Redis.Get(model.ARTICLE_VIEW_COUNT_PREFIX + "id:" + strconv.Itoa(value.ID)).Int()
		articleList[key].ViewCount += view_count
		if value.CoverImageURL != "" {
			articleList[key].CoverImageURL = conf.AppIni.DomainUrl + conf.AppIni.ImageUrl + value.CoverImageURL
		}
	}

	if len(articleList) > 0 {
		e.Json(c, &e.Return{Code:e.SERVICE_SUCCESS, Data: articleList[0]})
	} else {
		e.Json(c, &e.Return{Code:e.SERVICE_SUCCESS})
	}

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
		json.Status = model.ARTICLE_STATUS_NORMAL
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

func ArticleDelete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if id <= 0 {
		e.Json(c, &e.Return{Code:e.PRRAMS_ERROR})
	}
	article := &model.Article{ID:id}
	article.Find(article, "")
	if article.Title == "" {
		e.Json(c, &e.Return{Code:e.DATA_NOT_EXISTS})
		return
	}
	err := article.Update(article, map[string]interface{}{"status": model.ARTICLE_STATUS_DELETE})
	if err != nil {
		e.Json(c, &e.Return{Code:e.SERVICE_FIAL})
		return
	}
	e.Json(c, &e.Return{Code:e.SERVICE_SUCCESS})
}

func StatisticsViewCount()  {
	keys, _ := pkg.Redis.Scan(0, model.ARTICLE_VIEW_COUNT_PREFIX + "id:*", 1000).Val()
	article := &model.Article{}
	for _, value := range keys {
		view_count, _ := pkg.Redis.Get(value).Int()
		pkg.Redis.Del(value)

		str := strings.Split(value, model.ARTICLE_VIEW_COUNT_PREFIX + "id:")
		article.ID, _ = strconv.Atoi(str[1])
		err := article.Update(article, map[string]interface{}{"view_count": gorm.Expr("view_count + ?", view_count)})

		if err != nil {
			pkg.Redis.IncrBy(value, int64(view_count))
		}
	}
}