package routes

import (
	"fmt"
	// "log"
	"net/http"
	"encoding/json"
	// "github.com/PuerkitoBio?/goquery"
	"github.com/gorilla/mux"
	"models"
)

func PlayerDetail(w http.ResponseWriter, r *http.Request) {

	// Define constants
	vars := mux.Vars(r)
    battleTag := vars["battleTag"]
	url := "https://playoverwatch.com/en-us/heroes/" + battleTag

	// Fetch document body
    doc, err := goquery.NewDocument(url)

	if err != nil {
		log.Fatal(err)
		res, _ := json.Marshal(models.Response{"500", "heroDetail", "error with request"})
		fmt.Fprintln(w, string(res))
		return
	}

	playerName := doc.Find("div.masthead h1.header-masthead").text()
	playerPortrait, _ := doc.Find("div.masthead img.player-portrait").Attr("src")

	// Create response struct
	res, _ := json.Marshal(models.Response{"200", "PlayerDetail", "Player detail for " + battleTag})
	fmt.Fprintln(w, string(res))

}