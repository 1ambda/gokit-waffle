package service

import (
	"errors"
	"fmt"
)

// NumberService represents the feature: Inserting Number
type NumberService interface {
	Insert(string, int) (string, error)
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
	subs := &Submission{user: user, number: number}

	svc.repository.Store(subs)
	s, err := svc.repository.Find(user)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s has been inserted: %d", user, s.number), nil
}
