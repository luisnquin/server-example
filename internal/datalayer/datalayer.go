package datalayer

import (
	"fmt"

	"github.com/luisnquin/blind-creator-rest-api-test/internal/config"
	"github.com/luisnquin/blind-creator-rest-api-test/internal/log"
	"github.com/luisnquin/blind-creator-rest-api-test/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewForORM(config config.App) (*gorm.DB, error) {
	log.Trace().Msg("connecting to database...")

	db, err := gorm.Open(postgres.Open(generateDsnFromConfig(config)))
	if err != nil {
		return nil, err
	}

	log.Trace().Msg("successfully connected...")

	log.Trace().Msg("ensuring database tables...")

	err = db.AutoMigrate(&models.Customer{}, &models.Order{}, &models.OrderItem{}, &models.Product{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func generateDsnFromConfig(config config.App) string {
	var sslMode string

	if config.IsProduction() {
		sslMode = "require"
	} else {
		sslMode = "disable"
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Database.Host(),
		config.Database.User(),
		config.Database.Password(),
		config.Database.Name(),
		sslMode,
	)

	if port := config.Database.Port(); port != "" {
		dsn += fmt.Sprintf(" port=%s", port)
	}

	return dsn
}
