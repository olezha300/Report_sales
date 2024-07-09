package service

import (
	rgm "github.com/egorus1442/Report-Generation-Microservice"
	"github.com/egorus1442/Report-Generation-Microservice/pkg/repository"
)

type SalesListService struct {
	repo repository.SalesList
}

func NewSalesListService(repo repository.SalesList) *SalesListService {
	return &SalesListService{repo: repo}
}

func (s *SalesListService) Create(userId int, list rgm.SalesList) (int, error) {
	return s.repo.Create(userId, list)
}

func (s *SalesListService) GetAll(userId int) ([]rgm.SalesList, error) {
	return s.repo.GetAll(userId)
}
func (s *SalesListService) GetAllUserId() ([]int, error) {
	return s.repo.GetAllUserId()
}
func (s *SalesListService) GetUserNameById(id int) (string, error)  { return s.repo.GetUserNameById(id) }
func (s *SalesListService) GetBiggerSale() ([]rgm.SalesRepo, error) { return s.repo.GetBiggerSale() }
func (s *SalesListService) GetLowerSale() ([]rgm.SalesRepo, error)  { return s.repo.GetLowerSale() }
