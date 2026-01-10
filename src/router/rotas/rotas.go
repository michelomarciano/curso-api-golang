package rotas

import (
	"net/http"
	"github.com/gorilla/mux"
)

//Rota vai representar todas as rotas da nossa aplicação
type Rota struct {
	URI string
	Metodo string
	Funcao func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}

//Configurar vai configurar as rotas da nossa aplicação
func Configurar(router *mux.Router) *mux.Router {
	for _, rota := range rotasUsuarios {
		router.HandleFunc(rota.URI, rota.Funcao).Methods(rota.Metodo)
	}
	return router
}