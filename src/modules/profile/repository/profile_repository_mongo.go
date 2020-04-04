package repository

import (
	"time"

	"github.com/achimonchi/belajar_crud_mongodb/src/modules/profile/model"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type profileRepositoryMongo struct {
	db         *mgo.Database
	collection string
}

func NewProfileRepositoryMongo(db *mgo.Database, collection string) *profileRepositoryMongo {
	return &profileRepositoryMongo{
		db:         db,
		collection: collection,
	}
}

func (r *profileRepositoryMongo) Save(p *model.Profile) error {
	err := r.db.C(r.collection).Insert(p)
	return err
}

func (r *profileRepositoryMongo) Update(id string, p *model.Profile) error {
	p.UpdatedAt = time.Now()
	err := r.db.C(r.collection).Update(bson.M{"_id": id}, p)
	return err
}

func (r *profileRepositoryMongo) Delete(id string) error {
	err := r.db.C(r.collection).Remove(bson.M{"_id": id})
	return err
}

func (r *profileRepositoryMongo) FindAll() (model.Profiles, error) {
	var profiles model.Profiles
	err := r.db.C(r.collection).Find(bson.M{}).All(&profiles)
	if err != nil {
		return nil, err
	}
	return profiles, nil
}

func (r *profileRepositoryMongo) FindByID(id string) (*model.Profile, error) {
	var profile model.Profile
	err := r.db.C(r.collection).Find(bson.M{"_id": id}).One(&profile)
	if err != nil {
		return nil, err
	}
	return &profile, nil
}
