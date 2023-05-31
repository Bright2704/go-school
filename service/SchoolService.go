package service

import "example/BatteryTracking/entity"

type SchoolService interface {
	CreateSchool(*entity.UserSchool) error
	GetSchool(*string) (*entity.UserSchool, error)
	GetAll() ([]*entity.UserSchool, error)
	UpdateSchool(*entity.UserSchool) error
	DeleteSchool(*string) error
}