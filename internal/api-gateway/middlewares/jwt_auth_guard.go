package middlewares

import (
	"context"
	"fmt"
	proto "homify-go-grpc/api/authentication"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type AuthenticatedUser struct {
	Id    uint
	Email string
	Role  string
}

type JwtAuthGuard struct {
	client proto.AuthenticationClient
}

func NewJwtAuthGuard(c proto.AuthenticationClient) *JwtAuthGuard {
	return &JwtAuthGuard{
		client: c,
	}
}

func extractBearerToken(r *http.Request) (string, error) {
	// Get the Authorization header
	authHeader := r.Header.Get("Authorization")

	if authHeader == "" {
		return "", fmt.Errorf("missing Authorization header")
	}

	// Check if the header has the "Bearer" prefix
	if !strings.HasPrefix(authHeader, "Bearer ") {
		return "", fmt.Errorf("invalid Authorization header format")
	}

	// Extract the token excluding the "Bearer " prefix
	token := strings.TrimPrefix(authHeader, "Bearer ")

	return token, nil
}

func (ag *JwtAuthGuard) AuthGuard(ctx *gin.Context) {
	token, err := extractBearerToken(ctx.Request)

	if err != nil {
		ctx.AbortWithStatusJSON(401, gin.H{"message": "unauthorized"})
		return
	}

	clientCtx := context.Background()
	grpcReq := &proto.VerifyJwtTokenRequest{
		Token: token,
	}

	grpcRes, err := ag.client.VerifyJwtToken(clientCtx, grpcReq)

	if err != nil {
		ctx.AbortWithStatusJSON(401, gin.H{"message": "unauthorized"})
		return
	}

	user := AuthenticatedUser{
		Id:    uint(grpcRes.UserId),
		Email: grpcRes.Email,
		Role:  grpcRes.Role,
	}

	ctx.Set("user", user)
	ctx.Next()
}
