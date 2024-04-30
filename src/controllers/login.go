package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"webapp_go/src/config"
	"webapp_go/src/cookies"
	"webapp_go/src/models"
	"webapp_go/src/responses"
)

func Handle_Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	user, err := json.Marshal(map[string]string{
		"email":    r.FormValue("email"),
		"password": r.FormValue("password"),
	})

	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErroAPI{Erro: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/login", config.ApiURL)

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

	var authentication_data models.Authentication_data

	if err := json.NewDecoder(response.Body).Decode(&authentication_data); err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErroAPI{Erro: err.Error()})
		return
	}

	if err := cookies.Save(w, authentication_data.ID, authentication_data.Token); err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErroAPI{Erro: err.Error()})
		return
	}

	responses.JSON(w, http.StatusOK, nil)
}
