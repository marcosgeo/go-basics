package repositories

import (
	"github.com/gin-gonic/gin"
	"github.com/marcosgeo/go-basics/gitapp/api/domain/repositories"
	"github.com/marcosgeo/go-basics/gitapp/api/services"
	"github.com/marcosgeo/go-basics/gitapp/api/utils/errors"
	"net/http"
)

func CreateRepo(c *gin.Context) {
	var request repositories.CreateRepoRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		apiErr := errors.NewBadRequestError("Invalid json body.")
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	clientId := c.GetHeader("X-Client-Id")
	result, err := services.RepositoryService.CreateRepo(clientId, request)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func CreateRepos(c *gin.Context) {
	var request []repositories.CreateRepoRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		apiErr := errors.NewBadRequestError("Invalid json body.")
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	clientId := c.GetHeader("X-Client-Id")
	result, err := services.RepositoryService.CreateRepos(clientId, request)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(result.StatusCode, result)
}