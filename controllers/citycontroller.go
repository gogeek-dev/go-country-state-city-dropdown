package controllers

import (
	mysqldb "country-dropdown-gin/connection"
	"country-dropdown-gin/models"
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Cities(c *gin.Context) {
	db := mysqldb.SetupDB()
	sid := c.Request.PostFormValue("state_id")
	fmt.Println("sid", sid)
	citiesRows, err := db.Query("SELECT * FROM cities where state_id =?", sid)
	if err != nil {
		fmt.Println(err)
	}
	city := models.City{}
	res2 := []models.City{}
	for citiesRows.Next() {
		var id, state_id int
		var name string
		err = citiesRows.Scan(&id, &state_id, &name)
		if err != nil {
			panic(err.Error())
		}

		city.Cityid = id
		city.Stateid = state_id
		city.Cityname = name
		res2 = append(res2, city)

	}
	fmt.Println("res:", res2)

	json.NewEncoder(c.Writer).Encode(res2)

}

func Check(c *gin.Context) {
	pid := c.Request.PostFormValue("product_id")
	cid := c.Request.PostForm.Get("product_id")

	fmt.Println("the check pid is", pid)
	fmt.Println("the check cid is", cid)
	json.NewEncoder(c.Writer).Encode("Hi")
}
