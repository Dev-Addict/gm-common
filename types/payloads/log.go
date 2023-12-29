package payloads

import "github.com/Dev-Addict/gm-common/types"

type LogPayload struct {
	Level   types.LogLevel `json:"level" binding:"required" mapstructure:"level" validate:"required,oneof=info debug warn error fatal"`
	Service types.Service  `json:"service" binding:"required" mapstructure:"service" validate:"required,oneof=auth-service broker-service logger-service"`
	Message string         `json:"message" binding:"required" mapstructure:"message" validate:"required"`
	Data    string         `json:"data" mapstructure:"data"`
}
