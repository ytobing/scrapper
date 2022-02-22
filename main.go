package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	m "github.com/ytobing/scrapper/models"
	h "github.com/ytobing/scrapper/src/html"
	p "github.com/ytobing/scrapper/src/parser"
)

func main() {
	a, err := initApp()
	if err != nil {
		log.Println(err)
		return
	}

	h.RegisterRoutes(a)
	p.RegisterRoutes(a)

	log.Printf("Server is starting on port: %s", os.Getenv("PORT"))

	log.Fatal(http.ListenAndServe(os.Getenv("PORT"), a.Router))

}

func initApp() (res m.App, err error) {
	res.Router = httprouter.New()

	err = godotenv.Load()
	if err != nil {
		return
	}

	return
}
