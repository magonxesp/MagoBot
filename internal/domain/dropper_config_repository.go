package domain

type DropperConfigRepository interface {
	FindById(id string) (*DropperConfig, error)
	FindByUserId(userId int) (*DropperConfig, error)
	Save(config *DropperConfig) error
	Delete(config *DropperConfig) error
}
