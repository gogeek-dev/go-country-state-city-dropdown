package routes

import (
	"country-dropdown-gin/controllers"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(db *sql.DB) *gin.Engine {

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")
	r.GET("/", controllers.Viewcountry)
	r.POST("/states", controllers.State)
	r.POST("/cities", controllers.Cities)
	r.POST("/users", controllers.Users)

	return r
}
