package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIfIGetAnErrorIfIDIsBlank(t *testing.T) {
	order := Order{}

	assert.Error(t, order.Validate(), "invalid id")
}

func TestIfIGetAnErrorIfPriceIsBlank(t *testing.T) {
	order := Order{}
	order.ID = "123"

	assert.Error(t, order.Validate(), "invalid price")
}

func TestIfIGetAnErrorIfTaxIsBlank(t *testing.T) {
	order := Order{}
	order.ID = "123"
	order.Price = 10

	assert.Error(t, order.Validate(), "invalid tax")
}

func TestWithAllValidParams(t *testing.T) {
	order := Order{}
	order.ID = "123"
	order.Price = 10
	order.Tax = 2
	order.CalculateFinalPrice()

	assert.NoError(t, order.Validate())
	assert.Equal(t, float64(10), order.Price)
	assert.Equal(t, float64(2), order.Tax)
	assert.Equal(t, 12.0, order.FinalPrice)
}
