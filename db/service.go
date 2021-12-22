package db

import "stock"

type Service interface {
	GetStock(id int) (*stock.Stock, error)
	AddStock(stock *stock.Stock) error
	UpdateStock(stock *stock.Stock) error
	DeleteStock(id int) error
	ListStocks() ([]*stock.Stock, error)
}
