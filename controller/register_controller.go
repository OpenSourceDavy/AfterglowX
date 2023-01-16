package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/wwkeyboard/sunsetwx/domain"
	"github.com/wwkeyboard/sunsetwx/logs"
	"github.com/wwkeyboard/sunsetwx/settings"
	"golang.org/x/crypto/bcrypt"
)

type RegisterController struct {
	RegisterUsecase domain.RegisterUsecase
}

func (rc *RegisterController) Register(c *gin.Context) {
	var request domain.RegisterRequest

	err := c.ShouldBind(&request)
	if err != nil {
		logs.Log.Error("Register ShouldBind error, error message: %s", err)
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	_, err = rc.RegisterUsecase.GetUserByEmail(c, request.UserInfo.Email)
	if err == nil {
		c.JSON(http.StatusConflict, domain.ErrorResponse{Message: "User already exists with the given email."})
		return
	}

	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(request.UserInfo.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		logs.Log.Error("Register bcrypt error, error message: %s", err)
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	user := domain.User{
		UserID:   uuid.NewString(),
		UserName: request.UserInfo.UserName,
		Email:    request.UserInfo.Email,
		Password: string(encryptedPassword),
	}

	err = rc.RegisterUsecase.Create(c, &user)
	if err != nil {
		logs.Log.Error("Register Create error, error message: %s", err)
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	accessToken, err := rc.RegisterUsecase.CreateAccessToken(&user, settings.JwtSecret, settings.JwtExpHour)
	if err != nil {
		logs.Log.Error("Register CreateAccessToken error, error message: %s", err)
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	refreshToken, err := rc.RegisterUsecase.CreateRefreshToken(&user, settings.RefreshTokenSecret, settings.RefreshTokenExpHour)
	if err != nil {
		logs.Log.Error("Register CreateRefreshToken error, error message: %s", err)
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	signupResponse := domain.RegisterResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	c.JSON(http.StatusOK, signupResponse)
}
