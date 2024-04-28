package controllers

import (
	"net/http"
	"webapp_go/src/utils"
)

func Load_Login_Screen(w http.ResponseWriter, r *http.Request) {
	utils.Exec_Template(w, "login.html", nil)
}

func Load_Registration_Screen(w http.ResponseWriter, r *http.Request) {
	utils.Exec_Template(w, "registration.html", nil)
}
