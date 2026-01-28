package service

import (
	"learn-golang-echo/internal/model"
	"learn-golang-echo/internal/repository"
)

type UserService interface {
	GetAll() []model.User
	GetByID(id int) (model.User, bool)
	Create(user model.User) model.User
	Update(id int, user model.User) (model.User, bool)
	Delete(id int) bool
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) GetAll() []model.User {
	return s.repo.GetAll()
}

func (s *userService) GetByID(id int) (model.User, bool) {
	return s.repo.GetByID(id)
}

func (s *userService) Create(user model.User) model.User {
	return s.repo.Create(user)
}

func (s *userService) Update(id int, user model.User) (model.User, bool) {
	return s.repo.Update(id, user)
}

func (s *userService) Delete(id int) bool {
	return s.repo.Delete(id)
}
