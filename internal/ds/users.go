package ds

type Users struct {
	User_id  int    `json:"user_id" gorm:"primaryKey"`
	Login    string `gorm:"type:varchar(30); uniqueIndex; not null" json:"login"`
	Password string `gorm:"type:varchar(30); not null" json:"-"`
	IsAdmin  bool   `json:"is_admin"`
	ImageURL string `json:"image_url" gorm:"type:varchar(500);default:'https://avatars.githubusercontent.com/u/116744797?s=400&u=16c7a4c1dd9f1c4bea4f6a9f3daa8a3e0a0fd69f&v=4'"`
}
