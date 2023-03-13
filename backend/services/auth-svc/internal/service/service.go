package service

import "todos/internal/repository"

/************************/
/*  Прототипы сервисов  */
/************************/
type Authorization interface {
}

/************************/
/*  Хранилище сервисов  */
/************************/
type Service struct {
	Authorization
}

func NewService(r *repository.Repository) *Service {
	return &Service{}
}
