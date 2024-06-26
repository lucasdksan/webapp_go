package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"webapp_go/src/config"
	"webapp_go/src/cookies"
	"webapp_go/src/requests"
	"webapp_go/src/responses"

	"github.com/gorilla/mux"
)

func Create_User(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	user, err := json.Marshal(map[string]string{
		"name":     r.FormValue("name"),
		"email":    r.FormValue("email"),
		"nick":     r.FormValue("nick"),
		"password": r.FormValue("password"),
	})

	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErroAPI{Erro: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/users", config.ApiURL)

	response, err := http.Post(url, "application/json", bytes.NewBuffer(user))

	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErroAPI{Erro: err.Error()})
		return
	}

	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.Treat_Error_Status_Code(w, response)
		return
	}

	responses.JSON(w, response.StatusCode, nil)
}

func Un_Follow_User(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	user_id, err := strconv.ParseInt(params["id"], 10, 64)

	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErroAPI{Erro: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/users/%d/unfollow", config.ApiURL, user_id)

	response, err := requests.Make_request_with_authentication(r, http.MethodPost, url, nil)

	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErroAPI{Erro: err.Error()})
		return
	}

	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.Treat_Error_Status_Code(w, response)
		return
	}

	responses.JSON(w, response.StatusCode, nil)
}

func Follow_User(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	user_id, err := strconv.ParseInt(params["id"], 10, 64)

	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErroAPI{Erro: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/users/%d/follow", config.ApiURL, user_id)

	response, err := requests.Make_request_with_authentication(r, http.MethodPost, url, nil)

	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErroAPI{Erro: err.Error()})
		return
	}

	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.Treat_Error_Status_Code(w, response)
		return
	}

	responses.JSON(w, response.StatusCode, nil)
}

func Profile_Edit(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	user, err := json.Marshal(map[string]string{
		"name":  r.FormValue("name"),
		"email": r.FormValue("email"),
		"nick":  r.FormValue("nick"),
	})

	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErroAPI{Erro: err.Error()})
		return
	}

	cookie, _ := cookies.Read(r)
	user_id, _ := strconv.ParseUint(cookie["id"], 10, 64)
	url := fmt.Sprintf("%s/users/%d", config.ApiURL, user_id)

	response, err := requests.Make_request_with_authentication(r, http.MethodPut, url, bytes.NewBuffer(user))

	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErroAPI{Erro: err.Error()})
		return
	}

	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.Treat_Error_Status_Code(w, response)
		return
	}

	responses.JSON(w, response.StatusCode, nil)
}

func Update_Pass(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	pass, err := json.Marshal(map[string]string{
		"CurrentPass": r.FormValue("current"),
		"NewPass":     r.FormValue("newPass"),
	})

	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErroAPI{Erro: err.Error()})
		return
	}

	cookie, _ := cookies.Read(r)
	user_id, _ := strconv.ParseUint(cookie["id"], 10, 64)

	url := fmt.Sprintf("%s/users/%d/update-password", config.ApiURL, user_id)

	response, err := requests.Make_request_with_authentication(r, http.MethodPost, url, bytes.NewBuffer(pass))

	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErroAPI{Erro: err.Error()})
		return
	}

	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.Treat_Error_Status_Code(w, response)
		return
	}

	responses.JSON(w, response.StatusCode, nil)
}

func User_Delete(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Read(r)
	user_id, _ := strconv.ParseUint(cookie["id"], 10, 64)

	url := fmt.Sprintf("%s/users/%d", config.ApiURL, user_id)

	response, err := requests.Make_request_with_authentication(r, http.MethodDelete, url, nil)

	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErroAPI{Erro: err.Error()})
		return
	}

	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.Treat_Error_Status_Code(w, response)
		return
	}

	responses.JSON(w, response.StatusCode, nil)
}
