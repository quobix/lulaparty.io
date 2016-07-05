package data

import (
        "testing"
        "os"
        "strconv"
        . "github.com/smartystreets/goconvey/convey"
        "gopkg.in/mgo.v2"
        "fmt"
        "github.com/quobix/lulaparty.io/model"
        "gopkg.in/mgo.v2/bson"
)

var user        string
var pass        string
var host        string
var db          string
var jwtsec      string
var port        int
var sess        *mgo.Session
var ac          *model.AppConfig
var test_fbp    *model.FBProfile
var test_addr   *model.Address
var test_user   *model.User
var hostess1    *model.HostessProfile
var hostess2    *model.HostessProfile
var hostess3    *model.HostessProfile
var prov1       *model.ProviderProfile
var prov2       *model.ProviderProfile
var prov3       *model.ProviderProfile
var cust1       *model.CustomerProfile
var cust2       *model.CustomerProfile
var cust3       *model.CustomerProfile
var inv1        *model.Inventory
var inv2        *model.Inventory
var inv3        *model.Inventory
var invItem1    *model.InventoryItem
var invItem2    *model.InventoryItem
var invItem3    *model.InventoryItem
var invItem4    *model.InventoryItem
var invItem5    *model.InventoryItem
var g1          *model.Gallery
var g2          *model.Gallery
var g3          *model.Gallery
var gItem1      *model.GalleryItem
var gItem2      *model.GalleryItem
var gItem3      *model.GalleryItem
var p1,p2,p3    *model.Party
var pi1,pi2,pi3 *model.PartyInventory
var pinv1,pinv2 *model.PartyInvite
var rew1        *model.Reward

var uuid bson.ObjectId


func TestMain(m *testing.M) {
        Setup()
        fmt.Fprintf(os.Stderr, "starting data tests!\n")
        result := m.Run()
        fmt.Fprintf(os.Stderr, "finished data tests!\n")
//        Teardown()
        os.Exit(result)
}

func Setup() {
        host = os.Getenv("LLP_TEST_DB_HOST")
        pass = os.Getenv("LLP_TEST_DB_PASS")
        port, _ = strconv.Atoi(os.Getenv("LLP_TEST_DB_PORT"));
        user = os.Getenv("LLP_TEST_DB_USER")
        db = os.Getenv("LLP_TEST_DB")
        jwtsec = os.Getenv("LLP_JWTSECRET")

        test_fbp = &model.FBProfile {
                Id: bson.NewObjectId(),
                Firstname: "John",
                Lastname: "Appleseed",
                Email:  "john@appleseed.com" }

        test_addr = &model.Address {
                Id: bson.NewObjectId(),
                Street1: "1234 Happy Street",
                Street2: "Poptown",
                City:  "Smashville",
                State: "CA",
                Zip:    "90210",
        }

        test_user = &model.User {
                Id: bson.NewObjectId(),
                Cell: "(510) 321 3877",
                Email: "john@appleseed.com",
        }

        hostess1 = &model.HostessProfile {
                Id: bson.NewObjectId(),
        }
        hostess2 = &model.HostessProfile {
                Id: bson.NewObjectId(),
        }
        hostess3 = &model.HostessProfile {
                Id: bson.NewObjectId(),
        }
        prov1    = &model.ProviderProfile {
                Id: bson.NewObjectId(),
        }
        prov2    = &model.ProviderProfile {
                Id: bson.NewObjectId(),
        }
        prov3    = &model.ProviderProfile {
                Id: bson.NewObjectId(),
        }
        cust1    = &model.CustomerProfile {
                Id: bson.NewObjectId(),
        }
        cust2    = &model.CustomerProfile {
                Id: bson.NewObjectId(),
        }
        cust3    = &model.CustomerProfile {
                Id: bson.NewObjectId(),
        }
        inv1    = &model.Inventory {
                Id: bson.NewObjectId(),
        }
        inv2    = &model.Inventory {
                Id: bson.NewObjectId(),
        }
        inv3    = &model.Inventory {
                Id: bson.NewObjectId(),
        }

        invItem1    = &model.InventoryItem {
                Id: bson.NewObjectId(),
        }
        invItem2    = &model.InventoryItem {
                Id: bson.NewObjectId(),
        }
        invItem3    = &model.InventoryItem {
                Id: bson.NewObjectId(),
        }
        invItem4    = &model.InventoryItem {
                Id: bson.NewObjectId(),
        }

        g1    = &model.Gallery {
                Id: bson.NewObjectId(),
        }
        g2    = &model.Gallery {
                Id: bson.NewObjectId(),
        }
        g3    = &model.Gallery {
                Id: bson.NewObjectId(),
        }
        gItem1    = &model.GalleryItem {
                Id: bson.NewObjectId(),
        }
        gItem2    = &model.GalleryItem {
                Id: bson.NewObjectId(),
        }
        gItem3    = &model.GalleryItem {
                Id: bson.NewObjectId(),
        }
}

func Teardown() {
        sess.Close()
}

func TestGenerateURI(t *testing.T) {
        Convey("Given the ENV vars have been correctly set ", t, func() {

                Convey("The $LLP_TEST_DB_HOST env var should have been set", func() {
                        So(host, ShouldNotBeNil)
                        So(len(host), ShouldBeGreaterThan, 2)
                })
                Convey("The $LLP_TEST_DB_PASS env var should have been set", func() {
                        So(pass, ShouldNotBeNil)
                        So(len(pass), ShouldBeGreaterThan, 2)
                })
                Convey("The $LLP_TEST_DB_USER env var should have been set", func() {
                        So(user, ShouldNotBeNil)
                        So(len(user), ShouldBeGreaterThan, 2)
                })
                Convey("The $LLP_TEST_DB_PORT env var should have been set and greater than 0", func() {
                        So(port, ShouldNotBeNil)
                        So(port, ShouldBeGreaterThan, 0)
                })
                Convey("The $LLP_JWTSECRET env var should have been set", func() {
                        So(jwtsec, ShouldNotBeNil)
                        So(len(jwtsec), ShouldBeGreaterThan, 0)
                })

        })

        Convey("Given that we know how to connect to the DB", t, func() {

                var expected = "mongodb://" + user + ":" + pass + "@" + host + ":" + strconv.Itoa(port) + "/" + db
                Convey("The system connection URI for mongoDB should validate", func() {
                        So(expected, ShouldEqual, GenerateURI(user,pass,host,port,db))
                })

        })
}

func TestConnectDB(t *testing.T) {
        Convey("Given we have a valid URI and database credentials, we should be able to connect to the DB", t, func() {
                sess = ConnectDB(user,pass,host,port,db)
                So(sess, ShouldNotBeNil)

        })

        Convey("Given we have a valid URI and invalid database credentials, a connection should panic", t, func() {

                p := func() {
                        c := ConnectDB("inavlid","invalid",host,port,db)
                        defer c.Close()
                }
                So(p, ShouldPanic)
        })
}

func TestGetDB(t *testing.T) {
        d := GetDB(sess, db)
        Convey("Given a valid connection to the DB, we should be able to extract it from a session", t, func() {
                So(d, ShouldNotBeNil)
        })

}

func TestGetCollection(t *testing.T) {

        e := GetCollection(sess,db,"pop")

        Convey("Given a valid DB reference, we should be able to extract the collections", t, func() {
                So(e, ShouldNotBeNil)
        })
}

func TestGenerateCollectionName(t *testing.T) {
        Convey("Given we have a valid app config, flipping test mode on and off should generate correct names", t, func() {

                ac = &model.AppConfig{
                        DBName: db,
                        DBUser: user,
                        DBPassword: pass,
                        DBPort: port,
                        DBHost: host,
                        DBSession: ConnectDB(user, pass, host, port, db),
                        TestMode: true }

                So(ac, ShouldNotBeNil)
                So(GenerateCollectionName(ac, model.COLLECTION_FBPROFILE),
                        ShouldEqual, model.COLLECTION_FBPROFILE + model.COLLECTION_TEST_POSTFIX)
                So(GenerateCollectionName(ac, model.COLLECTION_USER),
                        ShouldEqual, model.COLLECTION_USER + model.COLLECTION_TEST_POSTFIX)
                So(GenerateCollectionName(ac, model.COLLECTION_ADDRESS),
                        ShouldEqual, model.COLLECTION_ADDRESS + model.COLLECTION_TEST_POSTFIX)

                So(GenerateCollectionName(ac, model.COLLECTION_ADDRESS),
                        ShouldNotEqual, model.COLLECTION_USER + model.COLLECTION_TEST_POSTFIX)

                ac.TestMode=false

                So(GenerateCollectionName(ac, model.COLLECTION_FBPROFILE),
                        ShouldEqual, model.COLLECTION_FBPROFILE)
                So(GenerateCollectionName(ac, model.COLLECTION_USER),
                        ShouldEqual, model.COLLECTION_USER)
                So(GenerateCollectionName(ac, model.COLLECTION_ADDRESS),
                        ShouldEqual, model.COLLECTION_ADDRESS)

                ac.TestMode=true // reset for everyone else.

        })

}


