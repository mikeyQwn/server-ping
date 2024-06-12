package models

type StatusResponse struct {
	IsOnline bool  `json:"is_online"`
	Error    error `json:"error,omitempty"`
}
