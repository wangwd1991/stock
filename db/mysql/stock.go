package mysql

import (
	"gorm.io/gorm"
	"stock"
	"time"
)

type Stock struct {
	gorm.Model

	Code       string  `gorm:"code"`
	Name       string  `gorm:"name"`
	ExpectSell float64 `gorm:"expect_sell"`
	ExpectBuy  float64 `gorm:"expect_buy"`
	Cost       float64 `gorm:"cost"`
	Number     int     `gorm:"number"`
}

func (Stock) TableName() string {
	return "stock"
}

func (s *Stock) Entry2Model(e *stock.Stock) {
	s.Code = e.Code
	s.Name = e.Name
	s.ExpectBuy = float64(e.ExpectBuy)
	s.ExpectSell = float64(e.ExpectSell)
	s.Cost = float64(e.Cost)
	s.Number = e.Number
}

func (s *Stock) Model2Entry() (e *stock.Stock) {
	return &stock.Stock{
		ID:         s.ID,
		Code:       s.Code,
		Name:       s.Name,
		ExpectSell: stock.MyFloat64(s.ExpectSell),
		ExpectBuy:  stock.MyFloat64(s.ExpectBuy),
		Cost:       stock.MyFloat64(s.Cost),
		Number: s.Number,
	}
}

func GetStock(id int) (*stock.Stock, error) {
	s := &Stock{}
	err := db.First(s, id).Error
	if err != nil {
		return nil, err
	}

	return s.Model2Entry(), nil
}

func AddStock(e *stock.Stock) error {
	s := &Stock{}
	s.Entry2Model(e)
	s.CreatedAt = time.Now()
	s.UpdatedAt = s.CreatedAt
	if err := db.Create(s).Error; err != nil {
		return err
	}
	e.ID = s.ID
	return nil
}

func UpdateStock(e *stock.Stock) error {
	s := &Stock{}
	s.Entry2Model(e)
	s.UpdatedAt = time.Now()
	return db.Updates(s).Error
}

func DeleteStock(id int) error {
	return db.Delete(&Stock{}, id).Error
}

func ListStock() ([]*stock.Stock, error) {
	var stocks []*Stock
	err := db.Find(&stocks).Error
	if err != nil {
		return nil, err
	}
	var entries []*stock.Stock
	for _, s := range stocks {
		entries = append(entries, s.Model2Entry())
	}
	return entries, nil
}
