package users

import (
	"fmt"
	"github.com/ecto0310/online_judge_backend/src/db"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
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
