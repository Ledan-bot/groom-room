package controllers

import (
	"fmt"
	"net/http"

	"github.com/Ledan-bot/groom-room/server/db"
	"github.com/Ledan-bot/groom-room/server/models"
	"github.com/gin-gonic/gin"
)

func GetAllAnimals(c *gin.Context) {
	animals, err := getallAnimals()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(202, animals)
}

func getallAnimals() (animals []models.Animal, err error) {
	db := db.Connection()

	sqlStatement := "SELECT * FROM animals"

	rows, err := db.Query(sqlStatement)
	fmt.Println("Under Query")
	fmt.Println(rows)
	if err != nil {
		return nil, fmt.Errorf(`error happened during SQL query: %v`, err)
	}
	fmt.Println("under Error")
	defer rows.Close()
	for rows.Next() {
		var animal models.Animal
		err = rows.Scan(&animal.ID, &animal.Name, &animal.Breed, &animal.CustomerID)
		if err != nil {
			fmt.Println(err)
			return nil, fmt.Errorf(`error happened during converting to struct: %v`, err)
		}
		animals = append(animals, animal)
	}
	return animals, err
}
