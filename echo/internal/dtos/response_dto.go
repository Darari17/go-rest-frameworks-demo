package dtos

type Response[T any] struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
	Data   T      `json:"data,omitempty"`
}
