package service

import (
	"golang/graph/model"
	"golang/models"
	"strconv"
)

func (s *Service) CreatCompany(companyData model.NewCompany) (*model.Company, error) {
	companyDetails := models.Company{
		Name:     companyData.Name,
		Location: companyData.Location,
		Salary:   companyData.Salary,
	}

	companyDetails, err := s.userRepo.CreatCompany(companyDetails)
	if err != nil {
		return nil, err
	}
	cid := strconv.FormatUint(uint64(companyDetails.ID), 10)
	return &model.Company{
		ID:       cid,
		Name:     companyDetails.Name,
		Location: companyDetails.Location,
		Salary:   companyDetails.Salary,
	}, nil
}

func (s *Service) ViewCompanyById(cid string) (*model.Company, error) {
	companyDetails, err := s.userRepo.ViewCompanyById(cid)
	if err != nil {
		return nil, err
	}
	cid = strconv.FormatUint(uint64(companyDetails.ID), 10)

	return &model.Company{
		ID:       cid,
		Name:     companyDetails.Name,
		Location: companyDetails.Location,
		Salary:   companyDetails.Salary,
	}, nil
}

func (s *Service) ViewAllcompany() ([]*model.Company, error) {
	companyDetails, err := s.userRepo.ViewAllCompany()
	if err != nil {
		return nil, err
	}
	cDatas := []*model.Company{}

	for _, v := range companyDetails {
		cData := model.Company{
			ID:       strconv.FormatUint(uint64(v.ID), 10),
			Name:     v.Name,
			Location: v.Location,
			Salary:   v.Salary,
		}
		cDatas = append(cDatas, &cData)
	}
	return cDatas, nil

}
