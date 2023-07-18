package usecase

import (
	entity "github.com/atorresleticia/go-bootcamp/internal/entity/order"
)

type OrderInput struct {
	ID    uint64
	Price float64
	Tax   float64
}

type OrderOutput struct {
	ID         uint64
	Price      float64
	Tax        float64
	FinalPrice float64
}

type CalculateFinalPrice struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewCalculateFinalPrice(orderRepository entity.OrderRepositoryInterface) *CalculateFinalPrice {
	return &CalculateFinalPrice{OrderRepository: orderRepository}
}

func (c *CalculateFinalPrice) Execute(input OrderInput) (*OrderOutput, error) {

	order, err := entity.NewOrder(input.Price, input.Tax)

	if err != nil {
		return nil, err
	}

	err = order.CalculateFinalPrice()

	if err != nil {
		return nil, err
	}

	err = c.OrderRepository.Save(order)

	if err != nil {
		return nil, err
	}

	return &OrderOutput{
		ID:         order.ID,
		Price:      order.Price,
		Tax:        order.Tax,
		FinalPrice: order.FinalPrice,
	}, nil
}
