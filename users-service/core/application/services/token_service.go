package services

type Payload map[string]any

type Token string

type ITokenService interface {
	Generate(payload Payload) (Token, error)

	Parse(token Token) (Payload, error)
}
