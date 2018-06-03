package repository

import (
	"log"

	. "github.com/nylo-andry/movies-service/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// MovieRepository is the entity responsible of getting movies from the database.
type MovieRepository struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	// COLLECTION is the name of the collection used in the databse.
	COLLECTION = "movies"
)

// Connect initializes the connection to the database.
func (m *MovieRepository) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

// FindAll queries all of the movies in the database.
func (m *MovieRepository) FindAll() ([]Movie, error) {
	var movies []Movie
	err := db.C(COLLECTION).Find(bson.M{}).All(&movies)
	return movies, err
}

// FindByID retrieves a single movie.
func (m *MovieRepository) FindByID(id string) (Movie, error) {
	var movie Movie
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&movie)
	return movie, err
}

// Insert creates a Movie document in the database
func (m *MovieRepository) Insert(movie Movie) error {
	err := db.C(COLLECTION).Insert(&movie)
	return err
}

// Delete removes a Movie document from the database.
func (m *MovieRepository) Delete(movie Movie) error {
	err := db.C(COLLECTION).Remove(&movie)
	return err
}

// Update updates a Movie document in the database.
func (m *MovieRepository) Update(movie Movie) error {
	err := db.C(COLLECTION).UpdateId(movie.ID, &movie)
	return err
}
