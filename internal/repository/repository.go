package repository

import (
	rgm "github.com/egorus1442/Report-Generation-Microservice"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user rgm.User) (int, error)
	GetUser(username, password string) (rgm.User, error)
}

type SalesList interface {
	Create(userId int, list rgm.SalesList) (int, error)
	GetAll(userId int) ([]rgm.SalesList, error)
	GetAllUserId() ([]int, error)
	GetUserNameById(id int) (string, error)
	GetBiggerSale() ([]rgm.SalesRepo, error)
	GetLowerSale() ([]rgm.SalesRepo, error)
}

type Repository struct {
	Authorization
	SalesList
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		SalesList:     NewSalesListPostgres(db),
	}
}
