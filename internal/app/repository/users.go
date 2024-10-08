package repository

import "web/internal/app/ds"

func (r *Repository) CreateUser(user ds.Users) error {
	return r.db.Create(user).Error
}

func (r *Repository) GetUserByID(id int) (*ds.Users, error) {
	user := &ds.Users{}

	err := r.db.First(user, "id = ?", "1").Error // find product with id = 1
	if err != nil {
		return nil, err
	}

	return user, nil
}
