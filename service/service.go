package service

import (
	"errors"
	"golang/graph/model"
	"golang/repository"
)

type Service struct {
	userRepo repository.UserRepo
}

type UserService interface {
	UserSignUp(userData model.NewUser) (*model.User, error)
	CreatCompany(companyData model.NewCompany) (*model.Company, error)
	ViewCompanyById(cid string) (*model.Company, error)
	ViewAllcompany() ([]*model.Company, error)
	CreatJob(jobData model.NewJob) (*model.Job, error)
	ViewJobById(id string) (*model.Job, error)
	ViewJobByCid(cid string) ([]*model.Job, error)
	ViewAllJobs() ([]*model.Job, error)
}

func NewService(userRepo repository.UserRepo) (UserService, error) {
	if userRepo == nil {
		return nil, errors.New("puserRepo cannot be null")
	}
	return &Service{
		userRepo: userRepo,
	}, nil
}
