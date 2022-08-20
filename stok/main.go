package main

import (
	"fmt"
	"github.com/dilaragorum/lovely-dress/model"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

var (
	Stocks = make(map[int]model.Stock)
)

func main() {
	Stocks[1] = model.Stock{Id: 1, Stock: 655}

	e := echo.New()
	e.GET("products/:id/stock", getStock)

	e.Logger.Fatal(e.Start(":5000"))
}

func getStock(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err.Error())
		return echo.ErrBadRequest
	}

	if val, ok := Stocks[id]; !ok {
		return echo.ErrNotFound
	} else {
		return c.JSON(http.StatusOK, val)
	}
}
