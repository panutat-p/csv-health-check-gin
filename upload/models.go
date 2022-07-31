package upload

type File struct {
	Data string `json:"data"`
}

// Result is a response payload
type Result struct {
	Total            int `json:"total"`
	Up               int `json:"up"`
	Down             int `json:"down"`
	ElapsedTimeMilli int `json:"elapsedTimeMilli"`
}
