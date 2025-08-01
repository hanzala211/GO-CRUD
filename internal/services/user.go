package services

import (
	"context"
	"time"

	"github.com/hanzala211/CRUD/internal/api/models"
	"github.com/hanzala211/CRUD/internal/repo"
)

type UserService struct {
	userRepo *repo.UserRepo
}

func NewUserService(userRepo *repo.UserRepo) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) CreateUser(ctx context.Context, user *models.User) error {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	return s.userRepo.CreateUser(ctx, user)
}

func (s *UserService) UpdateUser(ctx context.Context, user *models.User, userId string) error {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	return s.userRepo.UpdateUser(ctx, user, userId)
}

func (s *UserService) DeleteUser(ctx context.Context, userId string) error {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	return s.userRepo.DeleteUser(ctx, userId)
}

func (s *UserService) GetAllUsers(ctx context.Context) ([]*models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	return s.userRepo.GetUsers(ctx)
}

func (s *UserService) GetUserByID(ctx context.Context, user *models.User, userId string) error {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	return s.userRepo.GetUserByID(ctx, user, userId)
}
