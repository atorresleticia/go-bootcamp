package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIfGetsErrorWhenIDIsBlank(t *testing.T) {

	order := Order{}

	assert.Error(t, order.Validate(), "id is mandatory")
}

func TestCalculatePrice(t *testing.T) {

	order := Order{ID:"123", Price: 10, Tax: 20}
	order.CalculateFinalPrice()

	assert.NoError(t, order.Validate())
	assert.Equal(t, 30.0, order.FinalPrice)
}
