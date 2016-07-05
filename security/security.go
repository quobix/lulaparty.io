package security

import (
	"net/http"
	"os"
	"fmt"
	"strings"
	"time"
	"github.com/goinggo/tracelog"

)

func RequireTokenAuthentication(rw http.ResponseWriter, req *http.Request, next http.HandlerFunc) {



	tracelog.Trace("security","RequireTokenAuthentication","Checking http request for valid token...")
	m := CreateNewManager([]byte(os.Getenv("LLP_JWTSECRET")))
	header := strings.Replace(req.Header.Get("Authorization"),"Bearer ","",1)
	cl , err :=m.GetClaims(header)



	if err!= nil {
		fmt.Printf("%s", err)
	}

	if err == nil && cl != nil {
		tracelog.Trace("security","RequireTokenAuthentication","Valid token found, checking expiration....")
		tok_exp := time.Unix(cl.ExpiresAt,0)
		//now := time.Now()
		tracelog.Trace("security","RequireTokenAuthentication","token expires at "+ tok_exp.String())

		next(rw, req)
	} else {
		rw.WriteHeader(http.StatusUnauthorized)
	}
}