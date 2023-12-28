package payloads

type CreateUserPayload struct {
	Email     string `json:"email" binding:"required" mapstructure:"email" validate:"required,email"`
	FirstName string `json:"firstName" binding:"required" mapstructure:"firstName" validate:"required,name"`
	LastName  string `json:"lastName" binding:"required" mapstructure:"lastName" validate:"required,name"`
	Password  string `json:"password" binding:"required" mapstructure:"password" validate:"required,password"`
}
