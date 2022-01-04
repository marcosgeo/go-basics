package services

import (
	"fmt"
	"github.com/marcosgeo/go-basics/gitapp/api/config"
	"github.com/marcosgeo/go-basics/gitapp/api/domain/github"
	"github.com/marcosgeo/go-basics/gitapp/api/domain/repositories"
	"github.com/marcosgeo/go-basics/gitapp/api/log"
	github_provider "github.com/marcosgeo/go-basics/gitapp/api/providers/github"
	"github.com/marcosgeo/go-basics/gitapp/api/utils/errors"
	"net/http"
	"sync"
)

type repoService struct{}

type repoServiceInterface interface {
	CreateRepo(clientId string, request repositories.CreateRepoRequest) (*repositories.CreateRepoResponse,
		errors.ApiError)
	CreateRepos(clientId string, request []repositories.CreateRepoRequest) (repositories.CreateReposResponse,
		errors.ApiError)
}

var (
	RepositoryService repoServiceInterface
)

func init() {
	RepositoryService = &repoService{}
}

func (s *repoService) CreateRepo(clientId string, input repositories.CreateRepoRequest) (*repositories.
	CreateRepoResponse, errors.ApiError) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	request := github.CreateRepoRequest{
		Name:        input.Name,
		Description: input.Description,
		Private:     false,
	}

	log.Info("About tho send request to external api", fmt.Sprint("client_id:%s", clientId), "status:pending")
	response, err := github_provider.CreateRepo(config.GetGithubAccessToken(), request)
	if err != nil {
		log.Error("Response obtained from external api", err, fmt.Sprintf("client_id:%s", clientId), "status:error")
		return nil, errors.NewApiError(err.StatusCode, err.Message)
	}
	log.Info("Response obtained from external api", fmt.Sprintf("client_id:%s", clientId), "status:success")

	result := repositories.CreateRepoResponse{
		Id:    response.Id,
		Name:  response.Name,
		Owner: response.Owner.Login,
	}
	return &result, nil
}

func (s *repoService) CreateRepos(clientId string, requests []repositories.CreateRepoRequest) (repositories.
	CreateReposResponse, errors.ApiError) {
	input := make(chan repositories.CreateRepositoriesResult)
	output := make(chan repositories.CreateReposResponse)
	defer close(output)

	var wg sync.WaitGroup
	go s.handleRepoResults(&wg, input, output)

	// 3 requests to process
	for _, current := range requests {
		wg.Add(1)
		go s.createRepoConcurrent(clientId, current, input)
	}

	wg.Wait()
	close(input)

	result := <-output

	successCreations := 0
	for _, current := range result.Results {
		if current.Response != nil {
			successCreations ++
		}
	}
	if successCreations == 0 {
		result.StatusCode = result.Results[0].Error.Status()
	} else if successCreations == len(requests) {
		result.StatusCode = http.StatusCreated
	} else {
		result.StatusCode = http.StatusPartialContent
	}

	return result, nil
}

func (s *repoService) handleRepoResults(wg *sync.WaitGroup, input chan repositories.CreateRepositoriesResult, output chan repositories.CreateReposResponse) {
	var results repositories.CreateReposResponse
	for incomingEvent := range input {
		repoResult := repositories.CreateRepositoriesResult{
			Response: incomingEvent.Response,
			Error: incomingEvent.Error,
		}
		results.Results = append(results.Results, repoResult)
		wg.Done()
	}
	output <- results
}

func (s *repoService) createRepoConcurrent(clientId string, input repositories.CreateRepoRequest,
	output chan repositories.CreateRepositoriesResult) {
	if err := input.Validate(); err != nil {
		output <- repositories.CreateRepositoriesResult{Error: err}
		return
	}

	result, err := s.CreateRepo(clientId, input)
	if err != nil {
		output <- repositories.CreateRepositoriesResult{Error: err}
		return
	}
	output <- repositories.CreateRepositoriesResult{Response: result}
}
