package config

import "os"

const (
	secretGithubAccessToken = "SECRET_GITHUB_ACCESS_TOKEN"
	LogLevel                = "info"
	goEnvironment           = "GO_ENVIRIONMENT"
	production              = "productions"
)

var (
	githubAccessToken = os.Getenv(secretGithubAccessToken)
)

func GetGithubAccessToken() string {
	return githubAccessToken
}

func IsProduction() bool {
	return os.Getenv(goEnvironment) == production
}
