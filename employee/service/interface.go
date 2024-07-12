package service

import "algogrit.com/empserver/entities"

type EmployeeService interface {
	Index() ([]entities.Employee, error)
	Create(newEmp entities.Employee) (*entities.Employee, error)
}
