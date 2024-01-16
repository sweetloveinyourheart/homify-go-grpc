package handlers

import (
	"context"
	proto "homify-go-grpc/api/authentication"
	"homify-go-grpc/internal/api-gateway/dtos"
	"homify-go-grpc/internal/api-gateway/helpers"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AuthHandler struct {
	grpcClient proto.AuthenticationClient
	validator  *validator.Validate
}

func NewAuthHandler(c proto.AuthenticationClient, validate *validator.Validate) *AuthHandler {
	return &AuthHandler{
		grpcClient: c,
		validator:  validate,
	}
}

// SignUp godoc
// @Tags Authentication
// @Summary Register a new user account
// @Description This endpoint allows users to sign up and create a new account. It expects a JSON payload containing user information.
// If successful, it returns a 201 status code along with the created user data.
// @ID signUp
// @Accept json
// @Produce json
// @Param body body dtos.SignUpDTO true "User information for registration"
// @Success 201 {object} proto.SignUpResponse "Created"
// @Failure 400 {object} interface{} "Bad Request"
// @Router /sign-up [post]
func (h *AuthHandler) SignUp(ctx *gin.Context) {
	// ValidateSignUpDTO validates the SignUpDTO
	signUpData := dtos.SignUpDTO{}
	if bindError := ctx.ShouldBindJSON(&signUpData); bindError != nil {
		ctx.JSON(400, gin.H{"error": bindError.Error()})
		return
	}

	validatorError := h.validator.Struct(signUpData)
	if validatorError != nil {
		helpers.HandleValidationErrors(ctx, validatorError)
		return
	}

	authCtx := context.Background()

	grpcReq := &proto.SignUpRequest{
		Email:    signUpData.Email,
		Password: signUpData.Password,
		FullName: signUpData.FullName,
		Gender:   signUpData.Gender,
		Phone:    signUpData.Phone,
		Birthday: signUpData.Birthday,
	}

	grpcRes, err := h.grpcClient.SignUp(authCtx, grpcReq)
	if err != nil {
		log.Printf("SignUp thrown error: %v", err)
		ctx.JSON(400, gin.H{"error": "Create new account failed"})
		return
	}

	if !grpcRes.Success {
		ctx.JSON(400, gin.H{
			"message": "Failed",
			"data":    grpcRes,
		})
		return
	}

	ctx.JSON(201, gin.H{
		"message": "Created",
		"data":    grpcRes,
	})
}

// SignIn godoc
// @Tags Authentication
// @Summary Handles user sign-in
// @Description Handles the user sign-in process by validating input and authenticating the user via gRPC.
// @Accept json
// @Produce json
// @Param request body dtos.SignInDTO true "User sign-in data in JSON format"
// @Success 200 {object} proto.SignInResponse
// @Failure 400 {object} interface{}
// @Failure 401 {object} interface{}
// @Router /sign-in [post]
func (h *AuthHandler) SignIn(ctx *gin.Context) {
	// ValidateSignInDTO validates the SignInDTO
	signInData := dtos.SignInDTO{}
	if bindError := ctx.ShouldBindJSON(&signInData); bindError != nil {
		ctx.JSON(400, gin.H{"error": bindError.Error()})
		return
	}

	validatorError := h.validator.Struct(signInData)
	if validatorError != nil {
		helpers.HandleValidationErrors(ctx, validatorError)
		return
	}

	authCtx := context.Background()

	grpcReq := &proto.SignInRequest{
		Email:    signInData.Email,
		Password: signInData.Password,
	}

	grpcRes, err := h.grpcClient.SignIn(authCtx, grpcReq)
	if err != nil {
		log.Printf("SignIn thrown error: %v", err)
		ctx.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Created",
		"data":    grpcRes,
	})
}
