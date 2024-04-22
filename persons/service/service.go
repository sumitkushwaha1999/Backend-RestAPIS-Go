package service

import (
	"app1/persons/interfaces"
	"app1/persons/reporsitory"
	"app1/requests"
	"app1/responses"
	"context"
	"sync"
)

type personService struct {
	personRepo interfaces.PersonRepo
}

var (
	personServiceObject   sync.Once
	personServiceInstance interfaces.PersonService
)

func personServiceInterface() interfaces.PersonService {
	return &personService{
		personRepo: reporsitory.GetPersonRepoInstance(),
	}
}
func GetPersonServiceInstance() interfaces.PersonService {
	personServiceObject.Do(func() {
		personServiceInstance = personServiceInterface()
	})
	return personServiceInstance
}
func (p *personService) CreatePerson(ctx context.Context, req requests.Person) error {
	return p.personRepo.CreatePerson(ctx, req)
}
func (p *personService) GetPerson(ctx context.Context, ids []int) ([]responses.Person, error) {
	return p.personRepo.GetPerson(ctx, ids)
}
func (p *personService) UpdatePerson(ctx context.Context) error {
	return nil
}
func (p *personService) SofteDeletePerson(ctx context.Context) error {
	return nil
}
