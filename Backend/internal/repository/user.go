package repository

import (
	"fmt"

	"github.com/alamgir-ahosain/e-commerce-project/internal/models"
)

func initialUser(r *userRepo) {
	r.userList = []models.User{}
}

//Repository Design Pattern
type UserRepo interface {
	CreateUserFunc(u models.User) (*models.User, error)
	LoginUserFunc(login models.RequestLogin) (*models.User, error)
}

type userRepo struct {
	userList []models.User
}

// Public constructor
func NewUserRepo() UserRepo {
	repo := &userRepo{}
	initialUser(repo)
	return repo
}

// CreateUserFunc implements UserRepo.
func (ur *userRepo) CreateUserFunc(u models.User) (*models.User, error) {
	if u.ID != 0 {
		return &models.User{}, fmt.Errorf("duplicate user")
	}
	u.ID = len(ur.userList) + 1
	ur.userList = append(ur.userList, u)
	return &u, nil
}

// LoginUserFunc implements UserRepo.
func (ur *userRepo) LoginUserFunc(login models.RequestLogin) (*models.User, error) {
	for i, val := range ur.userList {
		if login.Email == "" || login.Password == "" {
			return &models.User{}, fmt.Errorf("email and password required")
		}

		if val.Email == login.Email && val.Password == login.Password {
			return &ur.userList[i], nil
		}
	}
	return &models.User{}, fmt.Errorf("user not found")
}
