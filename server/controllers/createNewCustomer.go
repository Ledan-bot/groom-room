package controllers

import (
	"errors"
	"net/http"

	"github.com/Ledan-bot/groom-room/server/db"
	"github.com/Ledan-bot/groom-room/server/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func CreateNewCustomer(c *gin.Context) {
	var validate = validator.New()
	var customer models.Customer
	if err := c.BindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	vaildationErr := validate.Struct(customer)
	if vaildationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": vaildationErr.Error()})
		return
	}

	InsertId, err := insertCustomer(customer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusCreated, gin.H{
		"ID":      InsertId,
		"Message": "New Customer Created",
	})
}

func insertCustomer(customer models.Customer) (int64, error) {
	db := db.Connection()
	defer db.Close()
	sqlStatement := `INSERT INTO customers (firstname, lastname, email) VALUES ($1, $2, $3) RETURNING customerid`
	var id int64
	err := db.QueryRow(sqlStatement, customer.FirstName, customer.LastName, customer.Email).Scan(&id)
	if err != nil {
		return 0, errors.New("unable to execute insertion into DB")
	}
	return id, nil
}
