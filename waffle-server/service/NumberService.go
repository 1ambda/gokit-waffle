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
	numberRepository NumberRepository
}

func NewNumberService(nr NumberRepository) NumberService {
	return &numberServiceInst{
		numberRepository: nr,
	}
}

func (numberServiceInst) Insert(user string, n int) (string, error) {
	if user == "" {
		return "", errors.New("Empty `user`")
	}

	return fmt.Sprintf("%s inserted %d", user, n), nil
}
