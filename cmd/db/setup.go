package main

import (
	"github.com/globalsign/mgo"
	"log"
	"os"
)

func main() {
	var err error
	mongoURL := os.Getenv("MONGO_URL")
	if mongoURL == "" {
		log.Fatal("MONGO_URL not provided")
	}

	session, err := mgo.Dial(mongoURL)
	defer session.Close()
	if err != nil {
		log.Fatalf("Dial '%s' failed: %v", mongoURL, err)
	}

	err = session.DB("").AddUser("mongo_user", "mongo_secret", false)
	if err != nil {
		log.Fatalf("Failed to add user: %v", err)
	}

	info := &mgo.CollectionInfo{}
	err = session.DB("").C("kudos").Create(info)
	if err != nil {
		log.Fatal(err)
	}
}
