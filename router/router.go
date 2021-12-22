package router

import (
	"github.com/gin-gonic/gin"
	"stock/db"
)

func StockRouter(e *gin.Engine, dbSrv db.Service) {
	h := NewStockHandler(dbSrv)
	e.GET("/stock/:id/detail", h.HandlerStock)
	e.GET("/stock/detail", h.HandlerStocks)
	e.GET("/stock/list", h.ListStock)
	e.PUT("/stock", h.UpdateStock)
	e.POST("/stock", h.AddStock)
	e.DELETE("/stock/:id", h.DeleteStock)
}

func Server(dbSrv db.Service) {
	r := gin.Default()
	StockRouter(r, dbSrv)
	r.Run("127.0.0.1:10000")
}
