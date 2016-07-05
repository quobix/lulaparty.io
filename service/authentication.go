package service

import (
	"encoding/json"
	"net/http"
	"github.com/quobix/lulaparty.io/model"
	"os"
	"github.com/quobix/lulaparty.io/security"
)

type TokenAuthentication struct {
	Token string `json:"token" form:"token"`
}

func Login(requestUser *model.User) (int, []byte) {


	m := security.CreateNewManager([]byte(os.Getenv("LLP_JWTSECRET")))
	token := m.CreateToken(requestUser)
	str, err := m.SignString(token)

	if err!=nil  {
		return http.StatusUnauthorized, []byte("")
	}
	response, _ := json.Marshal(TokenAuthentication{str})
	return http.StatusOK, response


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