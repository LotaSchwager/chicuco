package main

import (
	"chicuco/pkg/api"
	"log"
)

func main() {
	// Instancear la API
	api := &api.Api{Addr: ":8080"}

	// Montar los endpoints
	mux := api.Mount()

	// Iniciar el servidor
	if err := api.Run(mux); err != nil {
		log.Fatal("Error al iniciar el servidor:", err)
	}
}
