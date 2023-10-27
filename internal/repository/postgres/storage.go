package postgres

import (
	"fmt"

	"github.com/Longreader/order-owerview/config"
	"github.com/Longreader/order-owerview/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresStorage struct {
	db *gorm.DB
}

func NewDataBase(cfg config.PostgresConfig) (*PostgresStorage, error) {

	dsn := fmt.Sprintf(
		"postgres://%s:%s@postgres/%s?sslmode=disable",
		cfg.User,
		cfg.Password,
		cfg.Name,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error connection to database: %w", err)
	}
	return &PostgresStorage{db: db}, nil
}

func (ps *PostgresStorage) Migrate() error {
	err := ps.db.AutoMigrate(&models.Order{})
	if err != nil {
		return fmt.Errorf("error gorm migration: %w", err)
	}
	return nil
}

func (ps *PostgresStorage) SetOrder(order models.Order) error {
	tx := ps.db.Create(&order)
	if tx.Error != nil {
		return fmt.Errorf("error while insert data: %w", tx.Error)
	}
	return nil
}

func (ps *PostgresStorage) GetAllOrders() (orders []models.Order, err error) {
	tx := ps.db.Find(&orders)
	if tx.Error != nil {
		return nil, fmt.Errorf("error while recover data form database: %w", tx.Error)
	}
	return orders, nil
}

func (ps *PostgresStorage) Close() error {
	dbConn, _ := ps.db.DB()
	if err := dbConn.Close(); err != nil {
		return err
	}
	return nil
}
