package handlers

import (
	"net/http"

	"github.com/Cloudykit/jet/v6"
)

var views = jet.NewSet(
	jet.NewOsfileSystemLoader("./html"),
	jet.InDevelopmentMode(),
)

func Home(w http.ResponseWriter, r *http.Request) {

}

func renderPage()
