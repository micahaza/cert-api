package main

import (
	"fmt"
	"time"

	"github.com/micahaza/cert-api/app"
)

func main() {
	cdb := app.NewCertsDB()
	cert := app.Certificate{
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

	upd := app.Certificate{
		ID:        "00",
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
}
