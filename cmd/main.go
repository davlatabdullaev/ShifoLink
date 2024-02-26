package main

import (
	"context"
	"log"
	"shifolink/api"
	"shifolink/config"
	"shifolink/service"
	"shifolink/storage/postgres"
	_ "shifolink/api/docs"

)

func main() {

	cfg := config.Load()

	pgStore, err := postgres.New(context.Background(), cfg)
	if err != nil {
		log.Fatalln("error while connecting to db err: ", err.Error())
		return
	}
	
	defer pgStore.CloseDB()

	services := service.New(pgStore)

	server := api.New(services, pgStore)

	if err = server.Run("localhost:8080"); err != nil {
		log.Println("error while server run")
		return
	}

}
