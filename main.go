package main

import (
	"fmt"
	"log"
	"net/http"
	"overtime_system_menagement/src/bootstrap"
	"overtime_system_menagement/src/config"
	"overtime_system_menagement/src/datebase"
)

func main() {

	config.ToLoad()

	DB, err := datebase.Connection()

	if err != nil {
		log.Fatal(err)
	}

	defer DB.Close()

	r := bootstrap.Initialize(DB)

	fmt.Printf("Iniciando o Sevidor na Porta: %d\n", config.Door)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Door), r))
}
