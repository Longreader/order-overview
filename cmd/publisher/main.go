package main

import (
	"github.com/Longreader/order-owerview/config"
	"github.com/Longreader/order-owerview/internal/nats_streaming"
	"github.com/Longreader/order-owerview/internal/nats_streaming/producer"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.StandardLogger().Level = logrus.DebugLevel
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

func main() {
	cfg, err := config.NewConfig("pub")
	if err != nil {
		logrus.Fatalf("error configuration occured, %v", err)
	}

	nc, err := nats_streaming.Connect(cfg.NS)
	if err != nil {
		logrus.Fatalf("error connection to NATS Streaming, %v", err)
	}

	logrus.Debugf("Connected to NATS Producer")

	defer func() {
		if err = nc.Close(); err != nil {
			logrus.Fatalf("error closing producer, %v", err)
		}
	}()

	ns := producer.NewProducer(cfg.NS.ChanName, nc)
	err = ns.SubscribeAndWirte()
	if err != nil {
		logrus.Fatalf("sub and write error ocured, %v", err)
	}

}
