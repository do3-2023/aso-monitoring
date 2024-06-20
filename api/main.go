package main

import (
	"apimonitoring/handlers"
	"net/http"

	spinhttp "github.com/fermyon/spin/sdk/go/v2/http"
)

func init() {
	spinhttp.Handle(func(w http.ResponseWriter, r *http.Request) {
		router := spinhttp.NewRouter()
		router.GET("/health", handlers.Health)
		router.GET("/person", handlers.FindAllPerson)
		router.GET("/person/:id", handlers.FindOnePerson)
		router.POST("/person", handlers.CreatePerson)
		router.DELETE("/person/:id", handlers.DeletePerson)

		router.ServeHTTP(w, r)
	})
}

func main() {}
