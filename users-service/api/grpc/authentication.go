package grpc

import (
	"context"
	"github.com/Abdulrahman-Tayara/notes-app/proto/authentication"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/core/application/auth"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/core/application/interfaces"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/injection"
	"google.golang.org/grpc"
	"net/http"
)

type AuthenticationServer struct {
	authentication.UnimplementedAuthenticationServer

	tokenService interfaces.ITokenService
}

func NewAuthenticationServer(tokenService interfaces.ITokenService) *AuthenticationServer {
	return &AuthenticationServer{
		tokenService: tokenService,
	}
}

func (s *AuthenticationServer) ValidateToken(_ context.Context, request *authentication.ValidateTokenRequest) (*authentication.ValidateTokenResponse, error) {
	payload, err := s.tokenService.Parse(interfaces.Token(request.GetToken()))

	if err != nil {
		return &authentication.ValidateTokenResponse{
			StatusCode: http.StatusUnauthorized,
			Message:    err.Error(),
		}, nil
	}

	userPayload := auth.NewUserClaimsPayload(payload)

	if userPayload == nil {
		return &authentication.ValidateTokenResponse{
			StatusCode: http.StatusUnauthorized,
			Message:    "invalid payload",
		}, nil
	}

	return &authentication.ValidateTokenResponse{
		StatusCode: http.StatusOK,
		Message:    http.StatusText(http.StatusOK),
		Claims: &authentication.TokenClaims{
			UserId: userPayload.UserId,
			Email:  userPayload.Email,
		},
	}, nil
}

func RegisterAuthenticationServer(s *grpc.Server) {
	authentication.RegisterAuthenticationServer(s, NewAuthenticationServer(injection.InitTokenService()))
}
