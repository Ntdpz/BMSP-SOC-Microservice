package handlers

import (
	"bmsp-backend-service/services"
)

type handlers struct {
	service services.Services
}

func NewHandlers() *handlers {

	_service := services.NewServices()
	return &handlers{service: _service}
}
