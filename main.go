package main

import (
	"fmt"
	"github.com/bndr/gopencils"
	"html"
	"log"
	"net/http"
)

type Joke struct {
	Type string `json:"type"`
	Value struct {
		Id int `json:"id"`
		Joke string `json:"joke"`
		Categories []string `json:"categories"`
	} `json:"value"`
}

func main() {
	randomJoke := &Joke{}
	url := "http://api.icndb.com/jokes/"
	api := gopencils.Api(url)

	http.HandleFunc("/joke", func(w http.ResponseWriter, r *http.Request){
		api.Res("random",randomJoke).Get()
		fmt.Fprint(w, html.UnescapeString(randomJoke.Value.Joke))
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}
