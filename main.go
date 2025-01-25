package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"


	"github.com/lalo64/go-kafka/src"
)

func main() {

	//* Ejecutar producer y consumer en goroutines
	src.Producer()
	src.Consumer()

	//* Manejar señales del sistema para detener el programa

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs

	fmt.Println("Shutting downs gracefully....")

}