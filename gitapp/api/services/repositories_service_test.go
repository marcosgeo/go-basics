package services

import (
	"github.com/marcosgeo/go-basics/gitapp/api/clients/restclient"
	"github.com/marcosgeo/go-basics/gitapp/api/domain/repositories"
	"github.com/marcosgeo/go-basics/gitapp/api/utils/errors"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"
	"testing"
)

func TestMain(m *testing.M) {
	restclient.StartMockups()
	os.Exit(m.Run())
}

func TestCreateRepoInvalidInputName(t *testing.T) {
	request := repositories.CreateRepoRequest{}

	result, err := RepositoryService.CreateRepo(request)

	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, err.Status())
	assert.EqualValues(t, "invalid repository name", err.Message())
}

func TestCreateRepoErrorFromGithub(t *testing.T) {
	restclient.FlushMockups()
	restclient.AddMockup(restclient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusUnauthorized,
			Body:       ioutil.NopCloser(strings.NewReader(`{"message": "Requires authentication", "documentation_url": "https://developer.github.com/documentation"}`)),
		},
	})
	request := repositories.CreateRepoRequest{Name: "testing-repo"}

	result, err := RepositoryService.CreateRepo(request)
	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusUnauthorized, err.Status())
	assert.EqualValues(t, "Requires authentication", err.Message())
}

func TestCreateRepoNoError(t *testing.T) {
	restclient.FlushMockups()
	restclient.AddMockup(restclient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       ioutil.NopCloser(strings.NewReader(`{"id":123, "name": "testing", "owner": {"login": "marcosgeo"}}`)),
		},
	})
	request := repositories.CreateRepoRequest{Name: "testing"}

	result, err := RepositoryService.CreateRepo(request)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.EqualValues(t, 123, result.Id)
	assert.EqualValues(t, "testing", result.Name)
	assert.EqualValues(t, "marcosgeo", result.Owner)
}

func TestCreateRepoConcurrentInvalidRequest(t *testing.T) {
	request := repositories.CreateRepoRequest{}

	output := make(chan repositories.CreateRepositoriesResult)

	service := repoService{}
	go service.createRepoConcurrent(request, output)

	result := <-output
	assert.NotNil(t, result)
	assert.Nil(t, result.Response)
	assert.NotNil(t, result.Error)
	assert.EqualValues(t, http.StatusBadRequest, result.Error.Status())
	assert.EqualValues(t, "invalid repository name", result.Error.Message())
}

func TestCreateRepoConcurrentErrorFromGithub(t *testing.T) {
	restclient.FlushMockups()
	restclient.AddMockup(restclient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusUnauthorized,
			Body:       ioutil.NopCloser(strings.NewReader(`{"message": "Requires authentication", "documentation_url": "https://developer.github.com/docs"}`)),
		},
	})

	request := repositories.CreateRepoRequest{Name: "testing"}

	output := make(chan repositories.CreateRepositoriesResult)

	service := repoService{}
	go service.createRepoConcurrent(request, output)

	result := <-output
	assert.NotNil(t, result)
	assert.Nil(t, result.Response)
	assert.NotNil(t, result.Error)
	assert.EqualValues(t, http.StatusUnauthorized, result.Error.Status())
	assert.EqualValues(t, "Requires authentication", result.Error.Message())
}

func TestCreateRepoConcurrentNoError(t *testing.T) {
	restclient.FlushMockups()
	restclient.AddMockup(restclient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       ioutil.NopCloser(strings.NewReader(`{"id": 123, "name": "testing", "owner": {"login": "fulano"}}`)),
		},
	})

	request := repositories.CreateRepoRequest{Name: "testing"}

	output := make(chan repositories.CreateRepositoriesResult)

	service := repoService{}
	go service.createRepoConcurrent(request, output)

	result := <-output
	assert.NotNil(t, result)
	assert.Nil(t, result.Error)
	assert.NotNil(t, result.Response)
	assert.EqualValues(t, 123, result.Response.Id)
	assert.EqualValues(t, "testing", result.Response.Name)
	assert.EqualValues(t, "fulano", result.Response.Owner)
}

func TestHandleRepoResults(t *testing.T) {
	input := make(chan repositories.CreateRepositoriesResult)
	output := make(chan repositories.CreateReposResponse)

	var wg sync.WaitGroup

	service := repoService{}
	go service.handleRepoResults(&wg, input, output)

	wg.Add(1)
	go func() {
		input <- repositories.CreateRepositoriesResult{
			Error: errors.NewBadRequestError("invalid repository name"),
		}
	}()

	wg.Wait()
	close(input)

	result := <-output

	assert.NotNil(t, result)
	assert.EqualValues(t, 0, result.StatusCode)

	assert.EqualValues(t, 1, len(result.Results))
	assert.NotNil(t, result.Results[0].Error)
	assert.EqualValues(t, http.StatusBadRequest, result.Results[0].Error.Status())
	assert.EqualValues(t, "invalid repository name", result.Results[0].Error.Message())
}

func TestCreateReposInvalidRequest(t *testing.T) {
	requests := []repositories.CreateRepoRequest{
		{},
		{Name: ""},
	}

	result, err := RepositoryService.CreateRepos(requests)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.EqualValues(t, http.StatusBadRequest, result.StatusCode)
	assert.EqualValues(t, 2, len(result.Results))

	assert.Nil(t, result.Results[0].Response)
	assert.EqualValues(t, http.StatusBadRequest, result.Results[0].Error.Status())
	assert.EqualValues(t, "invalid repository name", result.Results[0].Error.Message())

	assert.Nil(t, result.Results[1].Response)
	assert.EqualValues(t, http.StatusBadRequest, result.Results[1].Error.Status())
	assert.EqualValues(t, "invalid repository name", result.Results[1].Error.Message())
}

func TestCreateReposOneSuccessOneFail(t *testing.T) {
	restclient.FlushMockups()
	restclient.AddMockup(restclient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       ioutil.NopCloser(strings.NewReader(`{"id": 123, "name": "testing", "owner": {"login": "fulano"}}`)),
		},
	})

	requests := []repositories.CreateRepoRequest{
		{},
		{Name: "testing"},
	}

	result, err := RepositoryService.CreateRepos(requests)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.EqualValues(t, http.StatusPartialContent, result.StatusCode)
	assert.EqualValues(t, 2, len(result.Results))

	for _, res := range result.Results {
		if res.Error != nil {
			assert.EqualValues(t, http.StatusBadRequest, res.Error.Status())
			assert.EqualValues(t, "invalid repository name", res.Error.Message())
			continue
		}

		assert.EqualValues(t, 123, res.Response.Id)
		assert.EqualValues(t, "testing", res.Response.Name)
		assert.EqualValues(t, "fulano", res.Response.Owner)
	}
}

func TestCreateReposAllSuccess(t *testing.T) {
	restclient.FlushMockups()
	restclient.AddMockup(restclient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       ioutil.NopCloser(strings.NewReader(`{"id": 123, "name": "testing", "owner": {"login": "fulano"}}`)),
		},
	})

	requests := []repositories.CreateRepoRequest{
		{Name: "testing"},
		{Name: "testing"},
	}

	result, err := RepositoryService.CreateRepos(requests)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.EqualValues(t, http.StatusCreated, result.StatusCode)
	assert.EqualValues(t, 2, len(result.Results))

	assert.Nil(t, result.Results[0].Error)
	assert.EqualValues(t, 123, result.Results[0].Response.Id)
	assert.EqualValues(t, "testing", result.Results[0].Response.Name)
	assert.EqualValues(t, "fulano", result.Results[0].Response.Owner)

}
