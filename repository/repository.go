package repository

import (
	"errors"
	"golang/models"

	"gorm.io/gorm"
)

type Repo struct {
	Db *gorm.DB
}

type UserRepo interface {
	CreatUser(userDetails models.User) (models.User, error)
	CreatCompany(companyData models.Company) (models.Company, error)
	ViewCompanyById(cid string) (models.Company, error)
	ViewAllCompany() ([]models.Company, error)
	CreatJob(jobData models.Job) (models.Job, error)
	ViewJobById(id string) (models.Job, error)
	ViewJobByCid(cid string) ([]models.Job, error)
	ViewAllJob() ([]models.Job, error)
}

func NewRepo(db *gorm.DB) (UserRepo, error) {
	if db == nil {
		return nil, errors.New("db cannot be nil")
	}
	return &Repo{
		Db: db,
	}, nil

}
