package models

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"errors"
	"io/ioutil"
	"os"
	"time"
	"www-th3-z-xyz/storage"
)

type Paste struct {
	Id int
	UploaderId string
	Filename string
	InsertDate time.Time
}

var pastesPath = "static/pastes/"

var pastesMaxGlobal = 1000
var pastesMaxUploader = 15
// 5Gb / MaxPastes
var pastesMaxSize = int64((1000 * 1000 * 1000 * 5) / pastesMaxGlobal)
// 3 Days
var pastesExpiration = int64(1*60*60*24*3)

func deletePaste(db *sql.DB, paste *Paste) {
	file := pastesPath + paste.Filename
	err := os.Remove(file)

	if err != nil {
		query := `
		DELETE FROM paste p
		WHERE p.id = ?
	`
		storage.PreparedExec(db, query, paste.Id)

		panic(err)
	}

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

func checkLimits(content []byte, uploaderId string) bool {
	// Length check in bytes is intentional
	if int64(len(content)) > pastesMaxSize {
		return false
	}

	pastes := GetPastes()

	if len(pastes) > pastesMaxGlobal {
		return false
	}

	userPastes := 0
	for _, paste := range pastes  {
		if paste.UploaderId == uploaderId {
			userPastes++
		}
	}

	if userPastes > pastesMaxUploader {
		return false
	}

	return true
}

func GetPastes() []Paste {
	var pastes []Paste

	rows := storage.PreparedQuery(
		storage.Db,
		"SELECT id, uploader_id, filename, insert_date FROM paste",
	)
	defer rows.Close()

	for rows.Next() {
		var paste Paste
		var timestamp int64
		err := rows.Scan(
			&paste.Id, &paste.UploaderId, &paste.Filename, &timestamp,
		)
		if err != nil {
			panic(err)
		}

		paste.InsertDate = time.Unix(timestamp, 0)
		pastes = append(pastes, paste)
	}

	return pastes
}

func GetPaste(db *sql.DB, pasteId int64) *Paste {
	query := `
		SELECT
			id,
			uploader_id,
			filename,
			insert_date
		FROM
			paste
		WHERE
			id = ?
	`
	row := storage.PreparedQueryRow(db, query, pasteId)

	var paste Paste
	row.Scan(&paste.Id, &paste.UploaderId, &paste.Filename, &paste.InsertDate)

	return &paste
}

func SearchPaste(db *sql.DB, filename string) *Paste {
	query := `
		SELECT
			id,
			uploader_id,
			filename,
			insert_date
		FROM
			paste
		WHERE
			filename = ?
	`
	row := storage.PreparedQueryRow(db, query, filename)

	var paste Paste
	row.Scan(&paste.Id, &paste.UploaderId, &paste.Filename, &paste.InsertDate)

	return &paste
}


func NewPaste(db *sql.DB, content []byte, uploaderId string) (*Paste, error) {
	// Move these outside of the function, into the handler
	prunePastes(db, GetPastes())

	if !checkLimits(content, uploaderId) {
		return nil, errors.New("paste limit reached")
	}

	h := sha256.New()
	h.Write(content)
	filename := hex.EncodeToString(h.Sum(nil))

	query := `
		INSERT INTO paste (
			uploader_id,
			filename,
			insert_date
		) VALUES (
			?,
			?,
			?
		)
	`

	pasteId, err := storage.PreparedExec(
		db, query, uploaderId, filename, time.Now().Unix(),
	)

	// Hash collided with another paste, return the existing one
	if err != nil {
		return SearchPaste(db, filename), nil
	}

	err = ioutil.WriteFile(pastesPath + filename, content, 0644)
	if err != nil {
		panic(err)
	}


	return GetPaste(db, pasteId), nil
}
