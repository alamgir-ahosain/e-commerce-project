package user

import (
	"github.com/alamgir-ahosain/e-commerce-project/config"
	"github.com/alamgir-ahosain/e-commerce-project/internal/repository"
)

type Handler struct {
	config   *config.Config
	userRepo repository.UserRepo
}

func NewHandler(config *config.Config, userRepo repository.UserRepo) *Handler {
	return &Handler{
		config:   config,
		userRepo: userRepo,
	}
}
