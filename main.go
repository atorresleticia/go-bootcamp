package main

import entity "github.com/atorresleticia/go-bootcamp/internal/entity/order"

func main() {

	order, err := entity.NewOrder("", 0, 1)

	if err != nil {
		println(err.Error())
	}

	if order != nil {
		println(order.ID)
	}

}
