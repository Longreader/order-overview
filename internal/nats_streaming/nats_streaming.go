package nats_streaming

import (
	"fmt"

	"github.com/Longreader/order-owerview/config"
	"github.com/nats-io/stan.go"
	"github.com/sirupsen/logrus"
)

func Connect(cfg config.NatsStreamingConfig) (stan.Conn, error) {

	host := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)

	sc, err := stan.Connect(cfg.ClusterID, cfg.ClientID, stan.NatsURL(host))
	if err != nil {
		logrus.Error("Error connection to NATS streaming")
		return nil, fmt.Errorf("error connection to NATS streamong, %w", err)
	}
	logrus.Debugf("%v conected to NATS and connection is %v", cfg.ClientID, stan.NatsURL(host))

	return sc, nil

}
