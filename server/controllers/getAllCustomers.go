package controllers

import (
	"fmt"
	"net/http"

	"github.com/Ledan-bot/groom-room/server/db"
	"github.com/Ledan-bot/groom-room/server/models"
	"github.com/gin-gonic/gin"
)

func GetAllCustomers(c *gin.Context) {
	customers, err := getAllCustomers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	c.JSON(http.StatusAccepted, customers)

}

func getAllCustomers() ([]models.Customer, error) {
	db := db.Connection()
	defer db.Close()
	var customers []models.Customer
	sqlStatement := `SELECT * FROM customers`
	rows, err := db.Query(sqlStatement)
	if err != nil {
		return nil, fmt.Errorf(`error happened during SQL query: %v`, err)
	}
	defer rows.Close()
	for rows.Next() {
		var customer models.Customer
		err = rows.Scan(&customer.ID, &customer.FirstName, &customer.LastName, &customer.Email)
		if err != nil {
			return nil, fmt.Errorf(`error happened during converting row into struct: %v`, err)
		}
		customers = append(customers, customer)
	}
	return customers, err
}
