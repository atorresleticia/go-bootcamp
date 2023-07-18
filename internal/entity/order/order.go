package entity

import "errors"

type Order struct {
	ID         uint64  `gorm:"primary_key;auto_increment" json:"id"`
	Price      float64 `gorm:"not null" json:"price"`
	Tax        float64 `gorm:"not null" json:"tax"`
	FinalPrice float64 `gorm:"null" json:"final_price"`
}

func NewOrder(price float64, tax float64) (*Order, error) {

	order := &Order{
		Price: price,
		Tax:   tax,
	}

	err := order.Validate()

	if err != nil {
		return nil, err
	}

	return order, nil
}

func (o *Order) Validate() error {

	if o.ID == 0 {
		return errors.New("id should be greater than zero")
	}

	return nil
}

func (o *Order) CalculateFinalPrice() error {
	o.FinalPrice = o.Price + o.Tax

	err := o.Validate()

	if err != nil {
		return err
	}

	return nil
}
