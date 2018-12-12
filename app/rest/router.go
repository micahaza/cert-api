package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/micahaza/cert-api/app/database"
)

// Server ...
type Server struct {
	DB     *database.CertsDB
	Router *http.ServeMux
	// email  EmailSender
}

// NewServer ...
func NewServer() *Server {
	return &Server{
		DB:     database.NewCertsDB(),
		Router: http.NewServeMux(),
	}
}

func (s *Server) handleCertificate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET": // list all certificates
			allCerts, _ := s.DB.AllCerts()
			json.NewEncoder(w).Encode(allCerts)
		case "POST": // create new certificate
			fmt.Println("POST ...")
		default:
			fmt.Println("Method not allowed")
		}
	}
}

func (s *Server) handleOneCertificate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "PUT": // update certificate
			allCerts, _ := s.DB.AllCerts()
			json.NewEncoder(w).Encode(allCerts)
		case "DELETE": // delete certificate
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
	s.Router.HandleFunc("/certificates/", s.handleCertificate())             // GET -> list all, POST -> create new
	s.Router.HandleFunc("/users/{id}/certificates", s.getUserCertificates()) //.Methods("GET")
	s.Router.HandleFunc("/certificates/{id}", s.handleOneCertificate())      //.Methods("PUT", "DELETE")
	// s.router.HandleFunc("/certificates/{id}", s.deleteCertificate())         //.Methods("DELETE")
	// s.router.HandleFunc("/certificates/", s.transferCertificate())       //.Methods("POST")
	// s.router.HandleFunc("/certificates/", s.acceptCertificateTransfer()) //.Methods("PUT")
}
