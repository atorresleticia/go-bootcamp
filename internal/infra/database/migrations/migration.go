package migrations

import (
	entity "github.com/atorresleticia/go-bootcamp/internal/entity/order"
	"github.com/jinzhu/gorm"
)

func AutoMigration(db *gorm.DB) {
	db.AutoMigrate(entity.Order{})
}
