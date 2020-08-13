package main

import (
	// "net/http"
	"time"
	"fmt"
	"strconv"
)

type SessionData map[string]interface{}

// start a session or resume an existing one
// save session id in a cookie
// store session data on the file system (one file per session id)

func sessionStart() {
	// todo
	// cookie := http.Cookie{Name: "healthy", Value: "very very very"}
	// http.SetCookie(res, &cookie)
}

func generateSessionId() string {
	curTime := int(time.Now().Unix())
	return fmt.Sprintf("%s", strconv.Itoa(curTime)) // this is terrible, make it better
}
