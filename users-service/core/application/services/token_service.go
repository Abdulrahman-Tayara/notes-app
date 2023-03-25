package services

import "time"

type GenerateInput struct {
	Payload   Payload
	ExpiresIn time.Time
}

type Payload map[string]any
type Token string

type ITokenService interface {
	Generate(input *GenerateInput) (Token, error)

	Parse(token Token) (Payload, error)
}
