package core

import "github.com/google/uuid"

type ID uuid.UUID

func NewID() ID {
	return ID(uuid.New())
}

func (i ID) String() string {
	return uuid.UUID(i).String()
}

func ParseSafely(s string) (ID, error) {
	id, err := uuid.Parse(s)

	if err != nil {
		return ID{}, err
	}

	return ID(id), nil
}

func Parse(s string) ID {
	id, _ := ParseSafely(s)

	return id
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
