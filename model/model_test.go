package model

import (
        "testing"
        . "github.com/smartystreets/goconvey/convey"
        "os"
        "strconv"
        "gopkg.in/mgo.v2"
        "time"
        "gopkg.in/mgo.v2/bson"
)

func TestAppConfig_CopyDBSession(t *testing.T) {
        host := os.Getenv("LLP_TEST_DB_HOST")
        pass := os.Getenv("LLP_TEST_DB_PASS")
        port, _ := strconv.Atoi(os.Getenv("LLP_TEST_DB_PORT"));
        user := os.Getenv("LLP_TEST_DB_USER")
        db := os.Getenv("LLP_TEST_DB")

        Convey("Given that we have a valid app config ", t, func() {

                ac := AppConfig {
                        DBName: db,
                        DBUser: user,
                        DBPassword: pass,
                        DBPort: port,
                        DBHost: host,
                        DBSession: nil }

                Convey("There should be a panic if the DB session is empty", func() {
                        So(func() { ac.CopyDBSession() }, ShouldPanic)

                        ac.DBSession = &mgo.Session{}
                        So(func() { ac.CopyDBSession() }, ShouldPanic)
                })

        })

}

func helperSetCreated(n PersistedEntity) {
        p := n.SetCreated()
        So(p, ShouldNotBeNil)
        So(p.Before(time.Now()), ShouldBeTrue)
}

func helperUpdated(n PersistedEntity) {
        p := n.Update()
        So(p, ShouldNotBeNil)
        So(p.Before(time.Now()), ShouldBeTrue)
}

func helperGetId(n PersistedEntity) {
        p := n.GetId()
        So(p, ShouldNotBeNil)
        So(p.String(), ShouldNotBeNil)
        So(p.Valid(), ShouldBeFalse)
        So(p.Hex(), ShouldNotEqual, bson.NewObjectId().Hex())
}

func helperSetId(n PersistedEntity) {
        n.SetId(bson.NewObjectId())
        p := n.GetId()
        So(p, ShouldNotBeNil)
        So(p.String(), ShouldNotBeNil)
        So(p.Hex(), ShouldNotEqual, bson.NewObjectId().Hex())
}