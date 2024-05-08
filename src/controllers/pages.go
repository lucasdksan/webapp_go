package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"webapp_go/src/config"
	"webapp_go/src/cookies"
	"webapp_go/src/models"
	"webapp_go/src/requests"
	"webapp_go/src/responses"
	"webapp_go/src/utils"

	"github.com/gorilla/mux"
)

func Load_Login_Screen(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Read(r)

	if cookie["token"] != "" {
		http.Redirect(w, r, "/home", http.StatusFound)
		return
	}

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

	cookie, _ := cookies.Read(r)
	user_id, _ := strconv.ParseUint(cookie["id"], 10, 64)

	utils.Exec_Template(w, "home.html", struct {
		Publications []models.Publication
		UserID       uint64
	}{
		Publications: publications,
		UserID:       user_id,
	})
}

func Load_Update_Post_Screen(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	publication_id, err := strconv.ParseUint(params["id"], 10, 64)

	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErroAPI{Erro: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/publication/%d", config.ApiURL, publication_id)

	response, err := requests.Make_request_with_authentication(r, http.MethodGet, url, nil)

	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErroAPI{Erro: err.Error()})
		return
	}

	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.Treat_Error_Status_Code(w, response)
		return
	}

	var publication models.Publication

	if err = json.NewDecoder(response.Body).Decode(&publication); err != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, responses.ErroAPI{Erro: err.Error()})
		return
	}

	utils.Exec_Template(w, "publication-update.html", publication)
}
