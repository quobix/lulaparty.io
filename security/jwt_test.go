package security

import (
        "testing"
        . "github.com/smartystreets/goconvey/convey"

        "github.com/quobix/lulaparty.io/model"
        "github.com/quobix/lulaparty.io/data"
        "github.com/dgrijalva/jwt-go"

        "github.com/quobix/lulaparty.io/util"
)

var m   *Manager
var u   *model.User
var to  *jwt.Token
var sto string

func TestCreateNewManager(t *testing.T) {
        Convey("we should be able to create a new manager without error", t, func() {

                m = CreateNewManager([]byte(JWTSECRET))
                So(m, ShouldNotBeNil)
        })
}

func TestManager_CreateToken(t *testing.T) {

        Convey("given we have a valid user and a manager, we should be able to create a valid new jwt token", t, func() {
                u = &model.User{Email: "pop@shop.com" }

                ac := data.CreateTestSession()
                _, err := data.CreateUserSimple(u, ac)
                So(err, ShouldBeNil)

                to = m.CreateToken(u)

                So(to, ShouldNotBeNil)
                So(to.Claims, ShouldNotBeNil)
        })
}

func TestManager_SignByte(t *testing.T) {

        Convey("given we have a valid token, we shoud be able to create a signed bytearray", t, func() {
                s, err:= m.SignByte(to)

                So(s, ShouldNotBeNil)
                So(err, ShouldBeNil)
                So(len(s), ShouldBeGreaterThan, 10)
        })

}

func TestManager_SignString(t *testing.T) {

        Convey("given we have a valid token, we shoud be able to create a signed string", t, func() {
                s, err:= m.SignString(to)

                So(s, ShouldNotBeNil)
                So(err, ShouldBeNil)
                So(len(s), ShouldBeGreaterThan, 10)

                sto = s
        })
}

func TestManager_GetClaims(t *testing.T) {

        Convey("given we have a valid token string, we should be able to validate it and extrapolate lulu claims", t, func() {

                b,err :=m.GetClaims(sto)
                md5 := util.GetMD5Hash(u.Id.String()+u.Email+u.Created.String()) // all valid tokens contain this hash

                So(b, ShouldNotBeNil)
                So(err, ShouldBeNil)
                So(b.Lph, ShouldEqual, md5)
                So(b.Lpu, ShouldEqual, u.Id.Hex())

                b,err =m.GetClaims("something random, should completely fail!")
                So(b, ShouldBeNil)
                So(err, ShouldNotBeNil)
        })

}