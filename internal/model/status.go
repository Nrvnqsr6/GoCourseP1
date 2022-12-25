package model

import (
	"github.com/google/uuid"
)

type Status struct {
	Id     uuid.UUID
	Status string
}

type Identifiable interface {
	GetUUID() uuid.UUID
}

// можно предавать все данные модели в поле типа fullData string, если понадобится
func CreateStatus(model Identifiable, changeStatus string) *Status {
	return &Status{
		Id:     model.GetUUID(),
		Status: changeStatus,
	}
}
