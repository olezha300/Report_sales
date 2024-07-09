package service

import (
	rgm "github.com/egorus1442/Report-Generation-Microservice"
	"github.com/egorus1442/Report-Generation-Microservice/pkg/repository"
)

type Authorization interface {
	CreateUser(user rgm.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type SalesList interface {
	Create(userId int, list rgm.SalesList) (int, error)
	GetAll(userId int) ([]rgm.SalesList, error)
	GetAllUserId() ([]int, error)
	GetUserNameById(id int) (string, error)
	GetBiggerSale() ([]rgm.SalesRepo, error)
	GetLowerSale() ([]rgm.SalesRepo, error)
}

type Service struct {
	Authorization
	SalesList
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		SalesList:     NewSalesListService(repos.SalesList),
	}
}
