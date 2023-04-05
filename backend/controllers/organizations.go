package controllers

import (
	"fmt"
	"net/http"
	"regexp"

	db "cs-backend/db"
	structures "cs-backend/structures"

	"github.com/gin-gonic/gin"
)

func GetOrganizations(c *gin.Context) {
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

func GetOrgSpecificData(c *gin.Context) {
	id := c.Param("id")

	r, _ := regexp.Compile("[a-zA-Z]+")

	if id == "" || r.MatchString(id) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid organization Id"})
		return
	}

	query := fmt.Sprintf("SELECT * from organizations WHERE id = %s", id)

	row := db.Db.QueryRow(query)
	if row == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "organization not found"})
		return
	}

	w := fmt.Sprintf(`select 
	CONCAT(CAST((CAST(active_hours_on_app AS INT)/50)*50 AS VARCHAR),' - ',CAST((CAST(active_hours_on_app AS INT)/50)*50+49 AS VARCHAR)) as hourrange, 
    COUNT(*)
	from users
	where organization_id = %s
	group by CAST(active_hours_on_app AS INT)/50
	order by CAST(active_hours_on_app AS INT)/50`, id)

	trwo, errr := db.Db.Query(w)
	if errr != nil {
		fmt.Println("ERROR", errr.Error())
		return
	}

	var data []structures.HoursOnProduct
	for trwo.Next() {
		var d structures.HoursOnProduct
		er := trwo.Scan(&d.Hourrange, &d.Count)
		if er != nil {
			fmt.Println("ERRR", er.Error())
			return
		}
		data = append(data, d)
	}
	fmt.Println(data)

	var d structures.OrganizationData
	d.HoursOnProductHistogramData = data
	err := row.Scan(&d.Id, &d.Name, &d.Address, &d.CreatedAt, &d.DealAmount, &d.DaysTillRenewal)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, d)
}
