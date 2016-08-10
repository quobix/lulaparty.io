
package controllers

import (
    "encoding/json"
    "net/http"
    "github.com/quobix/lulaparty.io/model"
    "github.com/quobix/lulaparty.io/service"
    //"strconv"
    "github.com/goinggo/tracelog"
    "net/mail"
    "github.com/gorilla/context"
)

func writeError(w http.ResponseWriter, msg string) {
    w.WriteHeader(http.StatusForbidden)
    json.NewEncoder(w).Encode(model.ServiceResponse{true,msg})
}


func Authenticate(w http.ResponseWriter, r *http.Request) {
    tracelog.Trace("controllers","Authenticate","Executing authentication controller")
    ac := context.Get(r, model.SYS_APPCONFIG).(*model.AppConfig)

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
    
    t  := service.ExchangeAccessToken(&ter)

    tracelog.Trace("controllers","Authenticate","Token exchange was successful")
    //tracelog.Trace("controllers","Authenticate","Token: " +  t.Token[0:8] + "...")
    //tracelog.Trace("controllers","Authenticate","Token Exp(secs): " + strconv.Itoa(t.ExpiryInSeconds))
    //tracelog.Trace("controllers","Authenticate","Token Exp(date): " + t.Expires.String())
    
    u, err := service.CheckForOrCreateExistingProfile(&ter, t, ac)

    jwt := service.GenerateToken(u)
    w.Header().Set("Content-Type", "application/json")
    h,_ := json.Marshal(jwt)
    
    w.WriteHeader(http.StatusOK)
    w.Write(h)
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






