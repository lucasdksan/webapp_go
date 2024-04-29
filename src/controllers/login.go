package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
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

	response, err := http.Post("http://localhost:8080/login", "application/json", bytes.NewBuffer(user))

	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErroAPI{Erro: err.Error()})
		return
	}

	fmt.Println(response.StatusCode, response.Body)
}
