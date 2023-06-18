package handlers

import services "crud-go/pkg/cars/service"

type Handler struct {
	Service services.CarService
}

func New() *Handler {
	return &Handler{
        Service: services.New(),
    }
}