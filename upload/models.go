package upload

type File struct {
	Data string `json:"data"`
}

type ResponsePayload struct {
	Total            int `json:"total"`
	Up               int `json:"up"`
	Down             int `json:"down"`
	ElapsedTimeMilli int `json:"elapsedTimeMilli"`
}
