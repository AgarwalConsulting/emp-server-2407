package repository

import "algogrit.com/empserver/entities"

//go:generate mockgen -source $GOFILE -destination mock_$GOPACKAGE.go -package $GOPACKAGE

type EmployeeRepository interface {
	ListAll() ([]entities.Employee, error)
	Save(newEmp entities.Employee) (*entities.Employee, error)
}
