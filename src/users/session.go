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
