package github

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateRepoRequestAsJson(t *testing.T){
	request := CreateRepoRequest{
		Name: "golang introduction",
		Description: "a golang introduction repository",
		Homepage: "https://github.com",
		Private: true,
		HasIssues: true,
		HasProjects: true,
		HasWiki: true,
	}

	// Marshal takes an input interface (struct) and attempts to create a valid json string
	bytes, err := json.Marshal(request)

	assert.Nil(t, err)
	assert.NotNil(t, bytes)
	fmt.Println(string(bytes))
}