package service

import (
	"encoding/json"
	"fmt"

	"github.com/Longreader/order-owerview/internal/models"
	"github.com/Longreader/order-owerview/internal/repository/memory"
	"github.com/Longreader/order-owerview/internal/repository/postgres"
	"github.com/sirupsen/logrus"
)

// Структура сервиса имплементирующая хранение данных
type Service struct {
	db    *postgres.PostgresStorage
	cache *memory.Cache
}

// Конструктор сервиса
func NewService(db *postgres.PostgresStorage) *Service {
	return &Service{
		db:    db,
		cache: memory.NewCache(),
	}
}

// Инициализация сервиса, восстановление кэша
func (s *Service) StartService() error {

	err := s.db.Migrate()
	if err != nil {
		return err
	}
	orders, err := s.db.GetAllOrders()
	if err != nil {
		return err
	}
	s.cache.SetOrders(orders)
	return nil
}

func (s *Service) GetOrder(orderUID string) (*models.Order, error) {
	order, err := s.cache.GetOrder(orderUID)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (s *Service) SetOrder(data []byte) error {
	var order models.Order
	err := json.Unmarshal(data, &order)
	if err != nil {
		return fmt.Errorf("error unmarshaling data: %w", err)
	}

	logrus.Debugf("ORDER UID IS %s", order.OrderUID)

	_, err = s.cache.GetOrder(order.OrderUID)
	if err == nil {
		return fmt.Errorf("order exist error: %w", err)
	}

	if err = s.db.SetOrder(order); err != nil {
		return err
	}
	s.cache.SetOrder(order)
	return nil
}
