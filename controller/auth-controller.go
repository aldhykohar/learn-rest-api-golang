package controller

import (
	"golang_rest_api/dto"
	"golang_rest_api/entity"
	"golang_rest_api/helper"
	"golang_rest_api/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// AuthController interface is a contract what this controller can do
type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type authController struct {
	// disini masukkan service yang kalian butuh
	// this is where you put your service
	authService service.AuthService
	jwtService  service.JWTService
}

//NewAuthController creates a new instance of AuthController
func NewAuthController(authService service.AuthService, jwtService service.JWTService) AuthController {
	return &authController{
		authService: authService,
		jwtService:  jwtService,
	}
}

func (c *authController) Login(ctx *gin.Context) {
	var logiDTO dto.LoginDTO
	errDTO := ctx.ShouldBind(&logiDTO)
	if errDTO != nil {
		response := helper.BuildeErrorResponse("Failed to proces request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	authResult := c.authService.VerifyCredential(logiDTO.Email, logiDTO.Password)
	if v, ok := authResult.(entity.User); ok {
		generateToken := c.jwtService.GenerateToken(strconv.FormatInt(int64(v.ID), 10))
		v.Token = generateToken
		response := helper.BuildResponse(true, "OK1", v)
		ctx.JSON(http.StatusOK, response)
		return
	}

	response := helper.BuildeErrorResponse("Please check check your credential again", "Invalid Credential", helper.EmptyObj{})
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
}

func (c *authController) Register(ctx *gin.Context) {
	var registerDTO dto.RegisterDTO

	errDTO := ctx.ShouldBind(&registerDTO)
	if errDTO != nil {
		response := helper.BuildeErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	if !c.authService.IsDuplicateEmail(registerDTO.Email) {
		response := helper.BuildeErrorResponse("Failed to process request", "Duplicate email", helper.EmptyObj{})
		ctx.JSON(http.StatusConflict, response)
	} else {
		createdUser := c.authService.CreateUser(registerDTO)
		token := c.jwtService.GenerateToken(strconv.FormatUint(createdUser.ID, 10))
		createdUser.Token = token
		response := helper.BuildResponse(true, "Ok!", createdUser)
		ctx.JSON(http.StatusCreated, response)
	}
}
