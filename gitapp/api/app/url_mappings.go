package app

import (
	"github.com/marcosgeo/go-basics/gitapp/api/controllers/polo"
	"github.com/marcosgeo/go-basics/gitapp/api/controllers/repositories"
)

func mapUrls() {
	router.GET("/marco", polo.Marco) // liveliness probe
	router.POST("/repository", repositories.CreateRepo)
	router.POST("/repositories", repositories.CreateRepos)
}