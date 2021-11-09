package github_provider

import (
	"encoding/json"
	"fmt"
	"github.com/marcosgeo/go-basics/gitapp/api/clients/restclient"
	"github.com/marcosgeo/go-basics/gitapp/api/domain/github"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	headerAuthorization = "Authorization"
	headerAuthorizationFormat = "token %s"
	urlCreateRepo = "https://api.github.com/user/repos"
)

func getAuthorizationHeader(accessToken string) string {
	return fmt.Sprintf(headerAuthorizationFormat, accessToken)
}
func CreateRepo(
	accessToken string, request github.CreateRepoRequest) (*github.CreateRepoResponse, *github.ErrorResponse) {
	headers := http.Header{}
	headers.Set(headerAuthorization, getAuthorizationHeader(accessToken))

	response, err := restclient.Post(urlCreateRepo, request, headers)
	if err != nil {
		log.Println("Error when trying to create new repo in github: %s", err.Error())
		return nil, &github.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, &github.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message: "Invalid response body",
		}
	}
	defer response.Body.Close()

	if response.StatusCode > 299 {
		var errResponse github.ErrorResponse
		if err := json.Unmarshal(bytes, &errResponse); err != nil {
			return nil, &github.ErrorResponse{
				StatusCode: http.StatusInternalServerError,
				Message: "Invalid json response body",
			}
		}
		errResponse.StatuCode = response.StatusCode
		return nil, &errResponse
	}
	return &github.CreateRepoResponse{}, nil
}