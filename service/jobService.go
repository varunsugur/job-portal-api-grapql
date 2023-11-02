package service

import (
	"golang/graph/model"
	"golang/models"
	"strconv"
)

func (s *Service) CreatJob(jobData model.NewJob) (*model.Job, error) {
	uid, err := strconv.ParseUint(jobData.Cid, 10, 32)
	if err != nil {
		return nil, err
	}

	j := models.Job{
		Cid:    uint(uid),
		Role:   jobData.Role,
		Salary: jobData.Salary,
	}

	jobDetails, err := s.userRepo.CreatJob(j)
	if err != nil {
		return nil, err
	}
	id := strconv.FormatUint(uint64(jobDetails.ID), 10)

	return &model.Job{
		ID:     id,
		Cid:    strconv.FormatUint(uint64(jobDetails.Cid), 10),
		Role:   jobDetails.Role,
		Salary: jobDetails.Salary,
	}, nil

}

func (s *Service) ViewJobById(id string) (*model.Job, error) {
	jobData, err := s.userRepo.ViewJobById(id)
	if err != nil {
		return nil, err
	}

	return &model.Job{
		ID:     strconv.FormatUint(uint64(jobData.ID), 10),
		Cid:    strconv.FormatUint(uint64(jobData.Cid), 10),
		Role:   jobData.Role,
		Salary: jobData.Salary,
	}, nil
}

func (s *Service) ViewJobByCid(cid string) ([]*model.Job, error) {
	jobDatas, err := s.userRepo.ViewJobByCid(cid)
	if err != nil {
		return nil, err
	}
	var jobDetails []*model.Job

	for _, v := range jobDatas {
		jobData := &model.Job{
			ID:     strconv.FormatUint(uint64(v.ID), 10),
			Cid:    strconv.FormatUint(uint64(v.Cid), 10),
			Role:   v.Role,
			Salary: v.Salary,
		}

		jobDetails = append(jobDetails, jobData)

	}
	return jobDetails, nil
}

func (s *Service) ViewAllJobs() ([]*model.Job, error) {
	jobDetails, err := s.userRepo.ViewAllJob()
	if err != nil {
		return nil, err
	}

	var jobDatas []*model.Job

	for _, v := range jobDetails {
		jobdata := &model.Job{
			ID:     strconv.FormatUint(uint64(v.ID), 10),
			Cid:    strconv.FormatUint(uint64(v.Cid), 10),
			Role:   v.Role,
			Salary: v.Salary,
		}
		jobDatas = append(jobDatas, jobdata)
	}
	return jobDatas, nil
}
