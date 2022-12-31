package application

import (
	"errors"
	"github.com/MagonxESP/MagoBot/internal/domain"
)

type DropperConfigDeleter struct {
	Repository domain.DropperConfigRepository
}

var NotConfigExistsError = errors.New("missing dropper user config")

func NewDropperConfigDeleter(repository domain.DropperConfigRepository) *DropperConfigDeleter {
	return &DropperConfigDeleter{
		Repository: repository,
	}
}

func (dc *DropperConfigDeleter) DeleteUserConfig(userId int) error {
	config, err := dc.Repository.FindByUserId(userId)

	if err != nil {
		return err
	}

	if config == nil {
		return NotConfigExistsError
	}

	return dc.Repository.Delete(config)
}
