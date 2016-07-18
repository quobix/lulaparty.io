
package controllers

import (
	"encoding/json"
	"net/http"
	"github.com/quobix/lulaparty.io/model"
	"github.com/quobix/lulaparty.io/service"
	"strconv"
	"github.com/goinggo/tracelog"
	"net/mail"
)

func writeError(w http.ResponseWriter, msg string) {
	w.WriteHeader(http.StatusForbidden)
	json.NewEncoder(w).Encode(model.ServiceResponse{true,msg})
}


func Authenticate(w http.ResponseWriter, r *http.Request) {
	tracelog.Trace("controllers","Authenticate","Executing authentication controller")

	var ter model.TokenExchangeRequest
	if err := json.NewDecoder(r.Body).Decode(&ter); err != nil {


		tracelog.Error(err, "controller","Token authentication failed, request malformed")
		writeError(w,"Token authentication failed, request malformed: " + err.Error())
		return
	}

	_, err := mail.ParseAddress(ter.Email)
	if err != nil {
		tracelog.Error(err, "controller","Token authentication failed, email invalid")
		writeError(w,"Token authentication failed, email invalid: " + err.Error())
		return
	}



	tracelog.Trace("controllers","Authenticate","Requesting token exchange for user [" + ter.Email + "]")
	tracelog.Trace("controllers","Authenticate","Exchanging short term token " + ter.AccessToken[0:8] + "...")

	t  := service.ExchangeAccessToken(&ter)

	tracelog.Trace("controllers","Authenticate","Token exchange was successful")
	tracelog.Trace("controllers","Authenticate","Token: " +  t.Token[0:8] + "...")
	tracelog.Trace("controllers","Authenticate","Token Exp(secs): " + strconv.Itoa(t.ExpiryInSeconds))


	tracelog.Trace("controllers","Authenticate","Token Exp(date): " + t.Expires.String())

	//responseStatus, token := service.Login(&ter)
	w.Header().Set("Content-Type", "application/json")
	//w.WriteHeader(responseStatus)
	w.Write([]byte("giggles"))
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






