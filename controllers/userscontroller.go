package controllers

import (
	mysqldb "country-dropdown-gin/connection"
	"strconv"

	"github.com/gin-gonic/gin"
	"tawesoft.co.uk/go/dialog"
)

func Users(c *gin.Context) {
	db := mysqldb.SetupDB()

	firstname := c.Request.PostFormValue("firstname")
	lastname := c.Request.PostFormValue("lastname")
	emailid := c.Request.PostFormValue("email")
	cid := c.Request.PostFormValue("cityid")
	cityid, _ := strconv.Atoi(cid)
	user, err := db.Prepare("INSERT INTO users(first_name,last_name,email_id,city_id) VALUES(?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	user.Exec(firstname, lastname, emailid, cityid)
	defer db.Close()
	c.Redirect(301, "/")
	dialog.Alert("Form submitted successfull")
}
