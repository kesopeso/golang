package stubs_test

import (
	"ch15/stubs"
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
)

type EntitiesStub struct {
	stubs.Entities
	getUser func(id string) (stubs.User, error)
	getPets func(userId string) ([]stubs.Pet, error)
}

func (es EntitiesStub) GetPets(userId string) ([]stubs.Pet, error) {
	return es.getPets(userId)
}

func TestGetPetNames(t *testing.T) {
	data := []struct {
		userId  string
		names   []string
		errMsg  string
		getPets func(userId string) ([]stubs.Pet, error)
	}{
		{"1", []string{"andrew", "mike", "thomas"}, "", func(userId string) ([]stubs.Pet, error) {
			return []stubs.Pet{{"andrew", 10}, {"mike", 20}, {"thomas", 30}}, nil
		}},
		{"5", nil, "no such user", func(userId string) ([]stubs.Pet, error) {
			return nil, errors.New("no such user")
		}},
	}

	l := stubs.Logic{}
	for _, d := range data {
		l.Entities = EntitiesStub{getPets: d.getPets}
		petNames, err := l.GetPetNames(d.userId)
		var errMsg string
		if err != nil {
			errMsg = err.Error()
		}
		if diff := cmp.Diff(errMsg, d.errMsg); diff != "" {
			t.Error(diff)
		}
		if diff := cmp.Diff(petNames, d.names); diff != "" {
			t.Error(diff)
		}
	}
}
