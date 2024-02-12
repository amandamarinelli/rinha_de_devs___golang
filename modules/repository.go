package modules

import (
	"errors"
	"rinha_de_devs___golang/models"
)

type (
	RepositoryTransaction struct{}
)

func NewRepositoryTransaction() RepositoryTransaction {
	return RepositoryTransaction{}
}

func (r RepositoryTransaction) GetClientById(id uint64) (models.Client, error) {
	c := models.Client{}
	if id == 1 {
		c = models.Client{
			Id:      1,
			Nome:    "fulano",
			Limit:   1000,
			Balance: 0,
		}
		return c, nil
	}

	err := errors.New("Seller not found")
	return c, err
}
