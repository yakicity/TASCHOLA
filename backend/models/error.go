package models

type HTTPError struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
	Place string `json:"place"`
}
