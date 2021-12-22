package db

import (
	"stock"
	"stock/db/mysql"
)

type mysqlService struct {
}

func (s *mysqlService) GetStock(id int) (*stock.Stock, error) {
	return mysql.GetStock(id)
}

func (s *mysqlService) AddStock(stock *stock.Stock) error {
	return mysql.AddStock(stock)
}
func (s *mysqlService) UpdateStock(stock *stock.Stock) error {
	return mysql.UpdateStock(stock)
}
func (s *mysqlService) DeleteStock(id int) error {
	return mysql.DeleteStock(id)
}
func (s *mysqlService) ListStocks() ([]*stock.Stock, error) {
	return mysql.ListStock()
}

func NewMysqlService() Service {
	s := &mysqlService{}
	return s
}
