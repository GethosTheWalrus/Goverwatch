package routes

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"github.com/gorilla/mux"
	"models"
)

func PlayerDetail(w http.ResponseWriter, r *http.Request) {

	// Define constants
	vars := mux.Vars(r)
    battleTag := vars["battleTag"]
	url := "https://playoverwatch.com/en-us/career/pc/us/" + battleTag

	// Fetch document body
    doc, err := goquery.NewDocument(url)

	if err != nil {
		log.Fatal(err)
		res, _ := json.Marshal(models.Response{"500", "heroDetail", "error with request"})
		fmt.Fprintln(w, string(res))
		return
	}

	playerName := doc.Find("div.masthead h1.header-masthead").Text()
	playerPortrait, _ := doc.Find("div.masthead img.player-portrait").Attr("src")
	playerRankIcon, _ := doc.Find("div.masthead div.competitive-rank img").Attr("src")
	playerRankNumber := doc.Find("div.masthead div.competitive-rank div.h6").Text()
	playerRank := [2]string{playerRankIcon, playerRankNumber}

	// Create player struct
	p := models.Player{playerName, battleTag, playerPortrait, playerRank}

	// Create response struct
	res, _ := json.Marshal(models.Response{"200", "PlayerDetail", p})
	fmt.Fprintln(w, string(res))

}