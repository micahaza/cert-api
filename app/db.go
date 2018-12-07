package app

import (
	"errors"
	"sync"
	"time"
)

/*
Representing database => users, certs, transfers
*/

// User ...
type User struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

// Transfer ...
type Transfer struct {
	CertID string `json:"certid"`
	ToUser string `json:"toId"`
}

// Certificate ...
type Certificate struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"createdAt"`
	Owner     string    `json:"owner"`
	Year      int       `json:"year"`
	Note      string    `json:"note"`
}

// CertsDB represents the database
type CertsDB struct {
	mu               sync.RWMutex
	users            map[string]User
	certificates     map[string]Certificate
	pendingTransfers map[string]Transfer
}

// NewCertsDB ...
func NewCertsDB() *CertsDB {
	return &CertsDB{
		users:            map[string]User{},
		certificates:     map[string]Certificate{},
		pendingTransfers: map[string]Transfer{},
	}
}

// Cert returns one cert if exists
func (db *CertsDB) Cert(id string) (Certificate, error) {
	db.mu.Lock()
	defer db.mu.Unlock()
	c, ok := db.certificates[id]
	if !ok {
		return Certificate{}, errors.New("Can't find certificate")
	}
	return c, nil
}

// AddCert adds a certificate, returns error if cert is already exists
func (db *CertsDB) AddCert(crt Certificate) error {
	db.mu.Lock()
	defer db.mu.Unlock()
	if _, ok := db.certificates[crt.ID]; ok {
		return errors.New("Certificate with this id already exists")
	}
	db.certificates[crt.ID] = crt
	return nil
}

// UpdateCert ..
func (db *CertsDB) UpdateCert(crt Certificate) (Certificate, error) {
	db.mu.Lock()
	defer db.mu.Unlock()
	if _, ok := db.certificates[crt.ID]; !ok {
		return Certificate{}, errors.New("Can't update non existing certificate")
	}
	crt.CreatedAt = db.certificates[crt.ID].CreatedAt
	db.certificates[crt.ID] = crt

	return db.certificates[crt.ID], nil
}
