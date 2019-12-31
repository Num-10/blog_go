package model

type User struct {
	ID       int    `gorm:"column:id;primary_key"`
	Password string `gorm:"column:password"`
	Username string `gorm:"column:username"`
}

// TableName sets the insert table name for this struct type
func (u *User) TableName() string {
	return "blog_user"
}

func (u *User) GetUser(where *User) {
	Db.Where(where).First(u)
}

func TestCreateUser()  {
	user1 := User{
		Username: "admin",
		Password: "123456",
	}
	Db.Create(&user1)
}