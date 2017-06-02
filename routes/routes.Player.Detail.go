package routes

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"strings"
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
func FeaturedStats(wg *sync.WaitGroup, featuredStatsObject *[]models.Stat, doc *goquery.Document, gameMode string, statNameParam string) {
	featuredStats := make([]models.Stat, 0, 0)
	doc.Find("#" + gameMode + " div.card").Each(func(index int, item *goquery.Selection) {

		// Scrape data
		statIcon, _ := item.Find("div.bg-icon").Html()
		statName := item.Find("div.card-content p.card-copy").Text()
		statValue := item.Find("div.card-content h3.card-heading").Text()

		if (statNameParam != "" && strings.ToLower(statNameParam) == strings.ToLower(statName)) || statNameParam == "" {
			// Create stat struct 
			s := models.Stat{statName, statValue, statIcon}
			featuredStats = append(featuredStats, s)
		}

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
func CareerStats(wg *sync.WaitGroup, careerStatsObject *[]models.StatCategory, doc *goquery.Document, gameMode string, statNameParam string) {
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

			if (statNameParam != "" && strings.ToLower(statNameParam) == strings.ToLower(statName)) || statNameParam == "" {
				// Create stat struct
				s := models.Stat{statName, statValue, ""}
				categoryStats = append(categoryStats, s)
			}

		})

		if len(categoryStats) > 0 {
			// Create stat category struct
			sc := models.StatCategory{categoryName, categoryIcon, categoryStats}
			careerStats = append(careerStats, sc)
		}

	})

	*careerStatsObject = careerStats
	defer wg.Done()
}

// HeroComparison will scrape Blizzard's Overwatch user profile page for that account's
// hero comparison metrics, such as played time, games won, etc. This will return an
// object that contains all available metrics for the specified game type
// 
// Parameters: 
// wg - A pointer to a wait group that will allow both goroutines to finish before creating the player object
// featuredStatsObject - A pointer to the object that will be used when creating the player object
// doc - The document body that GoQuery will be scraping from
// gameMode - The game mode that should be searched for in the DOM
func HeroComparison(wg *sync.WaitGroup, heroComparisonMetricsObject *[]models.HeroComparisonMetric, doc *goquery.Document, gameMode string, statNameParam string, heroNameParam string) {
	// Loop through the available metrics
	heroComparisonMetrics := make([]models.HeroComparisonMetric, 0, 0)
	doc.Find("#" + gameMode + " select[data-group-id='comparisons'] option").Each(func(index int, item *goquery.Selection) {

		// Get the identifier for each metric from the dropdown and loop through the
		// container for that identifier
		overwatchGuid, _ := item.Attr("value")
		metricName := item.Text()
		heroComparisonHeroes := make([]models.HeroComparisonHero, 0, 0)
		hch := models.HeroComparisonHero{"", "", "", ""}

		if strings.ToLower(statNameParam) == strings.ToLower(metricName) || statNameParam == "" {

			doc.Find("#" + gameMode + " div[data-category-id='" + overwatchGuid + "'] div.progress-category-item").Each(func(index int, progressItem *goquery.Selection) {

				// Scrape the stats from the specified container
				heroName := progressItem.Find("div.bar-text div.title").Text()
				metricValue := progressItem.Find("div.bar-text div.description").Text()
				heroImage, _ := progressItem.Find("img").Attr("src")
				metricPercent, _ := progressItem.Attr("data-overwatch-progress-percent")

				if strings.ToLower(heroNameParam) == strings.ToLower(heroName) || heroNameParam == "" {
					// Create a hero object and store the values,
					// then append the object to the slice of heroes
					hch = models.HeroComparisonHero{heroName, metricValue, heroImage, metricPercent}
					heroComparisonHeroes = append(heroComparisonHeroes, hch)
				}

			})

		}

		// Add the metric to the slice of metrics and add the metric
		// to the slice of metrics
		if len(heroComparisonHeroes) > 0 {
			hcm := models.HeroComparisonMetric{metricName, heroComparisonHeroes}
			heroComparisonMetrics = append(heroComparisonMetrics, hcm)
		}

	})

	*heroComparisonMetricsObject = heroComparisonMetrics
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
	playerRankNumber := doc.Find("div.masthead div.competitive-rank div.h6").First().Text()
	playerRank := [2]string{playerRankIcon, playerRankNumber}

	// Featured Stats
	wg.Add(1)
	featuredStatsObject := make([]models.Stat, 0, 0)
	go FeaturedStats(&wg, &featuredStatsObject, doc, gameMode, "")

	// Career stats
	wg.Add(1)
	careerStatsObject := make([]models.StatCategory, 0, 0)
	go CareerStats(&wg, &careerStatsObject, doc, gameMode, "")

	// Hero Comparison Metrics
	wg.Add(1)
	heroComparisonMetricsObject := make([]models.HeroComparisonMetric, 0, 0)
	go HeroComparison(&wg, &heroComparisonMetricsObject, doc, gameMode, "", "")

	// Create player struct
	wg.Wait()
	p := models.Player{playerName, battleTag, playerPortrait, playerRank, featuredStatsObject, careerStatsObject, heroComparisonMetricsObject}

	// Create response struct
	res, _ := json.Marshal(models.Response{"200", "PlayerDetail", p})
	fmt.Fprintln(w, string(res))

}

// FeaturedStatsDetail will return a list of a user's featured stats, or an individual featured stat
// print the result to the browser
// URL: /players/{platform}/{region}/{gameMode}/{battleTag}/featuredStats/{statName?} (only us and eu are currently supported for region)
func FeaturedStatsDetail(w http.ResponseWriter, r *http.Request) {

	// Create workgroup to wait for all goroutines to finish
	var wg sync.WaitGroup

	// Define constants
	vars := mux.Vars(r)
    battleTag := vars["battleTag"]
    region := vars["region"]
    platform := vars["platform"]
    gameMode := vars["gameMode"]
    statNameParam := vars["statName"]
	url := "https://playoverwatch.com/en-us/career/" + platform + "/" + region + "/" + battleTag

	// Fetch document body
    doc, err := goquery.NewDocument(url)

	if err != nil {
		log.Fatal(err)
		res, _ := json.Marshal(models.Response{"500", "heroDetail", "error with request"})
		fmt.Fprintln(w, string(res))
		return
	}

	// Featured Stats
	wg.Add(1)
	featuredStatsObject := make([]models.Stat, 0, 0)
	go FeaturedStats(&wg, &featuredStatsObject, doc, gameMode, statNameParam)

	wg.Wait()
	res, _ := json.Marshal(models.Response{"200", "FeaturedStatsDetail", featuredStatsObject})
	fmt.Fprintln(w, string(res))

}

// CareerStatsDetail will return a list of a user's career stats, or an individual career stat
// print the result to the browser
// URL: /players/{platform}/{region}/{gameMode}/{battleTag}/careerStats/{statName?} (only us and eu are currently supported for region)
func CareerStatsDetail(w http.ResponseWriter, r *http.Request) {

	// Create workgroup to wait for all goroutines to finish
	var wg sync.WaitGroup

	// Define constants
	vars := mux.Vars(r)
    battleTag := vars["battleTag"]
    region := vars["region"]
    platform := vars["platform"]
    gameMode := vars["gameMode"]
    statNameParam := vars["statName"]
	url := "https://playoverwatch.com/en-us/career/" + platform + "/" + region + "/" + battleTag

	// Fetch document body
    doc, err := goquery.NewDocument(url)

	if err != nil {
		log.Fatal(err)
		res, _ := json.Marshal(models.Response{"500", "heroDetail", "error with request"})
		fmt.Fprintln(w, string(res))
		return
	}

	// Career stats
	wg.Add(1)
	careerStatsObject := make([]models.StatCategory, 0, 0)
	go CareerStats(&wg, &careerStatsObject, doc, gameMode, statNameParam)

	wg.Wait()
	res, _ := json.Marshal(models.Response{"200", "CareerStatsDetail", careerStatsObject})
	fmt.Fprintln(w, string(res))

}

// HeroComparisonDetail will return a list of metrics by which a player's hero performance is measured for each hero and values
// associated with each metric. This request can be limited to a single hero, a single stat, or a single stat FOR a single hero
// print the result to the browser
// URL: /players/{platform}/{region}/{gameMode}/{battleTag}/heroComparison/[hero/{heroName},stat/{statName},hero/{heroName}/stat/{statName}]? (only us and eu are currently supported for region)
func HeroComparisonDetail(w http.ResponseWriter, r *http.Request) {

	// Create workgroup to wait for all goroutines to finish
	var wg sync.WaitGroup

	// Define constants
	vars := mux.Vars(r)
    battleTag := vars["battleTag"]
    region := vars["region"]
    platform := vars["platform"]
    gameMode := vars["gameMode"]
    heroNameParam := vars["heroName"]
    statNameParam := vars["statName"]
	url := "https://playoverwatch.com/en-us/career/" + platform + "/" + region + "/" + battleTag

	// Fetch document body
    doc, err := goquery.NewDocument(url)

	if err != nil {
		log.Fatal(err)
		res, _ := json.Marshal(models.Response{"404", "HeroComparisonDetail", "error with this request"})
		fmt.Fprintln(w, string(res))
		return
	}

	// Hero Comparison Metrics
	wg.Add(1)
	heroComparisonMetricsObject := make([]models.HeroComparisonMetric, 0, 0)
	go HeroComparison(&wg, &heroComparisonMetricsObject, doc, gameMode, statNameParam, heroNameParam)

	wg.Wait()
	res, _ := json.Marshal(models.Response{"200", "HeroComparisonDetail", heroComparisonMetricsObject})
	fmt.Fprintln(w, string(res))

}