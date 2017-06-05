package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"routes"
)

func main() {

    router := mux.NewRouter().StrictSlash(true)

    // Main Routes
    router.HandleFunc("/", routes.Index)
    router.HandleFunc("/heroes", routes.HeroesIndex)
    router.HandleFunc("/heroes/{heroName}", routes.HeroDetail)
    router.HandleFunc("/players/{platform}/{region}/{gameMode}/{battleTag}", routes.PlayerDetail)
    router.HandleFunc("/players/{platform}/{region}/{gameMode}/{battleTag}/featuredStats", routes.FeaturedStatsDetail)
    router.HandleFunc("/players/{platform}/{region}/{gameMode}/{battleTag}/featuredStats/{statName}", routes.FeaturedStatsDetail)
    router.HandleFunc("/players/{platform}/{region}/{gameMode}/{battleTag}/careerStats", routes.CareerStatsDetail)
    router.HandleFunc("/players/{platform}/{region}/{gameMode}/{battleTag}/careerStats/{statName}", routes.CareerStatsDetail)
    router.HandleFunc("/players/{platform}/{region}/{gameMode}/{battleTag}/heroComparison", routes.HeroComparisonDetail)
    router.HandleFunc("/players/{platform}/{region}/{gameMode}/{battleTag}/heroComparison/stat/{statName}", routes.HeroComparisonDetail)
    router.HandleFunc("/players/{platform}/{region}/{gameMode}/{battleTag}/heroComparison/hero/{heroName}", routes.HeroComparisonDetail)
    router.HandleFunc("/players/{platform}/{region}/{gameMode}/{battleTag}/heroComparison/hero/{heroName}/stat/{statName}", routes.HeroComparisonDetail)

    log.Fatal(http.ListenAndServe(":3000", router))

}