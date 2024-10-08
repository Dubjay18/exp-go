package utils

type RegisterUserRequest struct {
	Username string `gorm:"uniqueIndex;not null" json:"username" binding:"required"`
	Password string `gorm:"not null" json:"password" binding:"required,min=8"`
	Email    string `gorm:"uniqueIndex;not null" json:"email" binding:"required"`
}
