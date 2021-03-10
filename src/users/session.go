package users

import (
	"github.com/google/uuid"
)

var sessionDic = map[string]User{}

func newSession(user User) string {
	sessionID := uuid.New().String()
	sessionDic[sessionID] = user
	return sessionID
}

func destroySession(sessionId string) {
	sessionDic[sessionId] = User{Id: -1}
}

func CheckSession(user User, sessionId string) bool {
	return sessionDic[sessionId] == user
}
