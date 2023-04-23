package grpc

import (
	"context"
	"github.com/Abdulrahman-Tayara/notes-app/proto/authentication"
	"google.golang.org/grpc"
)

type AuthenticationServer struct {
	authentication.UnimplementedAuthenticationServer
}

func NewAuthenticationServer() *AuthenticationServer {
	return &AuthenticationServer{}
}

func (s *AuthenticationServer) ValidateToken(context.Context, *authentication.ValidateTokenRequest) (*authentication.ValidateTokenResponse, error) {
	return nil, nil
}

func RegisterAuthenticationServer(s *grpc.Server) {
	authentication.RegisterAuthenticationServer(s, NewAuthenticationServer())
}
