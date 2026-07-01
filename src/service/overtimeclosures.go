package service

import (
	"errors"
	"overtime_system_menagement/src/models"
	"overtime_system_menagement/src/repository"
	"time"
)

type ClousuresOvertimeServices struct {
	repositoryClusere *repository.OvertimeClousedRepository
}

func NewServiceClosuresOvertime(repository *repository.OvertimeClousedRepository) *ClousuresOvertimeServices {
	return &ClousuresOvertimeServices{repository}
}

func (cos *ClousuresOvertimeServices) CreateClosingInMonth(userInRequest uint64) (string, error) {

	var err error

	var cloused models.OvertimeCloused

	cloused.ClousedBy = userInRequest

	if userInRequest == 0 {

		return "usuário sem id", nil
	}

	todayDate := time.Now()

	if todayDate.IsZero() {
		return "", errors.New("data de criação inválida")
	}

	cloused.ClousedAt = todayDate

	cloused.Id, err = cos.repositoryClusere.MonthdyMonthByTechLeader(cloused)

	if err != nil {
		return "", err
	}

	return cloused.State, nil

}
