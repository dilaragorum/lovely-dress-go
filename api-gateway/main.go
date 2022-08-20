package main

import (
	"encoding/json"
	"fmt"
	"github.com/dilaragorum/lovely-dress/model"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
	"sync"
)

func main() {
	e := echo.New()
	e.GET("/products/:id", getProductConcurrently)
	e.Logger.Fatal(e.Start(":8080"))
}

func getProductConcurrently(c echo.Context) error {
	id := c.Param("id")
	p := model.Product{}

	var wg sync.WaitGroup

	wg.Add(3)

	go getPriceFromApi(id, &p, &wg)
	go getStockFromApi(id, &p, &wg)
	go getMetaFromApi(id, &p, &wg)

	wg.Wait()

	return c.JSON(http.StatusOK, p)
}

func getPriceFromApi(id string, p *model.Product, wg *sync.WaitGroup) {
	defer wg.Done()

	priceUrl := fmt.Sprintf("http://localhost:5001/products/%s/price", id)

	response, err := makeApiCall(priceUrl)
	if err != nil {
		fmt.Println("Error when getting price", err.Error())
		return
	}

	var priceModel model.ItemPrice
	_ = json.Unmarshal(response, &priceModel)

	p.Price = priceModel.Price
	p.Id = priceModel.Id
}

func getStockFromApi(id string, p *model.Product, wg *sync.WaitGroup) {
	defer wg.Done()

	stockUrl := fmt.Sprintf("http://localhost:5000/products/%s/stock", id)

	stockRes, err := makeApiCall(stockUrl)
	if err != nil {
		fmt.Println("Error when getting stock", err.Error())
		return
	}

	var stockModel model.Stock
	_ = json.Unmarshal(stockRes, &stockModel)

	p.Stock = stockModel.Stock
	p.Id = stockModel.Id
}

func getMetaFromApi(id string, p *model.Product, wg *sync.WaitGroup) {
	defer wg.Done()

	metaUrl := fmt.Sprintf("http://localhost:5002/products/%s/meta", id)

	metaRes, err := makeApiCall(metaUrl)
	if err != nil {
		fmt.Println("Error when getting meta", err.Error())
		return
	}

	var metaModel model.Meta
	_ = json.Unmarshal(metaRes, &metaModel)

	p.Name = metaModel.Name
	p.Description = metaModel.Description
	p.Id = metaModel.Id
}

func makeApiCall(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	responseBytes, _ := io.ReadAll(response.Body)
	return responseBytes, nil
}
