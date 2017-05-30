package routes

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"github.com/gorilla/mux"
	"models"
)

// FeaturedStats will scrape Blizzard's Overwatch user profile page for that account's
// featured statistics and add the data to the featuredStatsObject slice.
// 
// Parameters: 
// wg - A pointer to a wait group that will allow both goroutines to finish before creating the player object
// featuredStatsObject - A pointer to the object that will be used when creating the player object
// doc - The document body that GoQuery will be scraping from
// gameMode - The game mode that should be searched for in the DOM
func FeaturedStats(wg *sync.WaitGroup, featuredStatsObject *[]models.Stat, doc *goquery.Document, gameMode string) {
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

	*featuredStatsObject = featuredStats
	defer wg.Done()
}

// CareerStats will scrape Blizzard's Overwatch user profile page for that account's
// career statistics and add the data to the careerStatsObject slice.
// 
// Parameters: 
// wg - A pointer to a wait group that will allow both goroutines to finish before creating the player object
// featuredStatsObject - A pointer to the object that will be used when creating the player object
// doc - The document body that GoQuery will be scraping from
// gameMode - The game mode that should be searched for in the DOM
func CareerStats(wg *sync.WaitGroup, careerStatsObject *[]models.StatCategory, doc *goquery.Document, gameMode string) {
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

	*careerStatsObject = careerStats
	defer wg.Done()
}

// PlayerDetail will combine a user's featured stats, career stats, hero played time, into a JSON object and 
// print the result to the browser
// URL: /players/{platform}/{region}/{gameMode}/{battleTag} (only us and eu are currently supported for region)
func PlayerDetail(w http.ResponseWriter, r *http.Request) {

	// Create workgroup to wait for all goroutines to finish
	var wg sync.WaitGroup

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
	wg.Add(1)
	featuredStatsObject := make([]models.Stat, 0, 0)
	go FeaturedStats(&wg, &featuredStatsObject, doc, gameMode)

	// Career stats
	wg.Add(1)
	careerStatsObject := make([]models.StatCategory, 0, 0)
	go CareerStats(&wg, &careerStatsObject, doc, gameMode)

	// Create player struct
	wg.Wait()
	p := models.Player{playerName, battleTag, playerPortrait, playerRank, featuredStatsObject, careerStatsObject}

	// Create response struct
	res, _ := json.Marshal(models.Response{"200", "PlayerDetail", p})
	fmt.Fprintln(w, string(res))

}