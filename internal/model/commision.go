package model

import (
	"math"
)

const (
	TableName        = "commission_db"
	NormalCommission = 1.25
	Normal           = 3
	Shop             = 1
	Sales            = 2
	ShopCommission   = 13.2
	ShopTax          = 3
	SalesCommission  = 4.1
)

type Commission struct {
	Id        string
	Name      string
	Status    int
	CreatedAt string
	Price     float64
}

func NewCommission(id string, name string, status int, price float64) *Commission {

	return &Commission{
		Id:        id,
		Name:      name,
		Status:    status,
		CreatedAt: "",
		Price:     price,
	}
}

func (c *Commission) AddNewRulePrice(status int) {
	if status == Shop {
		c.Price = math.Round((c.Price * ShopCommission) + ShopTax)
		return
	}
	if status == Sales {
		c.Price = math.Round(c.Price * ShopCommission)
		return
	}
	c.Price = math.Round(c.Price * NormalCommission)
}
