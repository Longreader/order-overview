package producer

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/Longreader/order-owerview/internal/models"
	"github.com/Longreader/order-owerview/internal/utils"
	"github.com/nats-io/stan.go"
	"github.com/sirupsen/logrus"
)

type Producer struct {
	nc      stan.Conn
	channel string
}

func NewProducer(channel string, nc stan.Conn) *Producer {
	return &Producer{
		nc:      nc,
		channel: channel,
	}
}

func (p *Producer) SubscribeAndWirte() error {

	file, err := os.ReadFile("internal/nats_streaming/producer/example.json")
	if err != nil {
		return fmt.Errorf("error to read json pattern, %w", err)
	}
	var (
		order    models.Order
		orderUID string
	)

	err = json.Unmarshal(file, &order)
	if err != nil {
		return err
	}
	for {
		orderUID = utils.GenerateUID()
		order.OrderUID = orderUID

		data, err := json.Marshal(order)
		if err != nil {
			logrus.Errorf("error  json marshal, %v", err)
			continue
		}
		logrus.Debug("write to channel")
		err = p.nc.Publish(p.channel, data)
		if err != nil {
			logrus.Errorf("error publishing, %v", err)
			continue
		}
		logrus.Debug("published to channel")
		time.Sleep(time.Second * 10)

	}

}
