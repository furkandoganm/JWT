package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"projects/ProjectOneAPI_jwt_3/config/model"
	"time"
)

func (dbC *Client) GetMovies() ([]model.Movie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	result, err := dbC.Collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var req model.Movie
	var reqs []model.Movie

	for result.Next(ctx) {
		if err = result.Decode(&req); err != nil {
			return nil, err
		}
		reqs = append(reqs, req)
	}
	return reqs, nil
}

func (dbC *Client) AddMovie(req model.Movie) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	req.Id = primitive.NewObjectID()
	_, err := dbC.Collection.InsertOne(ctx, req)
	if err != nil {
		return err
	}
	return nil
}