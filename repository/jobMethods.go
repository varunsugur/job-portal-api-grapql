package repository

import (
	"errors"
	"golang/models"

	"github.com/rs/zerolog/log"
)

func (r *Repo) CreatJob(jobData models.Job) (models.Job, error) {
	result := r.Db.Create(&jobData)
	if result.Error != nil {
		return models.Job{}, errors.New("could not creat the record")
	}
	return jobData, nil
}

func (r *Repo) ViewJobById(id string) (models.Job, error) {
	var jobDetails models.Job
	result := r.Db.Where("id = ?", id).First(&jobDetails)
	if result.Error != nil {
		log.Info().Err(result.Error).Send()
		return models.Job{}, errors.New("could not find the company")
	}
	return jobDetails, nil
}

func (r *Repo) ViewJobByCid(cid string) ([]models.Job, error) {
	var jobDeatils []models.Job

	result := r.Db.Where("cid =?", cid).Find(&jobDeatils)
	if result.Error != nil {
		log.Info().Err(result.Error).Send()
		return nil, errors.New("could not find jobs")
	}
	return jobDeatils, nil
}

func (r *Repo) ViewAllJob() ([]models.Job, error) {
	var allJobs []models.Job

	result := r.Db.Find(&allJobs)
	if result.Error != nil {
		log.Info().Err(result.Error).Send()
		return nil, errors.New("could not find job")

	}
	return allJobs, nil
}
