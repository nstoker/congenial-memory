package main

import (
	"log"
	"net/http"
	"os"

	web "github.com/nstoker/congenial-memory/pkg/http"
	"github.com/nstoker/congenial-memory/pkg/storage"
)

func main() {
	httpPort := os.Getenv("PORT")
	repo := storage.NewMongoRepository()
	webService := web.New(repo)

	log.Printf("Running on port %s\n", httpPort)
	log.Fatal(http.ListenAndServe(httpPort, webService.Router))
}
