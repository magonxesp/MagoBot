package application

import "github.com/MagonxESP/MagoBot/internal/domain"

type DropperConfigFinder struct {
	Repository domain.DropperConfigRepository
}

func NewDropperConfigFinder(repository domain.DropperConfigRepository) *DropperConfigFinder {
	return &DropperConfigFinder{
		Repository: repository,
	}
}

func (dc *DropperConfigFinder) FindByUserId(userId int64) (*domain.DropperConfig, error) {
	return dc.Repository.FindByUserId(userId)
}
