package main

import (
	"fmt"
	"log"

	"github.com/robertobouses/client-kpi-api/app"
	"github.com/robertobouses/client-kpi-api/internal"
	"github.com/robertobouses/client-kpi-api/repository"

	"github.com/robertobouses/client-kpi-api/http"
)

func main() {

	fmt.Println("Holaa")

	db, err := internal.NewPostgres(internal.DBConfig{
		User:     "postgres",
		Pass:     "mysecretpassword",
		Host:     "localhost",
		Port:     "5432",
		Database: "easy_football_tycoon",
	})

	if err != nil {
		log.Println(err)
		panic(err)
	}

	clientRepo, err := repository.NewRepository(db)
	if err != nil {
		panic(err)
	}

	app := app.NewApp(clientRepo)

	clientHandler := http.NewHandler(app)

	s := http.NewServer(clientHandler)

	s.Run("8080")

}
