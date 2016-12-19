package service

import (
	"errors"
	"fmt"
)

// NumberService represents the feature: Inserting Number
type NumberService interface {
	Insert(string, int) (string, error)
}

// NumberServiceInst represents NumberService Instance
type NumberServiceInst struct{}

func (NumberServiceInst) Insert(user string, n int) (string, error) {
	if user == "" {
		return "", errors.New("Empty `user`")
	}

	return fmt.Sprintf("%s inserted %d", user, n), nil
}
