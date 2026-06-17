package service

import (
	"overtime_system_menagement/src/models"
	"overtime_system_menagement/src/repository"
)

type ContractService struct {
	contractRepository *repository.Contract
}

func NewContractService(contractRepository *repository.Contract) *ContractService {
	return &ContractService{contractRepository}
}

func (ce *ContractService) CreateContractEmployee(contract models.ContractEmployee) (uint64, error) {

	var err error

	contract.Status = "active"

	contract.Id, err = ce.contractRepository.CreateContract(contract)

	if err != nil {
		return 0, err
	}

	return contract.Id, nil
}

func (ce *ContractService) CreateContractCompany(contract models.ContractCompany) (models.ContractCompany, error) {

	var err error

	contract.State = "active"

	contract.Id, err = ce.contractRepository.CreateContractCompany(contract)

	if err != nil {
		return models.ContractCompany{}, err
	}

	return contract, nil
}
