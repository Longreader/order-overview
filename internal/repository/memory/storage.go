package memory

import (
	"fmt"

	"github.com/Longreader/order-owerview/internal/models"
)

type Cache struct {
	cache map[string]models.Order
}

func NewCache() *Cache {
	return &Cache{
		cache: make(map[string]models.Order),
	}
}

func (c *Cache) GetOrder(orderUID string) (*models.Order, error) {
	order, ok := c.cache[orderUID]
	if !ok {
		return nil, fmt.Errorf("order not found")
	}
	return &order, nil
}

func (c *Cache) SetOrder(order models.Order) {
	c.cache[order.OrderUID] = order
}

func (c *Cache) SetOrders(orders []models.Order) {
	for _, order := range orders {
		c.SetOrder(order)
	}
}
