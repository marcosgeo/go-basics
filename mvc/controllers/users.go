package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"

	"github.com/marcosgeo/go-basics/mvc/services"
	"github.com/marcosgeo/go-basics/mvc/utils"
)

func GetUser(c *gin.Context) {
	// get data from URL
	userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)

	// check for errors
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		utils.Respond(c, apiErr)
		return
	}

	// get the user from service and check for errors
	user, apiErr := services.UsersService.GetUser(userId)
	if apiErr != nil {
		utils.Respond(c, apiErr)
		return
	}

	// return the value
	c.JSON(http.StatusOK, user)
}
