package webServer

import (
	"log"
	"net/http"
	"time"

	rice "github.com/GeertJohan/go.rice"
	"github.com/gorilla/mux"
)

func serveAppHandler(app *rice.Box) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		indexFile, err := app.Open("index.html")
		if err != nil {
			http.Error(
				w,
				"Internal Server Error",
				http.StatusInternalServerError,
			)
			return
		}

		http.ServeContent(w, r, "index.html", time.Time{}, indexFile)
	}
}

func Server() {

	appBox, err := rice.FindBox("../static")
	if err != nil {
		log.Fatal(err)
	}

	muxRouter := mux.NewRouter()

	muxRouter.HandleFunc("/", serveAppHandler(appBox))

	log.Fatal(
		http.ListenAndServe(
			":8080",
			muxRouter,
		),
	)
}
