package controllers

import "net/http"

// CriarUsuario insere um usuário no banco de dados.
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando usuário."))
}

// Buscar usuário busca todos os usuários salvos no banco de dados.
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando todos os usuários."))
}

// BuscarUsuario busca um usuário específico no banco de dados.
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando um usuário."))
}

// AtualizarUsuario altera as informações de um usuário específico no banco de dados.
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando um usuário."))
}

// DeletarUsuario remove um usuário específico do banco de dados.
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando usuário."))
}
