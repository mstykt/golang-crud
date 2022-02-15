package request

type UserRequest struct {
	Name           string
	Surname        string
	City           string
	IdentityNumber string `json:"identityNumber" validate:"required"`
}
