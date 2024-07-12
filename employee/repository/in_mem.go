package repository

import "algogrit.com/empserver/entities"

type inMemRepo struct {
	emp []entities.Employee
}

func (repo *inMemRepo) ListAll() ([]entities.Employee, error) {
	return repo.emp, nil
}

func (repo *inMemRepo) Save(newEmp entities.Employee) (*entities.Employee, error) {
	newEmp.ID = len(repo.emp) + 1
	repo.emp = append(repo.emp, newEmp)

	return &newEmp, nil
}

func NewInMemRepository() EmployeeRepository {
	var employees = []entities.Employee{
		{1, "Gaurav", "LnD", 1001},
		{2, "Lakki", "Cloud", 2001},
		{3, "Muthu", "SRE", 10001},
	}

	return &inMemRepo{
		emp: employees,
	}
}
