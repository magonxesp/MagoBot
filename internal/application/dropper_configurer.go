package application

import (
	"errors"
	"github.com/MagonxESP/MagoBot/internal/domain"
)

type DropperConfigurer struct {
	Repository domain.DropperConfigRepository
}

var ExistingConfigError = errors.New("the user has existing dropper configuration")

func NewDropperConfigurer(repository domain.DropperConfigRepository) *DropperConfigurer {
	return &DropperConfigurer{
		Repository: repository,
	}
}

func (dc *DropperConfigurer) CreateNewConfig(config *domain.DropperConfig) error {
	existing, err := dc.Repository.FindByUserId(config.UserId)

	if err != nil {
		return err
	}

	if err == nil && existing != nil {
		return ExistingConfigError
	}

	return dc.Repository.Save(config)
}
