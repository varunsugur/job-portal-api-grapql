package repository

import (
	"errors"
	"golang/models"
)

func (r *Repo) CreatCompany(companyData models.Company) (models.Company, error) {
	result := r.Db.Create(&companyData)
	if result.Error != nil {
		return models.Company{}, errors.New("could not connect to database")
	}
	return companyData, nil
}

func (r *Repo) ViewCompanyById(cid string) (models.Company, error) {
	var companyDetails models.Company

	result := r.Db.Where("id=?", cid).First(&companyDetails)
	if result.Error != nil {
		return models.Company{}, errors.New("could not found company in the record")
	}
	return companyDetails, nil
}

func (r *Repo) ViewAllCompany() ([]models.Company, error) {
	var companyDetails []models.Company

	result := r.Db.Find(&companyDetails)
	if result.Error != nil {
		return []models.Company{}, errors.New("could not found companies in record")
	}
	return companyDetails, nil
}
