package main

import (
	"fmt"
	"github.com/dilaragorum/lovely-dress/model"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

var (
	MetaProducts = make(map[int]model.Meta)
)

func main() {
	MetaProducts[1] = model.Meta{Id: 1, Name: "Dress", Description: "Lovely Dress"}

	e := echo.New()
	e.GET("products/:id/meta", getProductMeta)

	e.Logger.Fatal(e.Start(":5002"))
}

func getProductMeta(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err.Error())
		return echo.ErrBadRequest
	}

	if val, ok := MetaProducts[id]; !ok {
		return echo.ErrNotFound
	} else {
		return c.JSON(http.StatusOK, val)
	}
}
