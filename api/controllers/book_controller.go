package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/valentergs/booksv2/api/auth"
	"github.com/valentergs/booksv2/api/models"
	"github.com/valentergs/booksv2/api/responses"
	"github.com/valentergs/booksv2/api/utils/formaterror"
)

func (server *Server) CreateBook(w http.ResponseWriter, r *http.Request) {
	book := models.Book{}
	json.NewDecoder(r.Body).Decode(&book)
	book.Prepare()
	err := book.Validate("create")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	bookCreated, err := book.SaveBook(server.DB)

	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, bookCreated.ID))
	responses.JSON(w, http.StatusCreated, bookCreated)
}

func (server *Server) GetBooks(w http.ResponseWriter, r *http.Request) {
	book := models.Book{}
	books, err := book.FindAllBooks(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, books)
}

func (server *Server) GetBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bid, err := strconv.Atoi(params["id"])
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	book := models.Book{}
	bookGotten, err := book.FindBookByID(server.DB, uint32(bid))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, bookGotten)
}

func (server *Server) UpdateBook(w http.ResponseWriter, r *http.Request) {
	book := models.Book{}
	params := mux.Vars(r)
	bid, err := strconv.Atoi(params["id"])
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	json.NewDecoder(r.Body).Decode(&book)
	tokenID, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	if tokenID != uint32(bid) {
		responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}
	book.Prepare()
	err = book.Validate("update")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	updatedBook, err := book.UpdateABook(server.DB, uint32(bid))
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, updatedBook)
}
