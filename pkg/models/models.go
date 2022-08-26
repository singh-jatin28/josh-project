package models

var StatusData map[string]string

func init() {
	StatusData = make(map[string]string)
}

type PostData struct {
	Websites []string `json:"websites"`
}
