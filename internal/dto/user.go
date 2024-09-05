package dto

// RegisterUserRequest is a struct for user registration request
type RegisterUserRequest struct {
	Username string `gorm:"uniqueIndex;not null" json:"username" binding:"required"`
	Password string `gorm:"not null" json:"password" binding:"required,min=8"`
	Email    string `gorm:"uniqueIndex;not null" json:"email" binding:"required"`
}

type RegisterUserResponse struct {
	Message  string `json:"message"`
	Token    string `json:"token"`
	Username string `json:"username"`
	Email    string `json:"email"`
	ID       string `json:"id"`
}
