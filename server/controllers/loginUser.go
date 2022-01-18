package controllers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/Ledan-bot/groom-room/server/db"
	"github.com/gin-gonic/gin"
)

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func LoginUser(c *gin.Context) {
	var login Login

	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := AuthenticateUser(login.Email, login.Password)
	fmt.Println(err)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error":  err.Error(),
			"status": "Not Authenticated"})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"status": "authenticated"})
}

func AuthenticateUser(email string, password string) (string, error) {
	db := db.Connection()
	defer db.Close()
	fmt.Println(password)
	sqlStatement := fmt.Sprintf("SELECT email, password FROM users WHERE email = '%s'", email)
	var login Login
	err := db.QueryRow(sqlStatement).Scan(&login.Email, &login.Password)
	if err != nil {
		return "", err
	}
	fmt.Println(login.Email, login.Password)
	auth := password == login.Password
	if !auth {
		return "", errors.New("incorrect password")
	}
	return "true", nil
}
