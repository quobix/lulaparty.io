package service

import (
	//"encoding/json"
	"net/http"
	"github.com/quobix/lulaparty.io/model"
	"os"
	//"github.com/quobix/lulaparty.io/security"
	"strings"
	"strconv"
	"time"
	"github.com/quobix/lulaparty.io/util"
	//"lulaparty.io/service"
	"github.com/quobix/lulaparty.io/data"
	//"gopkg.in/mgo.v2"
	"github.com/goinggo/tracelog"

)

type TokenAuthentication struct {
	Token string `json:"token" form:"token"`
}

func AuthenticateUser(ter *model.TokenExchangeRequest, ac *model.AppConfig) (int, []byte) {

	//m := security.CreateNewManager([]byte(os.Getenv("LLP_JWTSECRET")))
	_, err := data.GetUserByEmail(ter.Email, ac)
	if(err != nil) {
		tracelog.Trace("service","AuthenticateUser","New User! Not Found! Biscuits: [" + ter.Email + "]")
	}


/*


	token := m.CreateToken(requestUser)
	str, err := m.SignString(token)

	if err!=nil  {
		return http.StatusUnauthorized, []byte("")
	}
	response, _ := json.Marshal(TokenAuthentication{str})
	return http.StatusOK, response
*/
	return 0, nil

}

func RefreshToken(requestUser *model.User) []byte {
	/*
	authBackend := authentication.InitJWTAuthenticationBackend()
	token, err := authBackend.GenerateToken(requestUser.UUID)
	if err != nil {
		panic(err)
	}
	response, err := json.Marshal(parameters.TokenAuthentication{token})
	if err != nil {
		panic(err)
	}
	return response
	*/
	return nil
}

func Logout(req *http.Request) error {

	/*
	authBackend := authentication.InitJWTAuthenticationBackend()
	tokenRequest, err := jwt.ParseFromRequest(req, func(token *jwt.Token) (interface{}, error) {
		return authBackend.PublicKey, nil
	})
	if err != nil {
		return err
	}
	tokenString := req.Header.Get("Authorization")
	return authBackend.Logout(tokenString, tokenRequest)

	*/
	return nil
}

func ExchangeAccessToken(ter *model.TokenExchangeRequest) *model.AccessToken {

	cid := os.Getenv("LLP_FBAPPID")
	sec := os.Getenv("LLP_FBSECRET")
	r, err := http.Get(OAUTH_URI +
	"?client_id=" + cid +
	"&grant_type=fb_exchange_token&" +
	"&client_secret=" + sec +
	"&fb_exchange_token=" + ter.AccessToken)

	if err == nil {
		auth := util.ReadBody(r)
		var token model.AccessToken

		tokenArr := strings.Split(auth, "&")

		token.Token = strings.Split(tokenArr[0], "=")[1]
		expireInt, err := strconv.Atoi(strings.Split(tokenArr[1], "=")[1])

		if err == nil {
			token.ExpiryInSeconds = int(expireInt)
			ti :=time.Now().UTC()
			token.Expires = time.Date(ti.Year(), ti.Month(), ti.Day(),
				ti.Hour(), ti.Minute(), ti.Second()+token.ExpiryInSeconds,
				ti.Nanosecond(), time.UTC)
		}
		return &token
	}
	return new(model.AccessToken)
}
