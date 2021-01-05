package dto

// SignIn is dto for sign in handler
type SignIn struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
