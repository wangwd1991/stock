package api

import "stock"

type Service interface {
	GetByCode(*stock.Stock) error
	GetByName(*stock.Stock) error
}
