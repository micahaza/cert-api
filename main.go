package main

import (
	"net/http"
	"time"

	"github.com/micahaza/cert-api/app/database"
	"github.com/micahaza/cert-api/app/rest"
)

func main() {

	d := rest.NewServer()
	d.Routes()

	cert := database.Certificate{
		ID:        "001",
		Title:     "title",
		CreatedAt: time.Now(),
		Owner:     "user1",
		Year:      2018,
		Note:      "No note yet",
	}
	d.DB.AddCert(cert)
	cert2 := database.Certificate{
		ID:        "002",
		Title:     "title2",
		CreatedAt: time.Now(),
		Owner:     "user2",
		Year:      2018,
		Note:      "No note yet",
	}
	d.DB.AddCert(cert2)

	http.ListenAndServe(":8888", d.Router)
}
