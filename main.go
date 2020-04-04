package main

import (
	"fmt"
	"time"

	"github.com/achimonchi/belajar_crud_mongodb/src/modules/profile/model"

	"github.com/achimonchi/belajar_crud_mongodb/config"
	"github.com/achimonchi/belajar_crud_mongodb/src/modules/profile/repository"
)

func main() {
	fmt.Println("CRUD with Go and Mongo")

	db, err := config.GetMongoDB()

	if err != nil {
		fmt.Println(err)
	}

	profileRepository := repository.NewProfileRepositoryMongo(db, "users")
	saveProfile(profileRepository)
}

// function untuk save data
func saveProfile(profileRepository repository.ProfileRepository) {
	var p model.Profile
	p.ID = "u1"
	p.FirstName = "Reyhan"
	p.LastName = "Jovie"
	p.Email = "reyhanjovie01@gmail.com"
	p.Password = "123456"
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()

	err := profileRepository.Save(&p)
	if err != nil {
		fmt.Println(err)
	}
}

// function untuk update data
func updateProfile(id string, profileRepository repository.ProfileRepository) {

}

// function untuk delete data
func deleteProfile(id string, profileRepository repository.ProfileRepository) {

}

// function untuk lihat seluruh data
func findAll(profileRepository repository.ProfileRepository) {

}

// function untuk lihat data berdasarkan id
func findByID(id string, profileRepository repository.ProfileRepository) {

}
