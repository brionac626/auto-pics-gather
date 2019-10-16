package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	// "github.com/brionac626/auto-pics-gather/hentai"

	"github.com/gorilla/mux"
)

func initRounter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/hentai/download", hentaiAPI)
	r.HandleFunc("/pic", nil)

	return r
}

type request struct {
	URLs []string `json:"urls"`
}

func hentaiAPI(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var downloadRequest request

	if err := json.Unmarshal(body, &downloadRequest); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// hentai.Download(downloadRequest.URLs)
}
