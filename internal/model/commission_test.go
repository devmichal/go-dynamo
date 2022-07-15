package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChangePriceCommissionOfShop(t *testing.T) {

	price := Commission{Price: 4.1}

	price.AddNewRulePrice(Shop)

	assert.Equal(t, 57.00, price.Price)
}

func TestChangePriceCommissionOfSales(t *testing.T) {

	price := Commission{Price: 4.1}

	price.AddNewRulePrice(Sales)

	assert.Equal(t, 54.00, price.Price)
}

func TestChangePriceCommissionOfNormal(t *testing.T) {

	price := Commission{Price: 4.1}

	price.AddNewRulePrice(Normal)

	assert.Equal(t, 5.0, price.Price)
}
