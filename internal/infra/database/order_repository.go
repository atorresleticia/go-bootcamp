package database

import (
	"encoding/json"
	"log"

	entity "github.com/atorresleticia/go-bootcamp/internal/entity/order"
	"github.com/jinzhu/gorm"
)

type OrderRepository struct {
	Db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{Db: db}
}

func (r *OrderRepository) Save(order *entity.Order) error {

	err := r.Db.Debug().Model(&entity.Order{}).Create(&order).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *OrderRepository) Count() (int64, error) {

	var total int64

	log.Println("Counting all orders")

	err := r.Db.Debug().Model(&entity.Order{}).Count(&total).Error

	// err := r.Db.QueryRow("SELECT COUNT(*) FROM orders").Scan(&total)

	if err != nil {
		return 0, err
	}

	return total, nil
}

func (r *OrderRepository) FindAll() ([]entity.Order, error) {

	orders := []entity.Order{}
	reply, err := Get("orders")

	if err != nil {

		log.Println("Searching on MySQL")

		err := r.Db.Debug().Model(&entity.Order{}).Limit(100).Find(&orders).Error

		if err != nil {
			return []entity.Order{}, err
		}

		ordersBytes, _ := json.Marshal(orders)

		Set("orders", ordersBytes)

		return orders, nil
	}

	log.Println("Searching on Redis")

	json.Unmarshal(reply, &orders)

	return orders, nil
}
