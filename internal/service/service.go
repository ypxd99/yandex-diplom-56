package service

import "github.com/ypxd99/yandex-diplom-56/internal/repository"

type Service struct {
	repo repository.GophermartRepo
}

type GophermartService interface {
}

func InitService(repo repository.GophermartRepo) *Service {
	return &Service{repo: repo}
}
