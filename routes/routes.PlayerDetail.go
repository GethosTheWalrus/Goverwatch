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
    region := vars["region"]
    platform := vars["platform"]
    gameMode := vars["gameMode"]
	url := "https://playoverwatch.com/en-us/career/" + platform + "/" + region + "/" + battleTag

	// Fetch document body
    doc, err := goquery.NewDocument(url)

	if err != nil {
		log.Fatal(err)
		res, _ := json.Marshal(models.Response{"500", "heroDetail", "error with request"})
		fmt.Fprintln(w, string(res))
		return
	}

	// Player profile stats
	playerName := doc.Find("div.masthead h1.header-masthead").Text()
	playerPortrait, _ := doc.Find("div.masthead img.player-portrait").Attr("src")
	playerRankIcon, _ := doc.Find("div.masthead div.competitive-rank img").Attr("src")
	playerRankNumber := doc.Find("div.masthead div.competitive-rank div.h6").Text()
	playerRank := [2]string{playerRankIcon, playerRankNumber}

	// Featured Stats
	featuredStats := make([]models.Stat, 0, 0)
	doc.Find("#" + gameMode + " div.card").Each(func(index int, item *goquery.Selection) {

		// Scrape data
		statIcon, _ := item.Find("div.bg-icon").Html()
		statName := item.Find("div.card-content p.card-copy").Text()
		statValue := item.Find("div.card-content h3.card-heading").Text()

		// Create stat struct
		s := models.Stat{statName, statValue, statIcon}
		featuredStats = append(featuredStats, s)

	})

	// Career stats
	careerStats := make([]models.StatCategory, 0, 0)
	doc.Find("#" + gameMode + " div.card-stat-block").Each(func(index int, item *goquery.Selection) {

		// Scrape category info
		categoryName := item.Find("thead th span.stat-title").Text()
		categoryIcon, _ := item.Find("thead th svg").Html()

		// Loop through stats and scrape
		categoryStats := make([]models.Stat, 0, 0)
		item.Find("tbody tr").Each(func(index int, statItem *goquery.Selection) {

			// Scrape
			statName := statItem.Find("td").First().Text()
			statValue := statItem.Find("td").Last().Text()

			// Create stat struct
			s := models.Stat{statName, statValue, ""}
			categoryStats = append(categoryStats, s)

		})

		// Create stat category struct
		sc := models.StatCategory{categoryName, categoryIcon, categoryStats}
		careerStats = append(careerStats, sc)

	})

	// Create player struct
	p := models.Player{playerName, battleTag, playerPortrait, playerRank, featuredStats, careerStats}

	// Create response struct
	res, _ := json.Marshal(models.Response{"200", "PlayerDetail", p})
	fmt.Fprintln(w, string(res))

}