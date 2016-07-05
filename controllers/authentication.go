
package controllers

import (
	"encoding/json"
	"net/http"
	"github.com/quobix/lulaparty.io/model"
	"github.com/quobix/lulaparty.io/service"
)

func Login(w http.ResponseWriter, r *http.Request) {
	requestUser := new(model.User)
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&requestUser)

	responseStatus, token := service.Login(requestUser)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(responseStatus)
	w.Write(token)
}

func RefreshToken(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	requestUser := new(model.User)
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&requestUser)

	w.Header().Set("Content-Type", "application/json")
	w.Write(service.RefreshToken(requestUser))
}

func Logout(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	err := service.Logout(r)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}