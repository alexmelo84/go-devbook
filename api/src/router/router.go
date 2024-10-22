package router

import "github.com/gorilla/mux"

// Gerar retorna um objeto Router com as rotas configuradas
func Gerar() *mux.Router {
	return mux.NewRouter()
}
