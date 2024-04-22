package interfaces

import (
	"app1/requests"
	"app1/responses"
	"context"
)

type PersonService interface {
	CreatePerson(ctx context.Context, req requests.Person) error
	GetPerson(ctx context.Context, ids []int) ([]responses.Person, error)
	UpdatePerson(ctx context.Context) error
	SofteDeletePerson(ctx context.Context) error
}
type PersonRepo interface {
	CreatePerson(ctx context.Context, req requests.Person) error
	GetPerson(ctx context.Context, ids []int) ([]responses.Person, error)
	UpdatePerson(ctx context.Context) error
	SofteDeletePerson(ctx context.Context) error
}
