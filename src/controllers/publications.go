package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"webapp_go/src/config"
	"webapp_go/src/requests"
	"webapp_go/src/responses"
)

func Create_Publication(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	publication, err := json.Marshal(map[string]string{
		"title":   r.FormValue("title"),
		"content": r.FormValue("content"),
	})

	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErroAPI{Erro: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/publication", config.ApiURL)

	response, err := requests.Make_request_with_authentication(r, http.MethodPost, url, bytes.NewBuffer(publication))

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
