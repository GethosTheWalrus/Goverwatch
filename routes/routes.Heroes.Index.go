package routes

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"models"
)

// HeroesIndex will return a list of all Overwatch heroes
// along with their respective names, portraits, and role(s)
// URL: /heroes
func HeroesIndex(w http.ResponseWriter, r *http.Request) {
	// Define constants
	const selectionContainer = "body div.hero-portrait-detailed-container"
	const url = "https://playoverwatch.com/en-us/heroes/"

	// Check cache before making request
	cachedData := readCache("HeroesIndex", "heroes")
	if cachedData != "" {

		fmt.Fprintln(w, cachedData)
		return;

	}

	// Fetch document body
    doc, err := goquery.NewDocument(url)

	if err != nil {
		log.Fatal(err)
		res, _ := json.Marshal(models.Response{"500", "heroDetail", "error with request"})
		fmt.Fprintln(w, string(res))
		return
	}

	// Loop through hero tiles
	heroList := make([]models.HeroListHero, 0, 0)
	doc.Find(selectionContainer).Each(func(index int, item *goquery.Selection) {

		heroTile := item

		// Scrape values from web document
		heroName := heroTile.Find("span.container span.portrait-title").Text()
		heroPortrait, _ := heroTile.Find("a.hero-portrait-detailed img.portrait").Attr("src")
		heroRole, _ := heroTile.Attr("data-groups")

		// Create hero struct and append it to the hero list
		h := models.HeroListHero{heroName, heroRole, heroPortrait}
		heroList = append(heroList, h)
		
	})

	// Create response struct
	res, _ := json.Marshal(models.Response{"200", "heroesIndex", heroList})
	fmt.Fprintln(w, string(res))

	// Cache response for future use
	cache("HeroesIndex", "heroes", string(res))

}