package modules

import (
	"rinha_de_devs___golang/models"
)

type (
	ServiceTransaction struct {
		Repo Repository
	}
	Repository interface {
		GetClientById(id uint64) (models.Client, error)
	}
)

func NewService(r Repository) ServiceTransaction {
	return ServiceTransaction{
		Repo: r,
	}
}

func (s ServiceTransaction) PostTransaction(id uint64) (models.TransactionResponse, error) {
	user, err := s.Repo.GetClientById(id)
	if err != nil {
		return models.TransactionResponse{}, err
	}
	resp := models.TransactionResponse{
		Limit:   user.Limit,
		Balance: int64(user.Balance),
	}
	return resp, nil
}
