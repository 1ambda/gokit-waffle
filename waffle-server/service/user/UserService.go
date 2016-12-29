package user

import (
	. "github.com/1ambda/gokit-waffle/waffle-server/service/common"
	"github.com/1ambda/gokit-waffle/waffle-server/service/number"
)

type UserService interface {
	Users() []User
	// User(string) (int, error)
}

type service struct {
	repository number.NumberRepository
}

func NewUserService(r number.NumberRepository) UserService {
	return &service{repository: r}
}

func (svc *service) Users() []User {
	var users []User

	for _, subs := range svc.repository.FindAll() {
		users = append(users, subs.User)
	}

	return users
}
