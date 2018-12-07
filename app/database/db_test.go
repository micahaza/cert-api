package database

import (
	"testing"
	"time"
)

func TestCertsDB_AddCert(t *testing.T) {
	cdb := NewCertsDB()
	cert := Certificate{
		ID:        "001",
		Title:     "title",
		CreatedAt: time.Now(),
		Owner:     "user1",
		Year:      2018,
		Note:      "No note yet",
	}
	cdb.AddCert(cert)
}
