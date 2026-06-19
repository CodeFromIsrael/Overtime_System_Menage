package service

import (
	"errors"
	"overtime_system_menagement/src/models"
	"overtime_system_menagement/src/repository"
)

type AllocationsService struct {
	allocationRepository *repository.Allocations
}

func NewAllocationService(allocationRepository *repository.Allocations) *AllocationsService {
	return &AllocationsService{allocationRepository}
}

func (a *AllocationsService) CreateAllocation(allocation models.Allocations) (models.Allocations, error) {

	var err error

	if allocation.StartDate.Equal(allocation.EndDate) {
		return models.Allocations{}, errors.New("A data de inicio não pode ser igual a data de fim de contrato")
	}

	allocation.Id, err = a.allocationRepository.Create(allocation)

	if err != nil {
		return models.Allocations{}, err
	}

	return allocation, nil
}
