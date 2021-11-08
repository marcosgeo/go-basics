package app

import (
	"github.com/marcosgeo/go-basics/mvc/controllers"
)

func mapUrls(){
	router.GET("/users/:user_id", controllers.GetUser)
}