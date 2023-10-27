package main

import (
	"fmt"
	"net/http"

	"github.com/Longreader/order-owerview/config"
	"github.com/Longreader/order-owerview/internal/http/handlers"
	"github.com/Longreader/order-owerview/internal/http/routers"
	"github.com/Longreader/order-owerview/internal/nats_streaming"
	"github.com/Longreader/order-owerview/internal/nats_streaming/consumer"
	"github.com/Longreader/order-owerview/internal/repository/postgres"
	"github.com/Longreader/order-owerview/internal/service"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.StandardLogger().Level = logrus.DebugLevel
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

func main() {
	cfg, err := config.NewConfig("sub")
	if err != nil {
		logrus.Fatalf("error configuration occured, %v", err)
	}

	logrus.Debug("Config complited")

	db, err := postgres.NewDataBase(cfg.DB)
	if err != nil {
		logrus.Fatalf("error database connection, %v", err)
	}

	logrus.Debug("Database complited")

	defer func() {
		if err := db.Close(); err != nil {
			logrus.Fatalf("error closing connection to database, %v", err)
		}
	}()

	serviceLogic := service.NewService(db)
	if err = serviceLogic.StartService(); err != nil {
		logrus.Fatalf("error while start database and recover cache, %v", err)
	}

	logrus.Debug("Service created complited")

	nc, err := nats_streaming.Connect(cfg.NS)
	if err != nil {
		logrus.Fatalf("error connection NATS Streaming, %v", err)
	}

	defer func() {
		if err := nc.Close(); err != nil {
			logrus.Fatalf("error close connection NATS Streaming, %v", err)
		}
	}()

	ns := consumer.NewConsumer(cfg.NS.ChanName, nc, serviceLogic)
	ns.SubscribeAndRead()

	logrus.Debug("Consumer created")

	h := handlers.NewHandler(*serviceLogic)
	r := routers.NewRouter(h)

	logrus.Debug("Start http")

	logrus.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", cfg.Http.Port), r))
}
