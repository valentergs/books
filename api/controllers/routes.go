package controllers

import "github.com/valentergs/booksv2/api/middlewares"

func (s *Server) initializeRoutes() {

	// Home Route
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	// Login Route
	s.Router.HandleFunc("/login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")

	//Users routes
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.GetUsers)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(s.GetUser)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateUser))).Methods("PUT")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteUser)).Methods("DELETE")

	//Books routes
	s.Router.HandleFunc("/books", middlewares.SetMiddlewareJSON(s.CreateBook)).Methods("POST")
	s.Router.HandleFunc("/books", middlewares.SetMiddlewareJSON(s.GetBooks)).Methods("GET")
	s.Router.HandleFunc("/books/{id}", middlewares.SetMiddlewareJSON(s.GetBook)).Methods("GET")
	s.Router.HandleFunc("/books/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateBook))).Methods("PUT")
}
