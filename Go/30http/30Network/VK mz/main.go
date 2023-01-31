package main

import (
	"log"

	"30/handler"
	"30/vk"
)

func main() {
	handlers := new(handler.Handler)

	srv := new(vk.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured during running http server: %s", err.Error())
	}
}
