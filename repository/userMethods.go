package repository

import "golang/models"

func (r *Repo) CreatUser(userDetails models.User) (models.User, error) {
	result := r.Db.Create(&userDetails)
	if result.Error != nil {
		return models.User{}, nil
	}
	return userDetails, nil
}
