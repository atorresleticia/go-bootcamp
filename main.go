package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"

	"github.com/atorresleticia/go-bootcamp/internal/infra/database"
	"github.com/atorresleticia/go-bootcamp/internal/usecase"
)

func main() {

	db, err := sql.Open("sqlite3", "db.sqlite3")

	if err != nil {
		panic(err)
	}

	orderRepository := database.NewOrderRepository(db)
	uc := usecase.NewCalculateFinalPrice(orderRepository)

	orderInput := usecase.OrderInput{
		ID:    "1",
		Price: 100,
		Tax:   10,
	}

	orderOutput, err := uc.Execute(orderInput)

	if err != nil {
		panic(err)
	}

	fmt.Println(orderOutput)
}
