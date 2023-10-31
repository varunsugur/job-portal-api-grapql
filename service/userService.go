package service

import (
	"golang/graph/model"
	"golang/models"
	"golang/pkg"
	"strconv"
)

func (s *Service) UserSignUp(userData model.NewUser) (*model.User, error) {
	hashPassword, err := pkg.HashPassword(userData.Password)
	if err != nil {
		return nil, err
	}
	userDetails := models.User{
		UserName:     userData.Name,
		Email:        userData.Email,
		HashPassword: hashPassword,
	}
	userDetails, err = s.userRepo.CreatUser(userDetails)
	if err != nil {
		return nil, err
	}

	uid := strconv.FormatUint(uint64(userDetails.ID), 10)

	return &model.User{
		ID:        uid,
		Name:      userDetails.UserName,
		Email:     userDetails.Email,
		CreatedAt: userDetails.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: userDetails.UpdatedAt.Format("2006-01-02 15:04:05"),
	}, nil

}
