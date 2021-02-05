package controllers

import (
	mysqldb "country-dropdown-gin/connection"
	"country-dropdown-gin/models"
	"fmt"
	"html/template"

	"github.com/gin-gonic/gin"
)

func Viewcountry(c *gin.Context) {
	db := mysqldb.SetupDB()
	countryRows, err := db.Query("SELECT * FROM countries")
	if err != nil {
		fmt.Println(err)
	}
	country := models.Country{}
	res := []models.Country{}
	for countryRows.Next() {
		var id int
		var name string
		err = countryRows.Scan(&id, &name)
		if err != nil {
			panic(err.Error())
		}

		country.Countryid = id
		country.Countryname = name

		res = append(res, country)

	}
	fmt.Println("res:", res)
	t, _ := template.ParseFiles("templates/dropdown.html")
	t.Execute(c.Writer, res)
}
