package main

import (
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	imgOpt "imageOptimisation"
	"imageOptimisation/pkg/handler"
	"imageOptimisation/pkg/repository"
	"imageOptimisation/pkg/service"
	"log"
	"os"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("Error initializing configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading env variables: %s", err.Error())
	}

	db, err := repository.NewMySqlDB(repository.Config{
		Login:    viper.GetString("db.login"),
		Password: os.Getenv("DB_PASS"),
		DBName:   viper.GetString("db.dbName"),
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
	})
	if err != nil {
		log.Fatalf("Error initialization db: %s", err.Error())
	}
	defer db.Close()
	mq, err := repository.NewRabbitMQ(repository.Config{
		Login:    viper.GetString("rabbitmq.login"),
		Password: os.Getenv("MQ_PASS"),
		Host:     viper.GetString("rabbitmq.host"),
		Port:     viper.GetString("rabbitmq.port"),
	})
	defer func() {
		err := mq.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	channel, err := mq.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := channel.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	_, err = channel.QueueDeclare(
		"images_queue",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	repos := repository.NewRepository(db)
	servs := service.NewService(repos)
	handlers := handler.NewHandler(servs)

	serv := new(imgOpt.Server)
	if err := serv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatal(err)
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
