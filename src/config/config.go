package config

import (
	"os"
	"github.com/joho/godotenv"
	"log"
	"strconv"
	"fmt"
)

var (
	//StringConexaoBanco é a string de conexão com o banco de dados
	StringConexaoBanco = ""
	//Porta é a porta do servidor
	Porta = 0
)

//Carregar vai inicializar as variáveis de ambiente
func Carregar() {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	Porta, erro = strconv.Atoi(os.Getenv("API_PORT"))
	if erro != nil {
		Porta = 5000
	}

	StringConexaoBanco = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
			os.Getenv("DB_USUARIO"),
			os.Getenv("DB_SENHA"),
			os.Getenv("DB_NOME"))	
}