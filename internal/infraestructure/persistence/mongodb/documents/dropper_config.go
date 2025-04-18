package documents

import (
	"github.com/MagonxESP/MagoBot/internal/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DropperConfigDocument struct {
	Id           string             `bson:"id"`
	DocumentId   primitive.ObjectID `bson:"_id"`
	Url          string             `bson:"url"`
	ClientId     string             `bson:"client_id"`
	ClientSecret string             `bson:"client_secret"`
	UserId       int64              `bson:"user_id"`
}

func NewDropperConfigDocument(id string, url string, clientId string, clientSecret string, userId int64) *DropperConfigDocument {
	return &DropperConfigDocument{
		Id:           id,
		DocumentId:   primitive.NewObjectID(),
		Url:          url,
		ClientId:     clientId,
		ClientSecret: clientSecret,
		UserId:       userId,
	}
}

func NewDropperConfigDocumentFromDomain(config *domain.DropperConfig) *DropperConfigDocument {
	return &DropperConfigDocument{
		Id:           config.Id,
		DocumentId:   primitive.NewObjectID(),
		Url:          config.Url,
		ClientId:     config.ClientId,
		ClientSecret: config.ClientSecret,
		UserId:       config.UserId,
	}
}

func (d *DropperConfigDocument) ToDomainStruct() *domain.DropperConfig {
	return &domain.DropperConfig{
		Id:           d.Id,
		Url:          d.Url,
		ClientId:     d.ClientId,
		ClientSecret: d.ClientSecret,
		UserId:       d.UserId,
	}
}
