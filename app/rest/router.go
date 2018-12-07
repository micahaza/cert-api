package rest

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/micahaza/cert-api/app/database"
)

// Server ...
type Server struct {
	db     *database.CertsDB
	router *mux.Router
	// email  EmailSender
}

// NewServer ...
func NewServer() *Server {
	return &Server{
		db:     database.NewCertsDB(),
		router: mux.NewRouter().StrictSlash(true),
	}
}

func (s *Server) getAllCertificates() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// use thing
	}
}

func (s *Server) createCertificate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// use thing
	}
}

func (s *Server) getUserCertificates() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// use thing
	}
}

func (s *Server) deleteCertificate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// use thing
	}
}

func (s *Server) updateCertificate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// use thing
	}
}

func (s *Server) transferCertificate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// use thing
	}
}

func (s *Server) acceptCertificateTransfer() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// use thing
	}
}

// Routes ...
func (s *Server) Routes() {
	s.router.HandleFunc("/certificates/", s.getAllCertificates()).Methods("GET")
	s.router.HandleFunc("/certificates/", s.createCertificate()).Methods("POST")
	s.router.HandleFunc("/users/{id}/certificates", s.getUserCertificates()).Methods("GET")
	s.router.HandleFunc("/certificates/{id}", s.deleteCertificate()).Methods("DELETE")
	s.router.HandleFunc("/certificates/{id}", s.updateCertificate()).Methods("PUT")
	s.router.HandleFunc("/certificates/", s.transferCertificate()).Methods("POST")
	s.router.HandleFunc("/certificates/", s.acceptCertificateTransfer()).Methods("PUT")
}
