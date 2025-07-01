package service

import "github.com/quanergyO/ab_test_platform/internal/repository"

type Service struct {
}

func NewService(repo *repository.Repository) *Service {
	return &Service{}
}
