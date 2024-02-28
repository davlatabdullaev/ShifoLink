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

  // Config .env dan database ni manzilini olib keladi

	cfg := config.Load()

	// keyin olingan manzil postgresga berib yuboriladi va shu joydan service layerga malumot uzatiladi

	pgStore, err := postgres.New(context.Background(), cfg)
	if err != nil {
		log.Fatalln("error while connecting to db err: ", err.Error())
		return
	}
	
	defer pgStore.CloseDB()

    // service layerda biznes logikalar bajariladi

	services := service.New(pgStore)

	// keyin api orqali dastur ishga tushadi

	server := api.New(services, pgStore)

	if err = server.Run("localhost:8080"); err != nil {
		log.Println("error while server run")
		return
	}

}
