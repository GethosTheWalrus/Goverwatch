package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"routes"
)

func main() {

    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", routes.Index)
    //
    router.HandleFunc("/heroes", routes.HeroesIndex)
    router.HandleFunc("/heroes/{heroName}", routes.HeroDetail)
    router.HandleFunc("/players/{platform}/{region}/{gameMode}/{battleTag}", routes.PlayerDetail)
    router.HandleFunc("/players/{platform}/{region}/{gameMode}/{battleTag}/featuredStats", routes.FeaturedStatsDetail)
    router.HandleFunc("/players/{platform}/{region}/{gameMode}/{battleTag}/featuredStats/{statName}", routes.FeaturedStatsDetail)
    router.HandleFunc("/players/{platform}/{region}/{gameMode}/{battleTag}/careerStats", routes.CareerStatsDetail)
    router.HandleFunc("/players/{platform}/{region}/{gameMode}/{battleTag}/careerStats/{statName}", routes.CareerStatsDetail)

    log.Fatal(http.ListenAndServe(":3000", router))

}