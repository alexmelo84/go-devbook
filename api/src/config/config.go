package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// String de conexão com o banco de dados
	StringConexaoBd = ""
	// Porta da API
	Porta = 0
)

// Carregar inicializa as variáveis de ambiente da aplicação
func Carregar() {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	Porta, erro = strconv.Atoi(os.Getenv("API_PORT"))
	if erro != nil {
		Porta = 5000
	}

	StringConexaoBd = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("NETWORK_DB"),
		os.Getenv("DB_NAME"),
	)
}
