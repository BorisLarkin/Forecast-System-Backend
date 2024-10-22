package ds

type Users struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	Login    string `gorm:"type:varchar(30); uniqueIndex; not null" json:"login"`
	Password string `gorm:"type:varchar(30); not null" json:"-"`
	IsAdmin  bool   `json:"is_admin"`
}
