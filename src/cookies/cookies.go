package cookies

import (
	"net/http"
	"webapp_go/src/config"

	"github.com/gorilla/securecookie"
)

var s *securecookie.SecureCookie

func Configure() {
	s = securecookie.New(config.HashKey, config.BlockKey)
}

func Save(w http.ResponseWriter, ID, token string) error {
	data := map[string]string{
		"id":    ID,
		"token": token,
	}

	encoded_data, err := s.Encode("data", data)

	if err != nil {
		return err
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "data",
		Value:    encoded_data,
		Path:     "/",
		HttpOnly: true,
	})

	return nil
}

func Read(r *http.Request) (map[string]string, error) {
	cookie, err := r.Cookie("data")

	if err != nil {
		return nil, err
	}

	values := make(map[string]string)

	if err := s.Decode("data", cookie.Value, &values); err != nil {
		return nil, err
	}

	return values, nil
}
