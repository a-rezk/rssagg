package main

import (
	"net/http"
)

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	type Okay struct {
		Ok string `json:""`
	}
	okk := Okay{Ok: ""}
	respondWithJSON(w, 200, okk)
}
