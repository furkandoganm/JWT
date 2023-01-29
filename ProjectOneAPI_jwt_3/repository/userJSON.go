package repository

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"os"
	"projects/ProjectOneAPI_jwt_3/config/model"
)

func GetUserJSON() (model.Users, error) {
	var reqs model.Users
	file, err := os.ReadFile("./file/users.json")
	if err != nil {
		return reqs, err
	}

	err = json.Unmarshal(file, &reqs)
	if err != nil {
		return reqs, err
	}

	return reqs, nil
}

func AddUserJSON(req model.User) (model.User, error) {
	req.Id = primitive.NewObjectID()
	users, err := GetUserJSON()
	if err != nil {
		return req, err
	}

	users.User = append(users.User, req)

	jsonData, err := json.Marshal(users)
	if err != nil {
		return req, err
	}

	err = os.WriteFile("./file/users.json", jsonData, 0644)
	if err != nil {
		return req, err
	}

	return req, nil
}

func UpdateUserJSON(req model.User) (model.User, error) {
	users, err := GetUserJSON()
	if err != nil {
		return req, err
	}

	for i, user := range users.User {
		if req.Id == user.Id {
			users.User[i].UserName = req.UserName
			users.User[i].EMail = req.EMail
			users.User[i].Password = req.Password
			users.User[i].Status = req.Status
			users.User[i].Database = req.Database
			users.User[i].Collections = req.Collections
		}
	}

	jsonData, err := json.Marshal(users)
	if err != nil {
		return req, err
	}

	err = os.WriteFile("./file/users.json", jsonData, 0644)
	if err != nil {
		return req, err
	}

	return req, nil
}

func DeleteUserJSON(id primitive.ObjectID) error {
	var updateUser model.Users
	users, err := GetUserJSON()
	if err != nil {
		return err
	}

	for _, user := range users.User {
		if id != user.Id {
			updateUser.User = append(updateUser.User, user)
		}
	}
	
	jsonData, err := json.Marshal(updateUser)
	if err != nil {
		return err
	}

	err = os.WriteFile("./file/users.json", jsonData, 0644)
	if err != nil {
		return err
	}

	return nil
}
