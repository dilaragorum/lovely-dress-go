package model

type Stock struct {
	Id    int `json:"id"`
	Stock int `json:"stock"`
}

type ItemPrice struct {
	Id    int     `json:"id"`
	Price float64 `json:"price"`
}

type Meta struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Product struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Stock       int     `json:"stock"`
	Price       float64 `json:"price"`
}

func (s *Stock) ToProduct() *Product {
	return &Product{
		Id:    s.Id,
		Stock: s.Stock,
	}
}

func (i *ItemPrice) ToProduct() *Product {
	return &Product{
		Id:    i.Id,
		Price: i.Price,
	}
}
