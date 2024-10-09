package repository

import (
	"strconv"
	"strings"
	"web/internal/app/ds"
)

func (r *Repository) UserList() (*[]ds.Forecasts, error) {
	var Forecasts []ds.Forecasts
	r.db.Find(&Forecasts)
	return &Forecasts, nil
}

func (r *Repository) GetUserByID(id string) (*ds.Forecasts, error) {
	var Forecast ds.Forecasts
	intId, _ := strconv.Atoi(id)
	r.db.Find(&Forecast, intId)
	return &Forecast, nil
}
func (r *Repository) SearchUser(search string) (*[]ds.Users, error) {
	var Users []ds.Users
	r.db.Find(&Users)

	var filteredUsers []ds.Users
	for _, f := range Users {
		if strings.Contains(strings.ToLower(f.Login), strings.ToLower(search)) {
			filteredUsers = append(filteredUsers, f)
		}
	}
	return &filteredUsers, nil
}

func (r *Repository) CreateUser(Users ds.Users) error {
	return r.db.Create(Users).Error
}

func (r *Repository) DeleteUser(id string) {
	query := "DELETE FROM users WHERE id = $1"
	r.db.Exec(query, id)
}
