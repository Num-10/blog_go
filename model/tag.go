package model

import (
	"errors"
	"github.com/jinzhu/gorm"
	"time"
)

type Tag struct {
	ID        int         `gorm:"column:id;primary_key"`
	UserID    int   	  `gorm:"column:user_id"`
	Title     string	  `gorm:"column:title"`
	Sort      int    	  `gorm:"column:sort"`
	Created   int  		  `gorm:"column:created"`
	Updated   int   	  `gorm:"column:updated"`
}

// TableName sets the insert table name for this struct type
func (t *Tag) TableName() string {
	return "blog_tag"
}

func (t *Tag) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("Created", time.Now().Unix())
	return nil
}

func (t *Tag) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("Updated", time.Now().Unix())
	return nil
}

func (t *Tag) Create() error {
	return Db.Create(t).Error
}

func (t *Tag) Find(where interface{}, order string)  {
	Db.Where(where).Order(order).First(t)
}

func (t *Tag) Update(where, data interface{}) error {
	return Db.Model(&Tag{}).Where(where).Updates(data).Error
}

func (t *Tag) Delete() error {
	tx := Db.Begin()

	res := tx.Delete(t)
	if res.Error != nil || res.RowsAffected <= 0 {
		tx.Rollback()
		return errors.New("不存在该记录")
	}

	tx.Model(&Article{}).Where("tag_id = ?", t.ID).Update("tag_id", 0)

	tx.Commit()
	return nil
}