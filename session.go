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

var sessionData map[string]string

const sessionName = "session_id"

func sessionStart(res http.ResponseWriter, req *http.Request) {
	var cookie *http.Cookie = nil
	cookie, err := req.Cookie(sessionName)
	sessionId := ""

	if sessionData == nil {
		sessionData = make(map[string]string)
	}
	if err != nil {
		// session not there, create it
		sessionId = generateSessionId()
		sessionData[sessionName] = sessionId
		createSessionFile(sessionId)
		cookie = &http.Cookie{Name: sessionName, Value: sessionId}
	} else {
		sessionId = cookie.Value
		loadExistingSessionData(sessionId)
	}
	http.SetCookie(res, cookie)
}

func createSessionFile(sessionId string) {
	jsonData, err := json.Marshal(sessionData)
	if err != nil {
		log.Printf("%s", err)
	}
	err = ioutil.WriteFile("./storage/"+sessionId, jsonData, 0644)
	if err != nil {
		log.Printf("%s", err)
	}
}

func loadExistingSessionData(sessionId string) {
	data, err := ioutil.ReadFile("./storage/" + sessionId)
	if err != nil {
		log.Printf("%s", err)
		sessionData[sessionName] = sessionId
		return
	}
	if err = json.Unmarshal(data, &sessionData); err != nil {
		log.Printf("%s", err)
		sessionData[sessionName] = sessionId
	}
}

func generateSessionId() string {
	curTime := int(time.Now().Unix())
	return fmt.Sprintf("%s", strconv.Itoa(curTime)) // this is terrible, make it better
}
