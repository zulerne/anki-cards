package main

import "database/sql"

type User struct{}
type UserRepo interface {
	GetByID(id int) (*User, error)
}

type postgresRepo struct{ db *sql.DB }

func (r *postgresRepo) GetByID(id int) (*User, error) {
	return &User{}, nil
}

type mockRepo struct{}

func (r *mockRepo) GetByID(id int) (*User, error) {
	return &User{}, nil
}

type UserService struct {
	repo UserRepo
}

func NewUserService(repo UserRepo) *UserService {
	return &UserService{repo: repo}
}

func main() {
	// прод
	svc_prod := NewUserService(&postgresRepo{})
	_ = svc_prod
	// тест
	svc_test := NewUserService(&mockRepo{})
	_ = svc_test
}
