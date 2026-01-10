package main

import (
	"fmt"
	"log"
	"net/http"
	"api/src/router"
	"api/src/config"
)	

func main() {
	config.Carregar()
	fmt.Println("Server is running on port", config.Porta)
	router := router.Router()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), router))

}