package repository

import "learn-golang-echo/internal/model"

type UserRepository interface {
	GetAll() []model.User
	GetByID(id int) (model.User, bool)
	Create(user model.User) model.User
	Update(id int, user model.User) (model.User, bool)
	Delete(id int) bool
}

type userRepo struct {
	users  []model.User
	lastID int
}

func NewUserRepository() UserRepository {
	return &userRepo{
		users:  []model.User{},
		lastID: 0,
	}
}

func (r *userRepo) GetAll() []model.User {
	return r.users
}

func (r *userRepo) GetByID(id int) (model.User, bool) {
	for _, u := range r.users {
		if u.ID == id {
			return u, true
		}
	}
	return model.User{}, false
}

func (r *userRepo) Create(user model.User) model.User {
	r.lastID++
	user.ID = r.lastID
	r.users = append(r.users, user)
	return user
}

func (r *userRepo) Update(id int, user model.User) (model.User, bool) {
	for i, u := range r.users {
		if u.ID == id {
			user.ID = id
			r.users[i] = user
			return user, true
		}
	}
	return model.User{}, false
}

func (r *userRepo) Delete(id int) bool {
	for i, u := range r.users {
		if u.ID == id {
			r.users = append(r.users[:i], r.users[i+1:]...)
			return true
		}
	}
	return false
}
