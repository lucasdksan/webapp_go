package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"webapp_go/src/config"
	"webapp_go/src/models"
	"webapp_go/src/requests"
	"webapp_go/src/responses"
	"webapp_go/src/utils"
)

func Load_Login_Screen(w http.ResponseWriter, r *http.Request) {
	utils.Exec_Template(w, "login.html", nil)
}

func Load_Registration_Screen(w http.ResponseWriter, r *http.Request) {
	utils.Exec_Template(w, "registration.html", nil)
}

func Load_Home_Screen(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/publications", config.ApiURL)
	response, err := requests.Make_request_with_authentication(r, http.MethodGet, url, nil)

	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.Treat_Error_Status_Code(w, response)
		return
	}

	var publications []models.Publication

	if err := json.NewDecoder(response.Body).Decode(&publications); err != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	utils.Exec_Template(w, "home.html", publications)
}
