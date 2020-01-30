package models

import (
	"beta-th3-z-xyz/storage"
	"database/sql"
	"time"
	"errors"
)

type Paste struct {
	Id int
	UploaderId int
	Filename string
	InsertDate time.Time
}

var pastesMaxGlobal = 1000
var pastesMaxUploader = 15
// 5Gb / MaxPastes
var pastesMaxSize = int64((1000 * 1000 * 1000 * 5) / pastesMaxGlobal)
// 3 Days
var pastesExpiration = int64(1*60*60*24*3)

func deletePaste(db *sql.DB, paste *Paste) {
	query := `
		DELETE FROM paste p
		WHERE p.id = ?
	`
	storage.PreparedExec(db, query, paste.Id)
}

func prunePastes(db *sql.DB, pastes []Paste) {
	for _, paste := range pastes {
		age := time.Now().Unix() - paste.InsertDate.Unix()
		if age > pastesExpiration {
			deletePaste(db, &paste)
		}
	}
}

func checkLimits(content *string, uploaderId string) bool {
	if len(GetPastes()) > pastesMaxGlobal {
		return false
	}

	return true
}

func GetPastes() []Paste {
	var pastes []Paste

	return pastes
}

func NewPaste(db *sql.DB, content *string, uploaderId string) (*Paste, error) {
	prunePastes(db, GetPastes())

	if !checkLimits(content, uploaderId) {
		return nil, errors.New("paste limit reached")
	}
	var paste Paste

	return &paste, nil
}