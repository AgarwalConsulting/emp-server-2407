package service

import (
	"algogrit.com/empserver/employee/repository"
	"algogrit.com/empserver/entities"
)

type v1Svc struct {
	repo repository.EmployeeRepository
}

func (svc v1Svc) Index() ([]entities.Employee, error) {
	return svc.repo.ListAll()
}

func (svc v1Svc) Create(newEmp entities.Employee) (*entities.Employee, error) {
	return svc.repo.Save(newEmp)
}

func NewV1(repo repository.EmployeeRepository) EmployeeService {
	return v1Svc{repo}
}
