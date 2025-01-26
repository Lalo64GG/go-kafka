package main

import (
	"fmt"

	"github.com/lalo64/go-kafka/src/infraestructure"
	"github.com/lalo64/go-kafka/src/server"
)

const (
	HOST = "localhost"
	PORT ="8080"
)

func main() {

	go infraestructure.Consumer()

	srv := server.NewServer(HOST, PORT);

	if err := srv.Run(); err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
	}



	// //* Ejecutar producer y consumer en goroutines
	// go src.Producer()
	// go src.Consumer()

	// //* Manejar señales del sistema para detener el programa

	// sigs := make(chan os.Signal, 1)
	// signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	// <-sigs

	// fmt.Println("Shutting downs gracefully....")

}