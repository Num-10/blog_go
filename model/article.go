package model

type Article struct {
	Content       string `gorm:"column:content"`
	CoverImageURL string `gorm:"column:cover_image_url"`
	CreatedAt     int    `gorm:"column:created_at"`
	Desc          string `gorm:"column:desc"`
	ID            int    `gorm:"column:id;primary_key"`
	IsMarrow      int    `gorm:"column:is_marrow"`
	IsTop         int    `gorm:"column:is_top"`
	Sort          int    `gorm:"column:sort"`
	Status        int    `gorm:"column:status"`
	TagID         int    `gorm:"column:tag_id"`
	Title         string `gorm:"column:title"`
	UpdatedAt     int    `gorm:"column:updated_at"`
	UserID        int    `gorm:"column:user_id"`
}

// TableName sets the insert table name for this struct type
func (a *Article) TableName() string {
	return "blog_article"
}

