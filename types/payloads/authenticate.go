package payloads

type AuthenticatePayload struct {
	Email    string `json:"email" binding:"required" mapstructure:"email" validate:"required,email"`
	Password string `json:"password" binding:"required" mapstructure:"password" validate:"required,password"`
}
