// Package containing the functions that will be called for each route
package routes

import (
	"fmt"
	"net/http"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"models"
)

func Index(w http.ResponseWriter, r *http.Request) {
    res, _ := json.Marshal(models.Response{"200", "Index", "Welcome to the unofficial overwatch API written in Go!"})
    fmt.Fprintln(w, string(res))
}

func cache(requestType string, identifier string, data string) {

	db, err := sql.Open("mysql", "root@/goverwatch?sql_mode=TRADITIONAL")
	checkErr(err)

	// Insert cached data
	cache, err := db.Prepare("INSERT requests SET type=?, identifier=?, data=?, timestamp=NOW()")
	cache.Exec(requestType, identifier, data)

}

func readCache(requestType string, identifier string) string {

	db, err := sql.Open("mysql", "root@/goverwatch?sql_mode=TRADITIONAL")
	checkErr(err)

	// Update cached data hit count
	update, err := db.Prepare("UPDATE requests SET hit_count=? WHERE type=? AND identifier=? AND `timestamp` >= DATE_SUB(NOW(), INTERVAL 10 MINUTE)")

	// Get cached data
	var data string
	var hitCount int
	_ = db.QueryRow("SELECT data, hit_count FROM requests WHERE `type` = '" + requestType + "' AND `identifier` = '" + identifier + "' AND `timestamp` >= DATE_SUB(NOW(), INTERVAL 10 MINUTE)").Scan(&data, &hitCount)

	// Execute update once hit count has been
	// retrieved from database
	if data != "" {

		update.Exec(hitCount + 1, requestType, identifier)

	}

	db.Close()

	return data

}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}