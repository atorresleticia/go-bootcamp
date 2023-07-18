package entity

type OrderRepositoryInterface interface {
	Save(order *Order) error
	Count() (int64, error)
	FindAll() ([]Order, error)
}
