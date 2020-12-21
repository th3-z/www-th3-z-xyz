package models

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo"
)

// TODO: Cache
type Session struct {
	LastActivity int64
	Expires      int64
	Values       *map[string]string
	SessionKey   string
	UserId       int64
}

const sessionsPath = "storage/sessions"

const sessionExpiration = 60 * 60 * 24 * 7

func GetSession(c echo.Context) *Session {
	var session *Session

	sessionCookie, err := c.Cookie("sessionKey")
	if err != nil {
		// No session cookie
		session = NewSession(c)
		return session
	}

	h := sha256.New()
	h.Write([]byte(c.Request().UserAgent() + c.RealIP() + sessionCookie.Value))
	sessionFileName := sessionsPath + "/" + hex.EncodeToString(h.Sum(nil))

	data, err := ioutil.ReadFile(sessionFileName)
	if err != nil {
		// No session file, expired or bad key
		session = NewSession(c)
		return session
	}

	if err = json.Unmarshal(data, &session); err != nil {
		panic(err)
	}

	if session.Expires <= time.Now().Unix() {
		// Session has expired
		session = NewSession(c)
		return session
	}

	session.LastActivity = time.Now().Unix()
	session.Expires = time.Now().Unix() + sessionExpiration

	return session
}

func NewSession(c echo.Context) *Session {
	data := make([]byte, 32)
	if _, err := rand.Read(data); err != nil {
		panic(err)
	}

	h := sha256.New()
	h.Write(data)
	sessionKey := hex.EncodeToString(h.Sum(nil))

	session := Session{
		LastActivity: time.Now().Unix(),
		Expires:      time.Now().Unix() + sessionExpiration,
		SessionKey:   sessionKey,
	}

	WriteSession(c, &session)

	cookie := new(http.Cookie)
	cookie.Name = "sessionKey"
	cookie.Value = sessionKey
	cookie.Expires = time.Now().Add(sessionExpiration * time.Second)
	cookie.Path = "/"
	cookie.HttpOnly = false
	cookie.Secure = false
	c.SetCookie(cookie)

	return &session
}

func WriteSession(c echo.Context, session *Session) {
	h := sha256.New()
	h.Write([]byte(c.Request().UserAgent() + c.RealIP() + session.SessionKey))
	sessionFileName := sessionsPath + "/" + hex.EncodeToString(h.Sum(nil))

	sessionJson, err := json.Marshal(session)
	if err != nil {
		panic(err)
	}

	file, err := os.OpenFile(sessionFileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	file.Truncate(0)
	file.Write(sessionJson)
}
