package service

import (
	"errors"
	"overtime_system_menagement/src/models"
	"overtime_system_menagement/src/repository"
)

type CompanyService struct {
	companyRepository *repository.Company
}

func NewCompanyService(companyRepository *repository.Company) *CompanyService {
	return &CompanyService{companyRepository}
}

func (c *CompanyService) CreateCompany(company models.Company, typeCompany string) (models.Company, error) {

	var err error

	if err = company.Prepare(); err != nil {
		return models.Company{}, err
	}

	if typeCompany == "" {
		return models.Company{}, errors.New("tipo da empresa não especificado")
	}

	company.State = "active"

	company.Id, err = c.companyRepository.CreateCompany(company, typeCompany)

	if err != nil {
		return models.Company{}, err
	}

	return company, nil

}
