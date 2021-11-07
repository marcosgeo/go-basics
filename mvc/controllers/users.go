package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/marcosgeo/go-basics/mvc/services"
	"github.com/marcosgeo/go-basics/mvc/utils"
)

func GetUser(resp http.ResponseWriter, req *http.Request) {
	// get data from URL
	userId, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64)

	// check for errors
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		jsonValue, _ := json.Marshal(apiErr)
		resp.WriteHeader(apiErr.StatusCode)
		resp.Write(jsonValue)
		return
	}

	// get the user from service and check for errors
	user, apiErr := services.GetUser(userId)
	if apiErr != nil {
		jsonValue, _ := json.Marshal(apiErr)
		resp.WriteHeader(apiErr.StatusCode)
		resp.Write([]byte(jsonValue))
		return
	}

	// return the value
	jsonValue, _ := json.Marshal(user)
	resp.Write(jsonValue)

}
