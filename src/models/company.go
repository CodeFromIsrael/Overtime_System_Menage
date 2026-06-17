package models

import (
	"errors"
	"strings"
)

type Company struct {
	Id          uint64 `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	LegalName   string `json:"legal_name,omitempty"`
	TaxIndefier string `json:"tax_indefier,omitempty"`
	State       string `json:"status,omitempty"`
}

func (c *Company) Prepare() error {

	if erro := c.validate(); erro != nil {
		return erro
	}
	if erro := c.format(); erro != nil {
		return erro
	}

	return nil
}

func (c *Company) validate() error {

	if c.Name == "" {
		return errors.New("o campo nome e obrigatorio")
	}

	if c.LegalName == "" {
		return errors.New("o campo nome_legael e obrigatorio")
	}

	if c.TaxIndefier == "" {
		return errors.New("o campo cnpj e obrigatorio")
	}

	return nil
}

func (c *Company) format() error {

	c.Name = strings.TrimSpace(c.Name)

	c.LegalName = strings.TrimSpace(c.LegalName)

	c.TaxIndefier = clearString(c.TaxIndefier)

	c.State = clearString(c.State)

	return nil
}
