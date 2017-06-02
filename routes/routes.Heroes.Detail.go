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

// HeroDetail will return a data object consisting of a specific hero's
// difficulty, role, description, idle animation, and a list of that hero's
// ability names, descriptions, and icons
// URL: /heroes/{heroName}
func HeroDetail(w http.ResponseWriter, r *http.Request) {

	// Route parameters
	vars := mux.Vars(r)
	// Name of the requested hero
    heroName := vars["heroName"]
    // HTML container where data is being scraped from
	const selectionContainer = "body section#overview"
	// URL of the request
	url := "https://playoverwatch.com/en-us/heroes/" + heroName

	// Fetch document body
    doc, err := goquery.NewDocument(url)

	if err != nil {
		log.Fatal(err)
		res, _ := json.Marshal(models.Response{"500", "heroDetail", "error with request"})
		fmt.Fprintln(w, string(res))
		return
	}

	// Scrape values from web document

	// Hero difficulty rating
	heroDifficulty := doc.Find("body section#overview div.hero-detail-wrapper div.hero-detail-difficulty .star").Length() - doc.Find("div.hero-detail-wrapper div.hero-detail-difficulty .star.m-empty").Length()
	// Hero roles
	heroRole := doc.Find("body section#overview div.hero-detail-wrapper h4.hero-detail-role-name").Text()
	// Hero description
	heroDescription := doc.Find("body section#overview div.hero-detail-wrapper p.hero-detail-description").Text()
	
	// Hero videos (idle pose, etc)
	heroVideos := make([]string, 0, 0)
	doc.Find("div.hero-detail-video source").Each(func(index int, item *goquery.Selection) {

		src, _ := item.Attr("src")
		heroVideos = append(heroVideos, src)

	})

	// List of hero abilities
	heroAbilities := make([]models.Ability, 0, 0)
	doc.Find("body section#overview .hero-ability").Each(func(index int, item *goquery.Selection) {

		abilityTile := item

		// Scrape values from web document
		abilityName := abilityTile.Find(".hero-ability-descriptor h4.h5").Text()
		abilityDescription := abilityTile.Find(".hero-ability-descriptor p").Text()
		abilityIcon, _ := abilityTile.Find(".hero-ability-icon-container img.hero-ability-icon").Attr("src")

		// Create ability struct and append it to the ability slice
		a := models.Ability{abilityName, abilityDescription, abilityIcon}
		heroAbilities = append(heroAbilities, a)

	})

	// Create hero struct
	h := models.Hero{heroName, heroRole, heroDescription, heroDifficulty, heroVideos, heroAbilities}

	// Create response struct
	res, _ := json.Marshal(models.Response{"200", "heroDetail", h})
	fmt.Fprintln(w, string(res))

}