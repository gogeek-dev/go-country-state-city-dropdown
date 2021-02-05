package controllers

import (
	mysqldb "country-dropdown-gin/connection"
	"country-dropdown-gin/models"
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
)

func State(c *gin.Context) {
	db := mysqldb.SetupDB()
	cid := c.Request.PostFormValue("country_id")

	stateRows, err := db.Query("SELECT * FROM states where country_id = ?", cid)
	if err != nil {
		fmt.Println(err)
	}
	state := models.States{}
	res1 := []models.States{}

	for stateRows.Next() {
		var id, country_id int
		var name string
		err = stateRows.Scan(&id, &country_id, &name)
		if err != nil {
			panic(err.Error())
		}

		state.Stateid = id
		state.Countryid = country_id
		state.Statename = name

		res1 = append(res1, state)

	}
	fmt.Println("res1:", res1)

	json.NewEncoder(c.Writer).Encode(res1)

}
