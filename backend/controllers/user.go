package controllers

import (
	"cs-backend/db"
	"cs-backend/structures"
	"fmt"
	"net/http"
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllUsersInAnOrganization(c *gin.Context) {
	id := c.Param("id")
	sortBy := c.Query("sortBy")
	order := c.Query("order")
	page := c.Query("page")
	pageInt, _ := strconv.Atoi(page)

	limit := 8
	offset := pageInt * limit

	r, _ := regexp.Compile("[a-zA-Z]+")

	if id == "" || r.MatchString(id) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid organization Id"})
		return
	}

	query := fmt.Sprintf("SELECT * from users WHERE organization_id = %s", id)
	if sortBy != "" {
		query += fmt.Sprintf(" ORDER BY %s", sortBy)
		if order != "" {
			if order == "ASC" {
				query += " ASC"
			} else if order == "DESC" {
				query += " DESC"
			}
		}
	}
	query += fmt.Sprintf(" LIMIT %d OFFSET %d", limit, offset)

	fmt.Println(query)

	rows, err := db.Db.Query(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var data []structures.User
	for rows.Next() {
		var d structures.User
		err := rows.Scan(&d.Id, &d.FirstName, &d.LastName, &d.Email, &d.Gender, &d.Age, &d.ActiveHoursOnApp, &d.HasUnsubscribed, &d.OrganizationId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		data = append(data, d)
	}
	c.JSON(http.StatusOK, data)
}
