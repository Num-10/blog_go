package model

type Link struct {
	ID int    			`gorm:"column:id;primary_key"`
	UserID int		    `gorm:"column:user_id"`
	Title string		`gorm:"column:title"`
	Desc string			`gorm:"column:desc"`
	Link string			`gorm:"column:link"`
	Sort int			`gorm:"column:sort"`
	Created int			`gorm:"column:created"`
	Updated int			`gorm:"column:updated"`
}

func (l *Link) Create() error {
	return Db.Create(l).Error
}

func (l *Link) Update(where, data interface{}) error {
	return Db.Model(&Article{}).Where(where).Update(data).Error
}

func (l *Link) Find(where interface{}, order string) {
	Db.Where(where).Order(order).First(l)
}

func (l *Link) Delete() error {
	return Db.Delete(l).Error
}

func (l *Link) GetList(where map[string]interface{}, extra map[string]interface{}, links interface{}, count *int) {
	query := Db.Model(&Link{}).Where(where)
	if _, ok := extra["field"]; ok {
		query = query.Select(extra["field"])
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
	query = query.Scan(links)
	return
}
