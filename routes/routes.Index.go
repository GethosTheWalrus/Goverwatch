package routes

import (
	"fmt"
	"net/http"
	"encoding/json"
	"models"
)

func Index(w http.ResponseWriter, r *http.Request) {
    res, _ := json.Marshal(models.Response{"200", "Index", "Welcome to the unofficial overwatch API written in Go!"})
    fmt.Fprintln(w, string(res))
}