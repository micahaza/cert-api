package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/micahaza/cert-api/app/database"
)

// Server ...
type Server struct {
	db     *database.CertsDB
	router *http.ServeMux
	// email  EmailSender
}

// NewServer ...
func NewServer() *Server {
	return &Server{
		db:     database.NewCertsDB(),
		router: http.NewServeMux(),
	}
}

func (s *Server) handleCertificate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET": // list all certificates
			allCerts = s.db.AllCerts()
			console.log(allCerts)
			slcD := []string{"apple", "peach", "pear"}
			json.NewEncoder(w).Encode(slcD)
		case "POST": // create new certificate
			fmt.Println("POST ...")
		default:
			fmt.Println("Method not allowed")
		}
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
	s.router.HandleFunc("/certificates/", s.handleCertificate()) // GET -> list all, POST -> create new
	// s.router.HandleFunc("/certificates/", s.createCertificate())             // .Methods("POST")

	s.router.HandleFunc("/users/{id}/certificates", s.getUserCertificates()) //.Methods("GET")
	// s.router.HandleFunc("/certificates/{id}", s.deleteCertificate())         //.Methods("DELETE")
	s.router.HandleFunc("/certificates/{id}", s.updateCertificate()) //.Methods("PUT")
	// s.router.HandleFunc("/certificates/", s.transferCertificate())       //.Methods("POST")
	// s.router.HandleFunc("/certificates/", s.acceptCertificateTransfer()) //.Methods("PUT")
}
