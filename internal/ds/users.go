package ds

type Users struct {
	User_id  uint   `json:"user_id" gorm:"primaryKey"`
	Login    string `gorm:"type:varchar(30); uniqueIndex; not null" json:"login"`
	Password string `gorm:"type:varchar(30); not null" json:"password"`
	Role     int    `json:"role" gorm:"type:integer; not null"`
}

type UserRespReq struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type UserRegisterReq struct {
	Login          string `json:"login"`
	Password       string `json:"password"`
	RepeatPassword string `json:"repeat_password"`
}

type Role int

const (
	Guest Role = iota
	User
	Moderator
	Admin
)
