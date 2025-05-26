package migrations

import (
	"fmt"

	"github.com/go-rest-frameworks-demo/fiber/internal/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	if db == nil {
		return fmt.Errorf("DB Connection is nil")
	}

	return db.AutoMigrate(
		&models.User{},
		&models.Post{},
	)
}
