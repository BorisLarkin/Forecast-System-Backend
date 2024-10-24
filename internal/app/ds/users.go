package ds

type Users struct {
	User_id  int    `json:"user_id" gorm:"primaryKey"`
	Login    string `gorm:"type:varchar(30); uniqueIndex; not null" json:"login"`
	Password string `gorm:"type:varchar(30); not null" json:"-"`
	IsAdmin  bool   `json:"is_admin"`
}
