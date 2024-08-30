package controllers

import (
	"net/http"
)

func top(w http.ResponseWriter, req *http.Request) {
	generateHTML(w, "helllllllo", "layout", "top")
}
