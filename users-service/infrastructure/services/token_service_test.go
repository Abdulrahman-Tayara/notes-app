package services

import (
	"errors"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/core/application/services"
	"reflect"
	"testing"
	"time"
)

func TestJWTService_Generate(t *testing.T) {
	tests := map[string]struct {
		input     services.Payload
		returnErr bool
	}{
		"Filled payload": {
			input:     map[string]any{},
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
				SigningKey:       "MY_TEST_KEY",
				ExpirationPeriod: time.Hour,
				Issuer:           "unit_test",
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
		tokenGenerator func(svc *JWTService) services.Token
		outputPayload  services.Payload
		outputErr      error
	}{
		"Valid token": {
			tokenGenerator: func(svc *JWTService) services.Token {
				token, _ := svc.Generate(services.Payload{"key1": "value1", "key2": true})

				return token
			}, outputPayload: services.Payload{"key1": "value1", "key2": true}, outputErr: nil,
		},
		"Invalid token": {
			tokenGenerator: func(svc *JWTService) services.Token {
				return "random_string"
			}, outputPayload: services.Payload{"key1": "value1", "key2": true}, outputErr: errors.New("invalid token"),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			svc := NewJWTService(Config{
				SigningKey:       "MY_TEST_KEY",
				ExpirationPeriod: time.Hour,
				Issuer:           "unit_test",
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
