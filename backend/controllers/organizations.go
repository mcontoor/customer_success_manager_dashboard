package controllers

import (
	"fmt"
	"net/http"
	"time"

	db "cs-backend/db"
	structures "cs-backend/structures"

	"github.com/gin-gonic/gin"
)

type Organization struct {
	Id              int
	Name            string
	Address         string
	CreatedAt       *time.Time `json:"created_at"`
	DealAmount      int        `json:"deal_amount"`
	DaysTillRenewal int        `json:"days_till_renewal"`
}

func GetOrganizations(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "PUT, POST, GET, DELETE, OPTIONS")

	q := c.Query("q")

	sqlQuery := "SELECT * FROM organizations"
	if q != "" {
		s := fmt.Sprintf(" WHERE name ILIKE '%%%s%%'", q)
		sqlQuery += s
	}

	fmt.Println(sqlQuery)

	rows, err := db.Db.Query(sqlQuery)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var data []structures.Organization
	for rows.Next() {
		var d structures.Organization
		err := rows.Scan(&d.Id, &d.Name, &d.Address, &d.CreatedAt, &d.DealAmount, &d.DaysTillRenewal)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		data = append(data, d)
	}

	c.JSON(http.StatusOK, data)
}
