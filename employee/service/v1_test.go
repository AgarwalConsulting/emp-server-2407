package service_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	gomock "go.uber.org/mock/gomock"

	"algogrit.com/empserver/employee/repository"
	"algogrit.com/empserver/employee/service"
	"algogrit.com/empserver/entities"
)

func TestIndex(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockRepo := repository.NewMockEmployeeRepository(ctrl)
	sut := service.NewV1(mockRepo)

	expectedEmps := []entities.Employee{
		{1, "Gaurav", "LnD", 1001},
	}
	mockRepo.EXPECT().ListAll().Return(expectedEmps, nil)

	employees, err := sut.Index()

	assert.Nil(t, err)

	assert.NotNil(t, employees)
	assert.NotEqual(t, 0, len(employees))
	assert.Equal(t, expectedEmps, employees)
}
