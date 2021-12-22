package server

import (
	"fmt"
	"stock"
	"stock/api"
	"stock/db"
	"sync"
)

type Service interface {
	GetStockDetail(*stock.Stock) error
	GetStocks(stocks []*stock.Stock) error

	AddStock(stock *stock.Stock) error
	GetStock(id int) (*stock.Stock, error)
	UpdateStock(stock *stock.Stock) error
	DeleteStock(id int) error
	ListStocks() ([]*stock.Stock, error)
}

type service struct {
	lock  sync.Mutex
	dbSrv db.Service
}

func (s *service) GetStocks(stocks []*stock.Stock) error {
	wg := sync.WaitGroup{}
	for i := range stocks {
		sk := stocks[i]
		wg.Add(1)
		go func(stock *stock.Stock) {
			err := s.GetStockDetail(stock)
			if err != nil {
				fmt.Println("Get stock", stock.Name, "err: ", err.Error())
			}
			wg.Done()
		}(sk)
	}
	wg.Wait()
	return nil
}

func (s *service) GetStockDetail(sk *stock.Stock) error {
	return api.NewDFCF().GetByCode(sk)
}

func (s *service) GetStock(id int) (*stock.Stock, error) {
	return s.dbSrv.GetStock(id)
}

func (s *service) AddStock(stock *stock.Stock) error {
	return s.dbSrv.AddStock(stock)
}

func (s *service) UpdateStock(stock *stock.Stock) error {
	return s.dbSrv.UpdateStock(stock)
}
func (s *service) DeleteStock(id int) error {
	return s.dbSrv.DeleteStock(id)
}

func (s *service) ListStocks() ([]*stock.Stock, error) {
	return s.dbSrv.ListStocks()
}

func NewService(dbSrv db.Service) Service {
	return &service{
		dbSrv: dbSrv,
	}
}
