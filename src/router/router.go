package router

import (
	"github.com/gorilla/mux"
	"api/src/router/rotas"
)	

//Router vai retornar um router com as rotas definidas
func Router() *mux.Router {
	r := mux.NewRouter()
	return rotas.Configurar(r)
}