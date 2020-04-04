package model

import "time"

// Deskripsi dari Profile
type Profile struct {
	ID        string    `bson:"_id"`
	FirstName string    `bson:"first_name"`
	LastName  string    `bson:"last_name"`
	Email     string    `bson:"email"`
	Password  string    `bson:"password"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}

// List Profile
type Profiles []Profile
