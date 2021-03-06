package session

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

type User struct {
	Id       int64
	Name     string
	LastName string
	Email    string
}

const defaultSessionCookieName = "session_id"

type Session struct {
	data       map[string]interface{}
	id         string
	cookieName string
}

func (s *Session) Start(res http.ResponseWriter, req *http.Request) {
	if s.cookieName == "" {
		s.cookieName = defaultSessionCookieName
	}
	var cookie *http.Cookie = nil
	cookie, err := req.Cookie(s.cookieName)
	sessionId := ""

	if s.data == nil {
		s.data = make(map[string]interface{})
	}
	if err != nil {
		// session not there, create it
		sessionId = generateSessionId()
		s.data[s.cookieName] = sessionId
		s.id = sessionId
		s.updateFile()
		cookie = &http.Cookie{
			Name:     s.cookieName,
			Value:    s.id,
			Secure:   true,
			HttpOnly: true,
			SameSite: http.SameSiteStrictMode,
		}
	} else {
		cookie.Secure = true
		cookie.HttpOnly = true
		cookie.SameSite = http.SameSiteStrictMode
		sessionId = cookie.Value
		s.id = sessionId
		s.loadData()
	}
	http.SetCookie(res, cookie)
}

func (s *Session) Get(key string) (interface{}, bool) {
	value, ok := s.data[key]
	if !ok {
		return nil, false
	}
	return value, true
}

func (s *Session) Set(key string, value interface{}) {
	s.data[key] = value
	s.updateFile()
}

func (s *Session) Remove(key string) {
	if _, ok := s.Get(key); ok {
		delete(s.data, key)
		s.updateFile()
	}
}

func (s *Session) Destroy(res http.ResponseWriter) error {
	if err := os.Remove("./storage/" + s.id); err != nil {
		return err
	}
	cookie := &http.Cookie{
		Name:     s.cookieName,
		Value:    "",
		Expires:  time.Now(),
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	}
	http.SetCookie(res, cookie)
	s.data = nil
	s.id = ""
	s.cookieName = ""
	return nil
}

func (s *Session) GetUserFromSession() *User {
	var user *User = nil
	if usr, ok := s.Get("user"); ok {
		if usrMap, ok := usr.(map[string]interface{}); ok {
			user = &User{
				Id:       int64(usrMap["Id"].(float64)),
				Name:     usrMap["Name"].(string),
				LastName: usrMap["LastName"].(string),
				Email:    usrMap["Email"].(string),
			}
		}
	}
	return user
}

func (s *Session) LoggedIn() bool {
	_, ok := s.Get("user")
	if !ok {
		return false
	}
	// todo: use the user node here
	return true
}

func (s *Session) updateFile() {
	jsonData, err := json.Marshal(s.data)
	if err != nil {
		log.Printf("%s", err)
	}
	err = ioutil.WriteFile("./storage/"+s.id, jsonData, 0644)
	if err != nil {
		log.Printf("%s", err)
	}
}

func (s *Session) loadData() {
	data, err := ioutil.ReadFile("./storage/" + s.id)
	if err != nil {
		log.Printf("%s", err)
		s.data[s.cookieName] = s.id
		return
	}
	if err = json.Unmarshal(data, &s.data); err != nil {
		log.Printf("%s", err)
		s.data[s.cookieName] = s.id
	}
}

func generateSessionId() string {
	rand.Seed(time.Now().UnixNano())
	chars := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	idLen := 16
	id := make([]byte, idLen)
	for i := 0; i < idLen; i++ {
		id[i] = chars[rand.Intn(len(chars))]
	}
	return string(id)
}
