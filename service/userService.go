package service

import (
	"example/web-service-gin/models"
	"example/web-service-gin/repository"
)

type UserService struct {
	UserRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{UserRepo: userRepo}
}

func (s *UserService) GetUsers() ([]models.User, error) {
	return s.UserRepo.GetUsers()
}

func (s *UserService) GetUserByID(id uint) (*models.User, error) {
	return s.UserRepo.GetUserByID(id)
}

func (s *UserService) CreateUser(user *models.User) error {
	return s.UserRepo.CreateUser(user)
}
