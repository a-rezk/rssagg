package main

import "net/http"

func handlerErrors(w http.ResponseWriter, r *http.Request) {
	respondWithErr(w, 400, "Something went wrong!")
}
