package main

import (
	"log"

	"github.com/hibiken/asynq"
	"github.com/stg35/avito_test/internal/config"
	"github.com/stg35/avito_test/internal/db/postgres"
	"github.com/stg35/avito_test/internal/handler"
	"github.com/stg35/avito_test/internal/repository"
	"github.com/stg35/avito_test/internal/server"
	"github.com/stg35/avito_test/internal/service"
	"github.com/stg35/avito_test/internal/worker"
)

// @title Avito Test by Safoshkin Alexandr
// @version 1.0
// @description Dynamic user segmentation service

// @contact.email  alex30052003@icloud.com

// @host localhost:8000
// @BasePath /

func main() {
	config, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}
	db, err := postgres.NewConn(config.Database)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	repo, err := repository.NewRepository(db)
	if err != nil {
		log.Fatal(err)
	}
	service := service.NewService(repo)

	redisOpt := asynq.RedisClientOpt{
		Addr: config.Redis.Address,
	}
	taskDistributor := worker.NewRedisTaskDistributor(redisOpt)
	go runTaskProcessor(redisOpt, service)

	handler := handler.NewHandler(service, taskDistributor)

	server := server.NewServer(handler)
	err = server.Start(config.Server)
	if err != nil {
		log.Fatal(err)
	}
}

func runTaskProcessor(redisOpt asynq.RedisClientOpt, service *service.Service) {
	taskProcessor := worker.NewRedisTaskProcessor(redisOpt, service)
	err := taskProcessor.Start()
	if err != nil {
		log.Fatal(err)
	}
}
