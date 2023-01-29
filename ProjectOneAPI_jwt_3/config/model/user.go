package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Users struct {
	User []User `json:"users"`
}

type User struct {
	Id          primitive.ObjectID `json:"id"`
	CreatorId   primitive.ObjectID `json:"creator_id"`
	UserName    string             `json:"user_name"`
	EMail       string             `json:"e_mail"`
	Password    string             `json:"password"`
	Status      string             `json:"status"`
	Database    string             `json:"database"`
	Collections []string           `json:"collections"`
}

type BasicUser struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}
