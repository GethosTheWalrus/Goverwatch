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
    router.HandleFunc("/heroes", routes.HeroesIndex)
    router.HandleFunc("/heroes/{heroName}", routes.HeroDetail)
    router.HandleFunc("/players/{platform}/{region}/{gameMode}/{battleTag}", routes.PlayerDetail)

    log.Fatal(http.ListenAndServe(":3000", router))

}