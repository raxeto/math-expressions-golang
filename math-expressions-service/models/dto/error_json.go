package dto

type ErrorJson struct {
	Expression string `json:"expression"`
	Endpoint   string `json:"endpoint"`
	Frequency  uint   `json:"frequency"`
	Type       string `json:"type"`
}
