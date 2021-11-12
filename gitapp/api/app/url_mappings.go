package app

import (
	"github.com/marcosgeo/go-basics/gitapp/api/controllers/polo"
	"github.com/marcosgeo/go-basics/gitapp/api/controllers/repositories"
)

func mapUrls() {
	router.GET("/marco", polo.Polo) // liveliness probe
	router.POST("/repositories", repositories.CreateRepo)
}