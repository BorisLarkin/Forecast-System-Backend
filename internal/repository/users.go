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
	intId, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	if err := r.db.Where("user_id = ?", intId).First(&User).Error; err != nil {
		return nil, err
	}
	return &User, nil
}

func (r *Repository) GetUser(login string, pwd string) (*ds.Users, error) {
	var User ds.Users
	if login == "" || pwd == "" {
		return nil, fmt.Errorf("empty login/password")
	}
	if err := r.db.Where("login = ? and password = ?", login, pwd).First(&User).Error; err != nil {
		return nil, err
	}
	return &User, nil
}

func (r *Repository) RegiterUser(Users *ds.Users) (*ds.Users, error) {
	if Users.Login == "" || Users.Password == "" {
		return nil, fmt.Errorf("login and password are required")
	}
	candidate, err := r.GetUserByLogin(Users.Login)
	if err != nil {
		return nil, err
	}
	if candidate.Login == Users.Login {
		return nil, fmt.Errorf("user with such login already exists")
	}
	err = r.db.Table("users").Create(&Users).Error
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %s", err)
	}
	return Users, nil
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
	return user.Role == 3, nil
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

	if err := r.db.Save(user); err != nil {
		return err.Error
	}
	return nil
}

func (r *Repository) Login(login string, pwd string) (int, error) {
	i, err := dsn.GetCurrentUserID()
	if err == nil { //theres an active running session
		return 0, fmt.Errorf("an already running session exists: %s", i)
	}
	user, err := r.GetUser(login, pwd)
	if err != nil {
		return 0, err
	}
	strid := strconv.Itoa(user.User_id)
	err = dsn.SetCurrentUserID(strid)
	if err != nil {
		return 0, fmt.Errorf("error starting session")
	}
	return user.User_id, nil
}

func (r *Repository) Logout() error {
	i, err := dsn.GetCurrentUserID()
	if i == "null" || err != nil {
		return fmt.Errorf("no running session found")
	}
	if err := dsn.SetCurrentUserID("null"); err != nil {
		return fmt.Errorf("failed to deauth the user")
	}
	return nil
}

func (r *Repository) GetUserByLogin(login string) (*ds.Users, error) {
	var user ds.Users
	err := r.db.Table("users").Where("login = ?", login).Find(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
