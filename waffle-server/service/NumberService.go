package service

import (
	"errors"
	"fmt"
)

// NumberService represents the feature: Inserting Number
type NumberService interface {
	Generate(string, int) (string, error)
}

// NumberServiceInst represents NumberService Instance
type NumberServiceInst struct{}

func (NumberServiceInst) Generate(user string, n int) (string, error) {
	if user == "" {
		return "", errors.New("Empty `user`")
	}

	return fmt.Sprintf("%s generated %d", user, n), nil
}
