package services

import (
	"errors"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/core/application/interfaces"
	"reflect"
	"testing"
	"time"
)

func TestJWTService_Generate(t *testing.T) {
	tests := map[string]struct {
		input     *interfaces.GenerateInput
		returnErr bool
	}{
		"Filled payload": {
			input: &interfaces.GenerateInput{
				Payload:   map[string]any{},
				ExpiresIn: time.Now().Add(time.Hour),
			},
			returnErr: false,
		},
		"Nil payload": {
			input:     nil,
			returnErr: true,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			svc := NewJWTService(Config{
				SigningKey: "MY_TEST_KEY",
				Issuer:     "unit_test",
			})

			_, err := svc.Generate(test.input)

			if test.returnErr != (err != nil) {
				t.Errorf("unexpected return: %v", err)
			}
		})
	}
}

func TestJWTService_Parse(t *testing.T) {
	tests := map[string]struct {
		tokenGenerator func(svc *JWTService) interfaces.Token
		outputPayload  interfaces.Payload
		outputErr      error
	}{
		"Valid token": {
			tokenGenerator: func(svc *JWTService) interfaces.Token {
				token, _ := svc.Generate(&interfaces.GenerateInput{
					Payload:   interfaces.Payload{"key1": "value1", "key2": true},
					ExpiresIn: time.Now().Add(time.Hour),
				})

				return token
			}, outputPayload: interfaces.Payload{"key1": "value1", "key2": true}, outputErr: nil,
		},
		"Invalid token": {
			tokenGenerator: func(svc *JWTService) interfaces.Token {
				return "random_string"
			}, outputPayload: interfaces.Payload{"key1": "value1", "key2": true}, outputErr: errors.New("invalid token"),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			svc := NewJWTService(Config{
				SigningKey: "MY_TEST_KEY",
				Issuer:     "unit_test",
			})

			payload, err := svc.Parse(test.tokenGenerator(svc))

			if (err == nil) != (test.outputErr == nil) {
				t.Errorf("unexpexted output error, expected: %v, actual: %v", test.outputErr, err)
				return
			}

			if err == nil && !reflect.DeepEqual(payload, test.outputPayload) {
				t.Errorf("unexpexted payload, expected: %v, actual: %v", test.outputPayload, payload)
			}

		})
	}
}
