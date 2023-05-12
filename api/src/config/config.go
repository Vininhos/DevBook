package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// ConnectionString é a string de conexão com o MySQL.
	ConnectionString = ""

	// Porta onde a API vai estar rodando.
	Porta = 0

	// SecretKey é a chave que será utilizada assinar o token.
	SecretKey []byte
)

// Carregar vai inicializar as variáveis de ambiente.
func Carregar() {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	Porta, erro = strconv.Atoi(os.Getenv("API_PORT"))
	if erro != nil {
		Porta = 9000
	}

	SecretKey = []byte(os.Getenv("SECRET_KEY"))

	ConnectionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USUARIO"),
		os.Getenv("DB_SENHA"),
		os.Getenv("DB_ADDRESS"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NOME"),
	)
}
