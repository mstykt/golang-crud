package response

import "time"

type UserResponse struct {
	ID             uint64    `json:"id"`
	Name           string    `json:"name"`
	Surname        string    `json:"surname"`
	City           string    `json:"city"`
	IdentityNumber string    `json:"identityNumber"`
	CreatedDate    time.Time `json:"createdDate"`
	UpdatedDate    time.Time `json:"updatedDate"`
}
