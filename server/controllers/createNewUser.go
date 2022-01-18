package controllers

import (
	"errors"
	"net/http"

	"github.com/Ledan-bot/groom-room/server/db"
	"github.com/Ledan-bot/groom-room/server/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func CreateNewUser(c *gin.Context) {
	var validate = validator.New()
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	vaildationErr := validate.Struct(user)
	if vaildationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": vaildationErr.Error()})
		return
	}

	InsertId, err := insertuser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"ID":      InsertId,
		"Message": "New user Created",
	})
}

func insertuser(user models.User) (int64, error) {
	db := db.Connection()
	defer db.Close()
	sqlStatement := `INSERT INTO users (firstname, lastname, email) VALUES ($1, $2, $3) RETURNING userid`
	var id int64
	err := db.QueryRow(sqlStatement, user.FirstName, user.LastName, user.Email).Scan(&id)
	if err != nil {
		return 0, errors.New("unable to execute insertion into DB")
	}
	return id, nil
}
