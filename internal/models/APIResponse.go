package models

type APIResponse[T any] struct {
	OK     bool `json:"ok"`
	Result T    `json:"result"`
}
