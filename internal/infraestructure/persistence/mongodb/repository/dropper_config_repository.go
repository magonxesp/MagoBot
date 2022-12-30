package repository

import (
	"context"
	"errors"
	"github.com/MagonxESP/MagoBot/internal/domain"
	"github.com/MagonxESP/MagoBot/internal/infraestructure/helpers"
	"github.com/MagonxESP/MagoBot/internal/infraestructure/persistence/mongodb/documents"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const collection = "dropper_config"

type MongoDbDropperConfigRepository struct{}

func NewMongoDbDropperConfigRepository() *MongoDbDropperConfigRepository {
	return &MongoDbDropperConfigRepository{}
}

func findOneByFieldValue(field string, value interface{}) (*documents.DropperConfigDocument, error) {
	result := helpers.GetMongoCollection(collection).FindOne(context.TODO(), bson.D{{field, value}})

	var config *documents.DropperConfigDocument
	err := result.Decode(&config)

	if err != nil && errors.Is(err, mongo.ErrNoDocuments) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return config, nil
}

func (r *MongoDbDropperConfigRepository) FindById(id string) (*domain.DropperConfig, error) {
	config, err := findOneByFieldValue("id", id)

	if err != nil {
		return nil, err
	}

	if config != nil {
		return config.ToDomainStruct(), nil
	}

	return nil, nil
}

func (r *MongoDbDropperConfigRepository) FindByUserId(userId int) (*domain.DropperConfig, error) {
	config, err := findOneByFieldValue("user_id", userId)

	if err != nil {
		return nil, err
	}

	if config != nil {
		return config.ToDomainStruct(), nil
	}

	return nil, nil
}

func (r *MongoDbDropperConfigRepository) Save(config *domain.DropperConfig) error {
	existing, err := findOneByFieldValue("id", config.Id)

	if err != nil {
		return err
	}

	document := documents.NewDropperConfigDocumentFromDomain(config)

	if existing != nil {
		document.DocumentId = existing.DocumentId
		_, err = helpers.GetMongoCollection(collection).UpdateOne(context.Background(), bson.D{{"_id", existing.DocumentId}}, document)
	} else {
		_, err = helpers.GetMongoCollection(collection).InsertOne(context.Background(), document)
	}

	if err != nil {
		return err
	}

	return nil
}

func (r *MongoDbDropperConfigRepository) Delete(config *domain.DropperConfig) error {
	_, err := helpers.GetMongoCollection(collection).DeleteOne(context.Background(), bson.D{{"id", config.Id}})

	if err != nil {
		return err
	}

	return nil
}
