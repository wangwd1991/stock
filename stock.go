package stock

import (
	"strconv"
)

type Stock struct {
	ID   uint   `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`

	Open      MyFloat64 `json:"open"`
	Yesterday MyFloat64 `json:"yesterday"`

	Max        MyFloat64 `json:"max"`
	Min        MyFloat64 `json:"min"`
	Current    MyFloat64 `json:"current"`
	ExpectSell MyFloat64 `json:"expectSell,string,omitempty"`
	ExpectBuy  MyFloat64 `json:"expectBuy,string,omitempty"`
	Cost       MyFloat64 `json:"cost,string,omitempty"`
	Number     int       `json:"number,string,omitempty"`

	Gains   MyFloat64 `json:"gains"`
	Status  int       `json:"status"`
	Earn    MyFloat64 `json:"earn"`
	CurEarn MyFloat64 `json:"curEarn"`
}

type MyFloat64 float64

func (m MyFloat64) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatFloat(float64(m), 'f', 2, 32)), nil
}

func (a MyFloat64) Earn(b MyFloat64, num int) MyFloat64 {
	return MyFloat64((float64(a) - float64(b)) * float64(num))
}
