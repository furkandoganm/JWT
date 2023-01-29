package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Movie struct {
	Id             primitive.ObjectID `json:"id"`
	Director       string             `json:"director"`
	Scriptwriter   string             `json:"scriptwriter"`
	Name           string             `json:"name"`
	IMDBScore      string             `json:"imdb_score"`
	DatabaseName   string             `json:"database_name"`
	CollectionName string             `json:"collection_name"`
}
