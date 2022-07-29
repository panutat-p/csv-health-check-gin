package upload

type File struct {
	Data string `json:"data"`
}

type Result struct {
	Url        string `json:"url"`
	StatusCode int    `json:"statusCode"`
	Error      string `json:"error"`
}
