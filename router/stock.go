package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"stock"
	"stock/db"
	"stock/server"
	"strconv"
)

type response struct {
	Msg  string      `json:"msg"`
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

func (h *stockHandler) HandlerStocks(c *gin.Context) {
	stocks, err := h.srv.ListStocks()
	if err != nil {
		h.result(c, nil, err)
		return
	}
	if len(stocks) == 0 {
		h.result(c, nil, err)
		return
	}
	err = h.srv.GetStocks(stocks)
	h.result(c, stocks, err)
}

func (h *stockHandler) HandlerStock(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.result(c, id, err)
		return
	}
	stock, err := h.srv.GetStock(id)
	if err != nil {
		h.result(c, id, err)
		return
	}
	err = h.srv.GetStockDetail(stock)
	h.result(c, stock, err)
}

func (h *stockHandler) AddStock(c *gin.Context) {
	stock := stock.Stock{}
	err := c.BindJSON(&stock)
	if err != nil {
		fmt.Println("ERR:", err.Error())
		return
	}
	err = h.srv.AddStock(&stock)
	fmt.Println("ERR2:", err)
	h.result(c, stock, err)
}

func (h *stockHandler) UpdateStock(c *gin.Context) {
	stock := stock.Stock{}
	c.BindJSON(&stock)
	err := h.srv.UpdateStock(&stock)
	h.result(c, stock, err)
}

func (h *stockHandler) DeleteStock(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := h.srv.DeleteStock(id)
	h.result(c, id, err)
}

func (h *stockHandler) ListStock(c *gin.Context) {
	stocks, err := h.srv.ListStocks()
	h.result(c, stocks, err)
}

func (h *stockHandler) result(c *gin.Context, data interface{}, err error) {
	if err != nil {
		c.JSON(500, &response{
			Code: 500,
			Msg:  err.Error(),
			Data: nil,
		})
		return
	}
	c.JSON(200, &response{
		Code: 200,
		Msg:  "success",
		Data: data,
	})
}

type stockHandler struct {
	srv server.Service
}

func NewStockHandler(dbSrv db.Service) *stockHandler {
	return &stockHandler{
		srv: server.NewService(dbSrv),
	}
}
