package payloads

type UpdateMePayload struct {
	FirstName string `json:"firstName,omitempty" mapstructure:"firstName,omitempty" validate:"name"`
	LastName  string `json:"lastName,omitempty" mapstructure:"lastName,omitempty" validate:"name"`
}
