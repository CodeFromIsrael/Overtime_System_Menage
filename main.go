package main

import (
	"fmt"
	"log"
	"net/http"
	"overtime_system_menagement/src/config"
	"overtime_system_menagement/src/router"
)

func main() {

	config.ToLoad()

	r := router.Generete()

	fmt.Printf("Iniciando o Sevidor na Porta: %d\n", config.Door)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Door), r))
}
