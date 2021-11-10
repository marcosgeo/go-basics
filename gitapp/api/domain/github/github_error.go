package github

type ErrorResponse struct {
	Message string `json:"message"`
	DocumentationUrl string `json:"documentation_url"`
	Errors []Error `json:"errors"`
	StatusCode int `json:"status_code"`
}

type Error struct {
	Resource string `json:"resource"`
	Code string `json:"code"`
	Field string `json:"field"`
	Message string `json:"message"`
}