package controllers

import (
	"net/http"
)

// CriarUsuario cria usuário
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Cria usuário"))
}

// BuscarUsuarios cria usuário
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Busca usuários"))
}

// BuscarUsuario cria usuário
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Busca usuário"))
}

// AtualizarUsuario cria usuário
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualiza usuário"))
}

// DeletarUsuario cria usuário
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deleta usuário"))
}
