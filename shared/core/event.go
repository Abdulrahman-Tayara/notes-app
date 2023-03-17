package core

import "time"

type DomainEvent interface {
	CreatedAt() time.Time
	Identifier() string
}
