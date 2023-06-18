package middlewares

import (
	"context"
	"errors"
	pb "github.com/Abdulrahman-Tayara/notes-app/proto/authentication"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
	"os"
)

type AuthenticationConfig struct {
	AuthServiceUrl string
}

func Authentication(config *AuthenticationConfig) (gin.HandlerFunc, error) {
	url := config.AuthServiceUrl

	if url == "" {
		url = os.Getenv("AUTH_SERVICE_ADDRESS")
	}

	if url == "" {
		return nil, errors.New("invalid auth service address")
	}

	conn, err := grpc.Dial(config.AuthServiceUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := pb.NewAuthenticationClient(conn)

	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")

		if token == "" {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		response, err := client.ValidateToken(context.TODO(), &pb.ValidateTokenRequest{})
		if err != nil {
			return
		}

		if response.GetStatusCode() >= 300 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, map[string]any{"message": response.GetMessage()})
			return
		}

		claims := response.GetClaims()

		ctx.Set("user_id", claims.GetUserId())
		ctx.Set("email", claims.GetEmail())
	}, nil
}
