package dto

// Registration request data from the client.
type RegistrationRequest struct {
	Username       string `json:"username" validate:"required,min=2,max=24" message:"Username is required and must be between 2 and 24 characters"`
	Password       string `json:"password" validate:"required,min=8,max=100" message:"Password is required and must be between 8 and 100 characters"`
	SecretQuestion string `json:"secret_question" validate:"required,min=10,max=100" message:"Secret Question is required and must be between 10 and 100 characters"`
}
