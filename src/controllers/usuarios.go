package controllers

import (
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"api/src/respostas"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

//CriarUsuario é responsável por criar um novo usuário
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequisicao, erro := io.ReadAll(r.Body)
	if erro != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	//Faz o decode do JSON para o modelo de usuário
	var usuario modelos.Usuario
	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		respostas.ERRO(w, http.StatusBadRequest, erro)
		return
	}

	//Conecta no banco de dados
	db, erro := banco.Conectar()
	if erro != nil {
		respostas.ERRO(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	usuario.ID, erro = repositorio.Criar(usuario)
	if erro != nil {
		respostas.ERRO(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusCreated, usuario)
	
}

//BuscarUsuarios é responsável por buscar todos os usuários
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando usuários"))
	fmt.Println("Buscando usuários")
}

//BuscarUsuario é responsável por buscar um usuário específico
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando usuário"))
	fmt.Println("Buscando usuário")
}

//AtualizarUsuario é responsável por atualizar um usuário específico
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando usuário"))
	fmt.Println("Atualizando usuário")
}

//DeletarUsuario é responsável por deletar um usuário específico
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Deletando usuário")
}