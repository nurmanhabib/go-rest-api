package model

type Result struct {
	Code int `json:"code"`
	Data interface{} `json:"data,omitempty"`
	Message string `json:"message"`
}
