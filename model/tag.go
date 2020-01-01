package model

type Tag struct {
	ID        int         `gorm:"column:id;primary_key"`
	UserID    int   	  `gorm:"column:user_id"`
	Title     string	  `gorm:"column:title"`
	Sort      int    	  `gorm:"column:sort"`
	CreatedAt int  		  `gorm:"column:created_at"`
	UpdatedAt int   	  `gorm:"column:updated_at"`
}

// TableName sets the insert table name for this struct type
func (t *Tag) TableName() string {
	return "blog_tag"
}
