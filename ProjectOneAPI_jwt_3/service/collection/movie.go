package collection

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"projects/ProjectOneAPI_jwt_3/config"
	"projects/ProjectOneAPI_jwt_3/config/model"
	"projects/ProjectOneAPI_jwt_3/repository"
	"projects/ProjectOneAPI_jwt_3/service"
)

func Client(dbName, cName string) repository.Client {
	dbClient, err := config.ConnectionDB(dbName, cName, false)
	if err != nil {
		fmt.Println("connecting database")
	}
	return repository.Client{Collection: dbClient}
}

func GetMovie(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	service.EError(&w, "Body is not reading", err)

	var req model.Database
	err = json.Unmarshal(body, &req)

	users, err := repository.GetUserJSON()
	service.EError(&w, "Getting users from JSON file", err)

	userName := service.AUser
	fmt.Printf("\n\n%v\n\n", userName)
	var user model.User
	for _, u := range users.User {
		if u.UserName == userName {
			user.Id = u.Id
			user.UserName = u.UserName
			user.Database = u.Database
			user.Collections = u.Collections
			user.Status = u.Status
		}
	}
	chacking := false
	if user.Status == "godmin" {
		for _, lUser := range users.User {
			if req.DatabaseName == lUser.Database {
				for _, c := range lUser.Collections {
					if req.CollectionName == c {
						chacking = true
					}
				}
			}
		}
	} else if user.Status == "admin" {
		for _, lUser := range users.User {
			if user.Id == lUser.CreatorId || user.Id == lUser.Id {
				if req.DatabaseName == lUser.Database {
					for _, c := range lUser.Collections {
						if req.CollectionName == c {
							chacking = true
						}
					}
				}
			}
		}
	} else if user.Status == "user" {
		for _, lUser := range users.User {
			if user.Id == lUser.Id {
				if req.DatabaseName == lUser.Database {
					for _, c := range lUser.Collections {
						if req.CollectionName == c {
							chacking = true
						}
					}
				}
			}
		}
	}
	service.BError(&w, "There is no authentication to reach this database!", chacking)

	repo := Client(req.DatabaseName, req.DatabaseName)
	movies, err := repo.GetMovies()
	service.EError(&w, "Error getting movies", err)
	tMovies, err := json.Marshal(movies)
	service.EError(&w, "Error marshalling database movies to json", err)
	fmt.Fprintf(w, string(tMovies))
}

func AddMovie(w http.ResponseWriter, r *http.Request) {
	userName := service.AUser

	body, err := io.ReadAll(r.Body)
	service.EError(&w, "Body can not being read", err)

	var req model.Movie
	err = json.Unmarshal(body, &req)
	service.EError(&w, "Unmarshalling body", err)

	users, err := repository.GetUserJSON()
	service.EError(&w, "Getting users from JSON file", err)

	var user model.User
	for _, u := range users.User {
		if userName == u.UserName {
			user.Id = u.Id
			user.CreatorId = u.CreatorId
			user.Status = u.Status
			user.Collections = u.Collections
			user.Database = u.Database
		}
	}

	chacking := false
	if user.Status == "godmin" {
		for _, lUser := range users.User {
			if req.DatabaseName == lUser.Database {
				for _, c := range lUser.Collections {
					if req.CollectionName == c {
						chacking = true
					}
				}
			}
		}
	} else if user.Status == "admin" {
		for _, lUser := range users.User {
			if user.Id == lUser.CreatorId || user.Id == lUser.Id {
				if req.DatabaseName == lUser.Database {
					for _, c := range lUser.Collections {
						if req.CollectionName == c {
							chacking = true
						}
					}
				}
			}
		}
	} else if user.Status == "user" {
		for _, lUser := range users.User {
			if user.Id == lUser.Id {
				if req.DatabaseName == lUser.Database {
					for _, c := range lUser.Collections {
						if req.CollectionName == c {
							chacking = true
						}
					}
				}
			}
		}
	}

	service.BError(&w, "There is no authentication to reach this database", chacking)

	repo := Client(req.DatabaseName, req.CollectionName)

	err = repo.AddMovie(req)
	service.EError(&w, "adding database", err)

	fmt.Fprintf(w, "success")
}
