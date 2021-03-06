package handler

type RegisterRequest struct {
	Fullname string `form:"fullname" binding:"required"`
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required"`
}

type LoginRequest struct {
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required"`
}

type LoginResponseData struct {
	UserID      uint   `json:"user_id"`
	WalletID    string `json:"wallet_id"`
	Fullname    string `json:"fullname"`
	Email       string `json:"email"`
	AccessToken string `json:"access_token"`
}
