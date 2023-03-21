package core

import "github.com/google/uuid"

type ID string

func NewID() ID {
	return ID(uuid.New().String())
}

func (i ID) String() string {
	return string(i)
}

func ParseSafely(s string) (ID, error) {
	id, err := uuid.Parse(s)

	if err != nil {
		return "", err
	}

	return ID(id.String()), nil
}

type Entity interface {
	GetId() ID
}

type AggregateRoot struct {
	domainEvents []DomainEvent
}

func (r *AggregateRoot) ApplyEvent(event DomainEvent) {
	r.domainEvents = append(r.domainEvents, event)
}

func (r *AggregateRoot) DomainEvents() []DomainEvent {
	return r.domainEvents
}
