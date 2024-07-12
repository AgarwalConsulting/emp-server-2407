package repository

import "algogrit.com/empserver/entities"

type EmployeeRepository interface {
	ListAll() ([]entities.Employee, error)
	Save(newEmp entities.Employee) (*entities.Employee, error)
}
