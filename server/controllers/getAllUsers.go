package controllers

import (
	"fmt"
	"net/http"

	"github.com/Ledan-bot/groom-room/server/db"
	"github.com/Ledan-bot/groom-room/server/models"
	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {
	users, err := getAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, users)
}

func getAllUsers() (users []models.User, err error) {
	db := db.Connection()
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf(`error happened during SQL query: %v`, err)
	}
	defer rows.Close()
	for rows.Next() {
		var user models.User
		err = rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password)
		if err != nil {
			return nil, fmt.Errorf(`error happened during converting row into struct: %v`, err)
		}
		users = append(users, user)
	}
	return users, err
}
