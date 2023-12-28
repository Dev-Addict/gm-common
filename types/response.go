package types

type Response struct {
	Message string `json:"message"`
	Status  Status `json:"status"`
	Data    any    `json:"data,omitempty"`
}
