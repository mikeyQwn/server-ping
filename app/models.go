package models

type StatusResponse struct {
	IsOnline bool  `json:"status"`
	Error    error `json:"error,omitempty"`
}
