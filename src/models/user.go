package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
	"webapp_go/src/config"
	"webapp_go/src/requests"
)

type User struct {
	ID           uint64        `json:"id"`
	Name         string        `json:"name"`
	Email        string        `json:"email"`
	Nick         string        `json:"nick"`
	CreateAt     time.Time     `json:"createAt"`
	Followers    []User        `json:"followers"`
	Following    []User        `json:"following"`
	Publications []Publication `json:"publications"`
}

func Search_User_Complete(user_id uint64, r *http.Request) (User, error) {
	channel_user := make(chan User)
	channel_followers := make(chan []User)
	channel_following := make(chan []User)
	channel_publications := make(chan []Publication)

	go Search_Data_User(channel_user, user_id, r)
	go Search_Data_Followers(channel_followers, user_id, r)
	go Search_Data_Following(channel_following, user_id, r)
	go Search_Data_Publications(channel_publications, user_id, r)

	var (
		user         User
		followers    []User
		following    []User
		publications []Publication
	)

	for i := 0; i < 4; i++ {
		select {
		case user_loaded := <-channel_user:
			if user_loaded.ID == 0 {
				return User{}, errors.New("erro ao buscar o usuário")
			}

			user = user_loaded

		case followers_loaded := <-channel_followers:
			if len(followers_loaded) == 0 {
				followers = followers_loaded
			} else if followers_loaded == nil {
				return User{}, errors.New("erro ao buscar o seguidores")
			}

			followers = followers_loaded

		case following_loaded := <-channel_following:
			if len(following_loaded) == 0 {
				following = following_loaded
			} else if following_loaded == nil {
				return User{}, errors.New("erro ao buscar o seguindos")
			}

			following = following_loaded
		case publications_loaded := <-channel_publications:
			if len(publications_loaded) == 0 {
				publications = publications_loaded
			} else if publications_loaded == nil {
				return User{}, errors.New("erro ao buscar publicações")
			}

			publications = publications_loaded
		}
	}

	user.Followers = followers
	user.Following = following
	user.Publications = publications

	return user, nil
}

func Search_Data_User(channel chan<- User, user_id uint64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d", config.ApiURL, user_id)
	response, err := requests.Make_request_with_authentication(r, http.MethodGet, url, nil)

	if err != nil {
		channel <- User{}
		return
	}

	defer response.Body.Close()

	var user User

	if err = json.NewDecoder(response.Body).Decode(&user); err != nil {
		channel <- User{}
		return
	}

	channel <- user
}

func Search_Data_Followers(channel chan<- []User, user_id uint64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d/followers", config.ApiURL, user_id)
	response, err := requests.Make_request_with_authentication(r, http.MethodGet, url, nil)

	if err != nil {
		channel <- nil
		return
	}

	defer response.Body.Close()

	var followers []User

	if err = json.NewDecoder(response.Body).Decode(&followers); err != nil {
		channel <- nil
		return
	}

	if followers == nil {
		channel <- make([]User, 0)
	}

	channel <- followers
}

func Search_Data_Publications(channel chan<- []Publication, user_id uint64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d/publications", config.ApiURL, user_id)
	response, err := requests.Make_request_with_authentication(r, http.MethodGet, url, nil)

	if err != nil {
		channel <- nil
		return
	}

	defer response.Body.Close()

	var publications []Publication

	if err = json.NewDecoder(response.Body).Decode(&publications); err != nil {
		channel <- nil
		return
	}

	if publications == nil {
		channel <- make([]Publication, 0)
	}

	channel <- publications
}

func Search_Data_Following(channel chan<- []User, user_id uint64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d/following", config.ApiURL, user_id)
	response, err := requests.Make_request_with_authentication(r, http.MethodGet, url, nil)

	if err != nil {
		channel <- nil
		return
	}

	defer response.Body.Close()

	var following []User

	if err = json.NewDecoder(response.Body).Decode(&following); err != nil {
		channel <- nil
		return
	}

	if following == nil {
		channel <- make([]User, 0)
	}

	channel <- following
}
