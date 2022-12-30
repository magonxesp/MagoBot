package documents

import "go.mongodb.org/mongo-driver/bson/primitive"

type DropperConfig struct {
	Id           string             `bson:"id"`
	DocumentId   primitive.ObjectID `bson:"_id"`
	Url          string             `bson:"url"`
	ClientKey    string             `bson:"client_key"`
	ClientSecret string             `bson:"client_secret"`
	UserId       int                `bson:"user_id"`
}

func NewDropperConfig(id string, url string, clientKey string, clientSecret string, userId int) *DropperConfig {
	return &DropperConfig{
		Id:           id,
		Url:          url,
		ClientKey:    clientKey,
		ClientSecret: clientSecret,
		UserId:       userId,
	}
}
