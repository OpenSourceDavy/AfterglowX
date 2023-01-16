package service

import (
	"context"

	"github.com/wwkeyboard/sunsetwx/domain"
	"github.com/wwkeyboard/sunsetwx/util"
)

type refreshTokenUsecase struct {
	userRepository domain.UserRepository
}

func NewRefreshTokenUsecase(userRepository domain.UserRepository) domain.RefreshTokenUsecase {
	return &refreshTokenUsecase{
		userRepository: userRepository,
	}
}

func (rtu *refreshTokenUsecase) GetUser(c context.Context, data map[string]interface{}) (domain.User, error) {
	return rtu.userRepository.GetUser(data)
}

func (rtu *refreshTokenUsecase) CreateAccessToken(user *domain.User, secret string, expiry int) (accessToken string, err error) {
	return util.CreateAccessToken(user, secret, expiry)
}

func (rtu *refreshTokenUsecase) CreateRefreshToken(user *domain.User, secret string, expiry int) (refreshToken string, err error) {
	return util.CreateRefreshToken(user, secret, expiry)
}

func (rtu *refreshTokenUsecase) ExtractIDFromToken(requestToken string, secret string) (string, error) {
	return util.ExtractIDFromToken(requestToken, secret)
}
