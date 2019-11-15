package storage

import (
	"log"
	"os"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/nstoker/congenial-memory/pkg/core"
)

const (
	collectionName = "kudos"
)

// GetCollectionName does what it says on the tin
func GetCollectionName() string {
	return collectionName
}

// MongoRepository does what it says on the tin
type MongoRepository struct {
	logger  *log.Logger
	session *mgo.Session
}

// Find fetches a kudo from mongo
func (r MongoRepository) Find(repoID string) (*core.Kudo, error) {
	session := r.session.Copy()
	defer session.Close()
	coll := session.DB("").C(collectionName)

	var kudo core.Kudo
	err := coll.Find(bson.M{"repoID": repoID, "userId": kudo.UserID}).One(&kudo)
	if err != nil {
		r.logger.Printf("error: %v\n", err)
	}

	return &kudo, nil
}

// FindAll fetches all kudos from the database.
func (r MongoRepository) FindAll(selector map[string]interface{}) ([]*core.Kudo, error) {
	session := r.session.Copy()
	defer session.Close()
	coll := session.DB("").C(collectionName)

	var kudos []*core.Kudo
	err := coll.Find(selector).All(&kudos)
	if err != nil {
		r.logger.Printf("error: %v\n", err)
		return nil, err
	}

	return kudos, nil
}

// Delete deletes a kudo from mongo
func (r MongoRepository) Delete(kudo *core.Kudo) error {
	session := r.session.Copy()
	defer session.Close()
	coll := session.DB("").C(collectionName)

	return coll.Remove(bson.M{"repoId": kudo.RepoID, "userId": kudo.UserID})
}

// Update updates a kudo
func (r MongoRepository) Update(kudo *core.Kudo) error {
	session := r.session.Copy()
	defer session.Close()
	coll := session.DB("").C(collectionName)

	return coll.Update(bson.M{"repoId": kudo.RepoID, "userId": kudo.UserID}, kudo)
}

// Create kudos in the database
func (r MongoRepository) Create(kudos ...*core.Kudo) error {
	session := r.session.Copy()
	defer session.Close()
	coll := session.DB("").C(collectionName)

	for _, kudo := range kudos {
		_, err := coll.Upsert(bson.M{"repoId": kudo.RepoID, "userId": kudo.UserID}, kudo)
		if err != nil {
			return err
		}
	}

	return nil
}

// Count counts documents for a collection
func (r MongoRepository) Count() (int, error) {
	session := r.session.Copy()
	defer session.Close()
	coll := session.DB("").C(collectionName)

	return coll.Count()
}

// NewMongoSession dials mongodb and creates a session
func NewMongoSession() (*mgo.Session, error) {
	mongoURL := os.Getenv("MONGO_URL")
	if mongoURL == "" {
		log.Fatal(("MONGO_URL not provided"))
	}

	return mgo.Dial(mongoURL)
}

func newMongoRepositoryLogger() *log.Logger {
	return log.New(os.Stdout, "[mongoDB] ", 0)
}

// NewMongoRepository creates a new repo?
func NewMongoRepository() core.Repository {
	logger := newMongoRepositoryLogger()
	session, err := NewMongoSession()
	if err != nil {
		logger.Fatalf("Could not connect to the database: %v\n", err)
	}

	return MongoRepository{
		session: session,
		logger:  logger,
	}
}
