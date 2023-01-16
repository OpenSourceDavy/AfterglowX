package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wwkeyboard/sunsetwx/domain"
	"github.com/wwkeyboard/sunsetwx/logs"
	"github.com/wwkeyboard/sunsetwx/settings"
	"golang.org/x/crypto/bcrypt"
)

type LoginController struct {
	LoginUseCase domain.LoginUsecase
}

func (lc *LoginController) Login(c *gin.Context) {
	var request domain.User

	err := c.ShouldBind(&request)
	if err != nil {
		logs.Log.Error("Login ShouldBind error, error message: %s", err)
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	user, _ := lc.LoginUseCase.GetUserByEmail(c, request.Email)
	if user == (domain.User{}) {
		c.JSON(http.StatusNotFound, domain.ErrorResponse{
			Code:    -1,
			Message: "User not found with the given email/phone."})
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)) != nil {
		c.JSON(http.StatusUnauthorized, domain.ErrorResponse{
			Code:    -2,
			Message: "Invalid credentials"})
		return
	}

	accessToken, err := lc.LoginUseCase.CreateAccessToken(&user, settings.JwtSecret, settings.JwtExpHour)
	if err != nil {
		logs.Log.Error("Login CreateAccessToken error, error message: %s", err)
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	refreshToken, err := lc.LoginUseCase.CreateRefreshToken(&user, settings.RefreshTokenSecret, settings.RefreshTokenExpHour)
	if err != nil {
		logs.Log.Error("Login CreateRefreshToken error, error message: %s", err)
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	loginResponse := domain.LoginResponse{
		Code:         1,
		Message:      "success",
		Data:         user.UserID,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	c.JSON(http.StatusOK, loginResponse)
}
