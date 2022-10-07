package service

import (
	"context"

	"github.com/ValikoDorodnov/go_sample/internal/entity"
	"github.com/ValikoDorodnov/go_sample/internal/repository"
)

type GreetingService struct {
	repo *repository.GreetingRepository
}

func NewGreetingService(repo *repository.GreetingRepository) *GreetingService {
	return &GreetingService{
		repo: repo,
	}
}

func (gs *GreetingService) HelloProcess() (string, error) {
	return "Hello", nil
}

func (gs *GreetingService) HelloDbProcess(ctx context.Context) ([]*entity.GreetingEntity, error) {
	return gs.repo.All(ctx)
}
