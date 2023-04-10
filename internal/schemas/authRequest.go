package schemas

type AuthRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password"`
}
