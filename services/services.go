package services

import (
	"bmsp-backend-service/db"
	"bmsp-backend-service/repositories"
)

type Services struct {
	repo repositories.Repositories
}

func NewServices() Services {

	repo := repositories.NewRepositories(db.DBPg)
	return Services{repo: *repo}
}
