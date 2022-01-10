package controllers

import (
	"fmt"
	"net/http"

	"github.com/Ledan-bot/groom-room/server/db"
	"github.com/Ledan-bot/groom-room/server/models"
	"github.com/gin-gonic/gin"
)

func GetAllAppointments(c *gin.Context) {
	appointments, err := getAllAppointments()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	c.JSON(http.StatusAccepted, appointments)
}

func getAllAppointments() ([]models.Appointment, error) {
	db := db.Connection()

	var appointments []models.Appointment

	sqlStatement := `SELECT * FROM appointments`

	rows, err := db.Query(sqlStatement)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf(`error happened during SQL query: %v`, err)
	}
	defer rows.Close()
	for rows.Next() {
		var appointment models.Appointment

		err = rows.Scan(&appointment.ID, &appointment.CustomerID, &appointment.Date, &appointment.AnimalID)
		if err != nil {
			return nil, fmt.Errorf(`error happened during converting row into struct: %v`, err)
		}
		appointments = append(appointments, appointment)
	}
	return appointments, err
}
