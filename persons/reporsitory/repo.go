package reporsitory

import (
	"app1/db"
	"app1/persons/interfaces"
	"app1/requests"
	"app1/responses"
	"context"
	"fmt"
	"sync"

	"github.com/Masterminds/squirrel"
)

type personRepo struct{}

var (
	personRepoObject   sync.Once
	personRepoInstance interfaces.PersonRepo
)

func personRepoInterface() interfaces.PersonRepo {
	return &personRepo{}
}
func GetPersonRepoInstance() interfaces.PersonRepo {
	personRepoObject.Do(func() {
		personRepoInstance = personRepoInterface()
	})
	return personRepoInstance
}

const (
	PERSON = "personIdentity"
)

func (p *personRepo) CreatePerson(ctx context.Context, req requests.Person) error {
	conn, err := db.DBClient.Conn(ctx)
	if err != nil {
		return err
	}
	defer conn.Close()
	qbuilder := squirrel.Insert(PERSON).Columns("name", "age", "phoneNumber", "status").Values(req.Name, req.Age, req.PhoneNumber, "A")
	query, qargs, err := qbuilder.ToSql()
	if err != nil {
		return err
	}
	_, err = conn.ExecContext(ctx, query, qargs...)
	if err != nil {
		return err
	}
	return nil
}
func (p *personRepo) GetPerson(ctx context.Context, ids []int) ([]responses.Person, error) {
	conn, err := db.DBClient.Conn(ctx)
	if err != nil {
		return nil, err
	}
	fmt.Println(ids)
	defer conn.Close()
	qbuilder := squirrel.Select("id", "name", "age", "phoneNumber", "status").From(PERSON).
		Where(squirrel.Eq{"id": ids})
	query, qargs, err := qbuilder.ToSql()
	if err != nil {
		return nil, err
	}
	rows, err := conn.QueryContext(ctx, query, qargs...)
	if err != nil {
		return nil, err
	}
	var response []responses.Person
	defer rows.Close()
	for rows.Next() {
		var resp responses.Person
		err = rows.Scan(&resp.Id, &resp.Name, &resp.Age, &resp.PhoneNumber, &resp.Status)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(resp)
		response = append(response, resp)
	}
	fmt.Println(response)
	return response, nil
}
func (p *personRepo) UpdatePerson(ctx context.Context) error {
	return nil
}
func (p *personRepo) SofteDeletePerson(ctx context.Context) error {
	return nil
}
