package utils

type RegisterUserRequest struct {
	Username string `gorm:"uniqueIndex;not null" json:"username"`
	Password string `gorm:"not null" json:"password"`
	Email    string `gorm:"uniqueIndex;not null" json:"email"`
}
