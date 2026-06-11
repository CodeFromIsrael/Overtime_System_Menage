package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	StringConnectionDb = ""
	Door               = 0
	SecretKey          []byte
)

func ToLoad() {

	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	Door, erro = strconv.Atoi(os.Getenv("API_PORT"))

	if erro != nil {
		Door = 9000
	}

	StringConnectionDb = fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s?parseTime=true", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_BANK"))

	SecretKey = []byte(os.Getenv("SECRET_KEY"))

}
