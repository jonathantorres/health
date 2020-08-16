package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

const defaultSessionCookieName = "session_id"

type Session struct {
	data       map[string]string
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
		s.data = make(map[string]string)
	}
	if err != nil {
		// session not there, create it
		sessionId = generateSessionId()
		s.data[s.cookieName] = sessionId
		s.id = sessionId
		s.createFile()
		cookie = &http.Cookie{Name: s.cookieName, Value: s.id}
	} else {
		sessionId = cookie.Value
		s.id = sessionId
		s.loadData()
	}
	http.SetCookie(res, cookie)
}

func (s *Session) Get(key string) (string, bool) {
	value, ok := s.data[key]
	if !ok {
		return "", false
	}
	return value, true
}

func (s *Session) createFile() {
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
	curTime := int(time.Now().Unix())
	return fmt.Sprintf("%s", strconv.Itoa(curTime)) // this is terrible, make it better
}
