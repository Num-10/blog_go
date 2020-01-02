package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Article struct {
	ID            int    `gorm:"column:id;primary_key"`
	UserID        int    `gorm:"column:user_id"`
	TagID         int    `gorm:"column:tag_id"`
	Title         string `gorm:"column:title"`
	Content       string `gorm:"column:content"`
	CoverImageURL string `gorm:"column:cover_image_url"`
	Created       int    `gorm:"column:created"`
	Desc          string `gorm:"column:desc"`
	IsMarrow      int    `gorm:"column:is_marrow"`
	IsTop         int    `gorm:"column:is_top"`
	Sort          int    `gorm:"column:sort"`
	Status        int    `gorm:"column:status"`
	Updated       int    `gorm:"column:updated"`
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

