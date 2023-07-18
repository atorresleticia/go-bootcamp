package main

import (
	"fmt"

	_ "github.com/mattn/go-sqlite3"

	"github.com/atorresleticia/go-bootcamp/internal/infra/database"
	"github.com/atorresleticia/go-bootcamp/internal/infra/database/migrations"
	"github.com/atorresleticia/go-bootcamp/internal/usecase"
)

func main() {

	connection := database.Connect()
	defer connection.Close()

	migrations.AutoMigration(connection)

	orderRepository := database.NewOrderRepository(connection)
	uc := usecase.NewCalculateFinalPrice(orderRepository)

	orderOutput, err := uc.Execute(usecase.OrderInput{
		ID:    1,
		Price: 100,
		Tax:   10,
	})

	if err != nil {
		panic(err)
	}

	fmt.Println(orderOutput)
	fmt.Println(orderRepository.Count())
}
