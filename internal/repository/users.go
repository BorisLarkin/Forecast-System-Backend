package repository

import (
	"fmt"
	"strconv"
	"web/internal/ds"
	"web/internal/dsn"
)

func (r *Repository) UserList() (*[]ds.Forecasts, error) {
	var Forecasts []ds.Forecasts
	r.db.Find(&Forecasts)
	return &Forecasts, nil
}

func (r *Repository) GetUserByID(id string) (*ds.Users, error) {
	var User ds.Users
	intId, _ := strconv.Atoi(id)
	r.db.Find(&User, intId)
	return &User, nil
}

func (r *Repository) CreateUser(Users *ds.Users) error {
	return r.db.Create(Users).Error
}

func (r *Repository) DeleteUser(id string) {
	query := "DELETE FROM users WHERE user_id = $1"
	r.db.Exec(query, id)
}
func (r *Repository) CurrentUser_IsAdmin() (bool, error) {
	user_id, err := dsn.GetCurrentUserID()
	if err != nil {
		return false, err
	}
	user, err := r.GetUserByID(user_id)
	if err != nil {
		return false, err
	}
	return user.IsAdmin, nil
}

func (r *Repository) UpdateUser(newUser ds.Users, id string) error {
	var user ds.Users
	intid, err := strconv.Atoi(id)
	if err != nil {
		return nil
	}
	if err := r.db.First(&user, intid).Error; err != nil {
		return fmt.Errorf("user %d not found", intid)
	}

	if newUser.Login != "" && newUser.Password != "" {
		user.Password = newUser.Password
		user.Login = newUser.Login
	} else {
		return fmt.Errorf("cruical info empty")
	}

	if err := r.db.Save(user).Error; err != nil {
		return err
	}
	return nil
}
func (r *Repository) Auth(id string) error {
	i, err := dsn.GetCurrentUserID()
	if i == "null" || err != nil { //theres no active running session
		_, err := r.GetUserByID(id)
		if err != nil {
			return fmt.Errorf("issue with retrieving user data")
		}
		err = dsn.SetCurrentUserID(id)
		if err != nil {
			return fmt.Errorf("error starting session")
		}
		return nil
	}
	return fmt.Errorf("an already running session exists")
}
func (r *Repository) Deauth() error {
	i, err := dsn.GetCurrentUserID()
	if i == "null" || err != nil {
		return fmt.Errorf("no running session found")
	}
	if err := dsn.SetCurrentUserID("null"); err != nil {
		return fmt.Errorf("failed to deauth the user")
	}
	return nil
}
