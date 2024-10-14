package ds

type Users struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	Login    string `gorm:"type:varchar(255)" json:"login"`
	Password string `gorm:"type:varchar(255)" json:"-"`
	IsAdmin  bool   `json:"is_admin"`
}
