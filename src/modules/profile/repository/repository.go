package repository

import (
	"github.com/achimonchi/belajar_crud_mongodb/src/modules/profile/model"
)

type ProfileRepository interface {
	Save(*model.Profile) error
	Update(string, *model.Profile) error
	Delete(string) error
	FindAll() (model.Profiles, error)
	FindByID(string) (*model.Profile, error)
}
