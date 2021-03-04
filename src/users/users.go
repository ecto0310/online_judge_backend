package users

import (
	"database/sql"
	"fmt"
	"github.com/ecto0310/online_judge_backend/src/db"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

func Register(c echo.Context) error {
	user := User{}
	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, Response{Error: "Invalid data"})
	}
	if len(user.Name) < 3 || len(user.Password) < 8 {
		return c.JSON(http.StatusUnprocessableEntity, Response{Error: "Invalid name or password"})
	}
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response{Error: "Encryption failed"})
	}
	user.Password = ""
	result, err := db.Db.Exec(fmt.Sprintf("INSERT INTO users (name, encrypted_password) VALUES ('%s', '%s')", user.Name, encryptedPassword))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, Response{Error: "Failed to register to DB"})
	}
	user.Id, err = result.LastInsertId()
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, Response{Error: "Failed to register to DB"})
	}
	return c.JSON(http.StatusOK, Response{Success: true, User: user})
}

func Login(c echo.Context) error {
	user := User{}
	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, Response{Error: "Invalid data"})
	}
	encryptedPassword := "$2a$10$xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
	err = db.Db.QueryRow(fmt.Sprintf("SELECT id, encrypted_password FROM users WHERE name = '%s'", user.Name)).Scan(&user.Id, &encryptedPassword)
	if err != sql.ErrNoRows && err != nil {
		return c.JSON(http.StatusUnprocessableEntity, Response{Error: "Failed to access to DB"})
	}
	err = bcrypt.CompareHashAndPassword([]byte(encryptedPassword), []byte(user.Password))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return c.JSON(http.StatusUnprocessableEntity, Response{Error: "Mismatched name or password"})
	}
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, Response{Error: "Encryption failed"})
	}
	user.Password = ""
	cookie := http.Cookie{Name: "session", Value: newSession(user), Expires: time.Now().AddDate(0, 0, 7), Path: "/"}
	c.SetCookie(&cookie)
	return c.JSON(http.StatusOK, Response{Success: true, User: user})
}

func Logout(c echo.Context) error {
	sessionId, err := c.Cookie("session")
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, Response{Error: "Invalid data"})
	}
	destroySession(sessionId.Value)
	return c.JSON(http.StatusOK, Response{Success: true})
}
