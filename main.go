package main

import (
	"fmt"
	"time"

	"github.com/micahaza/cert-api/app/database"
	"github.com/micahaza/cert-api/app/rest"
)

func main() {
	cdb := database.NewCertsDB()
	cert := database.Certificate{
		ID:        "001",
		Title:     "title",
		CreatedAt: time.Now(),
		Owner:     "user1",
		Year:      2018,
		Note:      "No note yet",
	}
	cdb.AddCert(cert)

	cert, err := cdb.Cert("001")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(cert)

	upd := database.Certificate{
		ID:        "002",
		Title:     "Changed title",
		CreatedAt: time.Now(),
		Owner:     "user1",
		Year:      2016,
		Note:      "Something important",
	}
	c, err := cdb.UpdateCert(upd)
	if err != nil {
		fmt.Println(c, err)
	}
	cert, err = cdb.Cert("001")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("updated", cert)

	d := rest.NewServer()
	d.Routes()
	fmt.Println(d)
}
