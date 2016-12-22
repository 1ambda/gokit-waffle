package number

import (
	"errors"
	"fmt"

	. "github.com/1ambda/gokit-waffle/waffle-server/service/common"
)

// NumberService represents the feature: Inserting Number
type NumberService interface {
	Insert(string, int) (string, error)
	Query(string) (int, error)
}

// numberServiceInst represents NumberService Instance
type numberServiceInst struct {
	repository NumberRepository
}

func NewNumberService(r NumberRepository) NumberService {
	return &numberServiceInst{
		repository: r,
	}
}

func (svc numberServiceInst) Insert(u string, n int) (string, error) {
	if u == "" {
		return "", errors.New("Empty `user`")
	}

	user := User(u)
	number := Number(n)
	subs := &Submission{User: user, Number: number}

	svc.repository.Store(subs)
	s, err := svc.repository.Find(user)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s has been inserted: %d", user, s.Number), nil
}

func (svc numberServiceInst) Query(u string) (int, error) {
	if u == "" {
		return 0, errors.New("Empty `user`")
	}

	user := User(u)
	subs, err := svc.repository.Find(user)
	if err != nil {
		return 0, err
	}

	return subs.GetTotal(), err
}
