package main

import (
	"fmt"
	"github.com/dilaragorum/lovely-dress/model"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

var (
	Prices = make(map[int]model.ItemPrice)
)

func main() {
	Prices[1] = model.ItemPrice{Id: 1, Price: 27.5}

	e := echo.New()
	e.GET("products/:id/price", getPrice)

	e.Logger.Fatal(e.Start(":5001"))
}

func getPrice(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err.Error())
		return echo.ErrBadRequest
	}

	if val, ok := Prices[id]; !ok {
		return echo.ErrNotFound
	} else {
		return c.JSON(http.StatusOK, val)
	}
}
