package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"polybay/configurator"
	"polybay/internal/apps"
)

func main() {
	config, err := configurator.Read(os.Getenv("CONFIG_PATH"))
	if err != nil {
		panic(err)
	}
	app := apps.New()

	log.Println(fmt.Sprintf("Listening on %s port", config.Port))
	http.ListenAndServe(fmt.Sprintf(":%s", config.Port), app.Route())
}
