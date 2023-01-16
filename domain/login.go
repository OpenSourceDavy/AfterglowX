package domain

import "context"

type LoginRequest struct {
	UserInfo User
}

type LoginResponse struct {
	Code         int    `json:"code"`
	Message      string `json:"message"`
	Data         string `json:"data"`
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type LoginUsecase interface {
	GetUserByEmail(c context.Context, email string) (User, error)
	CreateAccessToken(user *User, secret string, expiry int) (accessToken string, err error)
	CreateRefreshToken(user *User, secret string, expiry int) (refreshToken string, err error)
}
