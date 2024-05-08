package controllers

import (
	"net/http"
	"webapp_go/src/cookies"
)

func Handle_Logout(w http.ResponseWriter, r *http.Request) {
	cookies.Delete(w)

	http.Redirect(w, r, "/login", http.StatusFound)
}
