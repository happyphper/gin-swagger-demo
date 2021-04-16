package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/happyphper/gin-swagger-demo/response"
	"net/http"
	"time"
)

type LoginForm struct {
	// 账号
	Username string `json:"username" validate:"required,min=2" example:"happyphper"`
	// 密码
	Password string `json:"password" validate:"required,min=6" example:"123123123"`
}

// Login
// @Summary 登录
// @Description 登录接口，适用于账号、密码登录
// @Tags Auth
// @Accept json
// @Produce json
// @Param username body handlers.LoginForm false "JSON"
// @Success 200 {object} response.Response{code=int,data=response.TokenModel} "desc"
// @Success 200 {object} response.Response{code=int,data=response.TokenModel} "desc"
// @Router /login [post]
func Login(ctx *gin.Context) {
	var form LoginForm
	if err := ctx.ShouldBindJSON(&form); err != nil {
		ctx.JSON(http.StatusOK, response.ValidationRes(err.Error()))
		return
	}

	if form.Username != "happyphper" || form.Password != "123123123" {
		ctx.JSON(http.StatusOK, response.AuthErrorRes())
		return
	}

	fmt.Println(form)

	tokenModel := response.TokenModel{
		Token:     "110123123",
		ExpiresIn: time.Now().Add(time.Hour * 24 * 15).Unix(),
	}

	ctx.JSON(http.StatusOK, tokenModel.ToResponse())
}

type RegisterForm struct {
	// 账号
	Username string `json:"username" validate:"required,min=2" example:"18010021234"`
	// 密码
	Password string `json:"password" validate:"required,min=6" example:"123123123"`
}
// Register
// @Summary 注册
// @Description 适用于新账号注册
// @Tags Auth
// @Accept json
// @Produce json
// @Param username body handlers.RegisterForm false "JSON"
// @Success 200 {object} response.Response{code=int,data=response.TokenModel} "desc"
// @Success 200 {object} response.Response{code=int,data=response.TokenModel} "desc"
// @Router /register [post]
func Register(ctx *gin.Context) {
	var form LoginForm
	if err := ctx.BindJSON(&form); err != nil {
		ctx.JSON(http.StatusOK, response.ValidationRes(err.Error()))
		return
	}

	if form.Username == "happyphper" {
		ctx.JSON(http.StatusOK, response.UsernameExistsRes())
		return
	}

	tokenModel := response.TokenModel{
		Token:        "110123123",
		ExpiresIn: time.Now().Add(time.Hour * 24 * 15).Unix(),
	}

	ctx.JSON(http.StatusOK, tokenModel.ToResponse())
}
