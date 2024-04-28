package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

type ErroAPI struct {
	Erro string `json:"erro"`
}

func JSON(w http.ResponseWriter, status_code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status_code)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Fatal(err)
	}
}

func Treat_Error_Status_Code(w http.ResponseWriter, r *http.Response) {
	var err ErroAPI

	json.NewDecoder(r.Body).Decode(&err)
	JSON(w, r.StatusCode, err)
}
