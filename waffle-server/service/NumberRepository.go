package service

import (
	"fmt"
	"sync"
)

type User string
type Number int

type Submission struct {
	user   User
	number Number
}

func NewSubmission(user string, n int) *Submission {
	return &Submission{
		user:   User(user),
		number: Number(n),
	}
}

func (s *Submission) Update(n Number) {
	s.number += n
}

type NumberRepository interface {
	Store(s *Submission) error
	Find(u User) (*Submission, error)
	FindAll() []*Submission
}

type numberRepositoryInst struct {
	mtx         sync.RWMutex
	submissions map[User]*Submission
}

func NewNumberRepository() NumberRepository {
	return &numberRepositoryInst{
		submissions: make(map[User]*Submission),
	}
}

func (r *numberRepositoryInst) Store(s *Submission) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	subs, exist := r.submissions[s.user]
	if !exist {
		r.submissions[s.user] = s
	} else {
		subs.Update(s.number)
	}

	return nil
}

func (r *numberRepositoryInst) Find(u User) (*Submission, error) {
	r.mtx.RLock()
	defer r.mtx.Unlock()

	s, exist := r.submissions[u]
	if !exist {
		return nil, fmt.Errorf("Can't find User: %s", u)
	}

	return s, nil
}

func (r *numberRepositoryInst) FindAll() []*Submission {
	r.mtx.RLock()
	defer r.mtx.Unlock()

	ss := make([]*Submission, 0, len(r.submissions))
	for _, s := range r.submissions {
		ss = append(ss, s)
	}

	return ss
}