package service

import (
	"context"

	"github.com/wwkeyboard/sunsetwx/domain"
	"github.com/wwkeyboard/sunsetwx/util"
)

type registerUseCase struct {
	userRepository domain.UserRepository
}

func NewRegisterUseCase(userRepository domain.UserRepository) domain.RegisterUsecase {
	return &registerUseCase{
		userRepository: userRepository,
	}
}

func (rs *registerUseCase) Create(c context.Context, user *domain.User) error {
	return rs.userRepository.CreateUser(user)
}

func (rs *registerUseCase) GetUserByEmail(c context.Context, email string) (domain.User, error) {
	return rs.userRepository.GetUserByEmail(email)
}

func (rs *registerUseCase) CreateAccessToken(user *domain.User, secret string, expiry int) (accessToken string, err error) {
	return util.CreateAccessToken(user, secret, expiry)
}

func (rs *registerUseCase) CreateRefreshToken(user *domain.User, secret string, expiry int) (refreshToken string, err error) {
	return util.CreateRefreshToken(user, secret, expiry)
}
