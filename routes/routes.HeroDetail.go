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

	// Define constants
	vars := mux.Vars(r)
    heroName := vars["heroName"]
	const selectionContainer = "body section#overview"
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
	heroDifficulty := doc.Find("body section#overview div.hero-detail-wrapper div.hero-detail-difficulty .star").Length() - doc.Find("div.hero-detail-wrapper div.hero-detail-difficulty .star.m-empty").Length()
	heroRole := doc.Find("body section#overview div.hero-detail-wrapper h4.hero-detail-role-name").Text()
	heroDescription := doc.Find("body section#overview div.hero-detail-wrapper p.hero-detail-description").Text()
	
	heroVideos := make([]string, 0, 0)
	doc.Find("div.hero-detail-video source").Each(func(index int, item *goquery.Selection) {

		src, _ := item.Attr("src")
		heroVideos = append(heroVideos, src)

	})

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