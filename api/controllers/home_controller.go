package controllers

import (
	"net/http"

	"github.com/valentergs/booksv2/api/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Bem-vindo ao meu maravilhoso API")
}
