package service

import (
	"github.com/quanergyO/ab_test_platform/internal/models"
	"github.com/quanergyO/ab_test_platform/internal/repository"
)

type IGenerator interface {
	GeneratePlugin(spec models.ABTestSpec) error
}

type Service struct {
	IGenerator
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		IGenerator: NewGenerator(),
	}
}
