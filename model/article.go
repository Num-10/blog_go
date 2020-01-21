package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

const (
	ARTICLE_STATUS_NORMAL = 1
	ARTICLE_STATUS_HIDE = -1
	ARTICLE_STATUS_DELETE = -9

	ARTICLE_VIEW_COUNT_PREFIX = "article_view_count|"
)

type Article struct {
	ID            int    `gorm:"column:id;primary_key" json:"article_id"`
	UserID        int    `gorm:"column:user_id" json:"user_id"`
	TagID         int    `gorm:"column:tag_id" json:"tag_id"`
	Title         string `gorm:"column:title" json:"title"`
	Content       string `gorm:"column:content" json:"content"`
	CoverImageURL string `gorm:"column:cover_image_url" json:"cover_image_url"`
	Created       int    `gorm:"column:created" json:"created"`
	Desc          string `gorm:"column:desc" json:"desc"`
	IsMarrow      int    `gorm:"column:is_marrow" json:"is_marrow"`
	IsTop         int    `gorm:"column:is_top" json:"is_top"`
	Sort          int    `gorm:"column:sort" json:"sort"`
	Status        int    `gorm:"column:status" json:"status"`
	ViewCount	  int	 `gorm:"column:view_count" json:"view_count"`
	Updated       int    `gorm:"column:updated" json:"updated"`
}

// TableName sets the insert table name for this struct type
func (a *Article) TableName() string {
	return "blog_article"
}

func (a *Article) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("Created", time.Now().Unix())
	return nil
}

func (a *Article) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("Updated", time.Now().Unix())
	return nil
}

func (a *Article) Create() error {
	return Db.Create(a).Error
}

func (a *Article) Find(where interface{}, order string) {
	Db.Where(where).Order(order).First(a)
}

func (a *Article) Update(where, data interface{}) error {
	return Db.Model(&Article{}).Where(where).Update(data).Error
}

func (a *Article) GetList(where map[string]interface{}, extra map[string]interface{}, articles interface{}, count *int) {
	query := Db.Model(&Article{}).Where(where)
	if _, ok := extra["field"]; ok {
		query = query.Select(extra["field"])
	}
	if _, ok := extra["multi_like_search"]; ok && extra["multi_like_search"] != "" {
		extra["multi_like_search"] = "%"+ (extra["multi_like_search"]).(string) + "%"
		query = query.Where("`title` like ? or `desc` like ? or `content` like ?", extra["multi_like_search"], extra["multi_like_search"], extra["multi_like_search"])
	}
	if _, ok := extra["status"]; !ok {
		query = query.Where("status <> ?", ARTICLE_STATUS_DELETE)
	}
	if _, ok := extra["group"]; ok {
		query = query.Group((extra["group"]).(string))
	}
	if _, ok := extra["count"]; ok {
		query = query.Count(count)
		return
	}
	if _, ok := extra["order"]; ok {
		query = query.Order(extra["order"])
	}
	page, ok := extra["page"];
	pageSize, pok := extra["page_size"];
	if ok && pok {
		query = query.Limit(pageSize).Offset(((page).(int) - 1) * (pageSize).(int))
	}
	query = query.Scan(articles)
	return
}

