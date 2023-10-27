package consumer

import (
	"github.com/Longreader/order-owerview/internal/service"
	"github.com/nats-io/stan.go"
	"github.com/sirupsen/logrus"
)

type Consumer struct {
	nc      stan.Conn
	channel string
	service *service.Service
}

func NewConsumer(channel string, nc stan.Conn, service *service.Service) *Consumer {
	return &Consumer{
		channel: channel,
		nc:      nc,
		service: service,
	}
}

func (c *Consumer) SubscribeAndRead() {
	_, err := c.nc.Subscribe(
		c.channel, func(msg *stan.Msg) {
			err := c.service.SetOrder(msg.Data)
			if err != nil {
				logrus.Errorf("error while set order, %v", err)
			} else {
				logrus.Debug("read from channel")
			}
		})
	if err != nil {
		logrus.Fatal(err)
	}
}
