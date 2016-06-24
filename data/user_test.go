package data

import (
        "testing"
        . "github.com/smartystreets/goconvey/convey"

        "gopkg.in/mgo.v2/bson"
        "github.com/quobix/lulaparty.io/model"
)

func TestCreateAddress(t *testing.T) {

        if(ac == nil) {
                ac = CreateTestSession()
        }

        Convey("Given that we have an address, we should be able to persist it", t, func() {
                uuid = test_addr.Id
                test_addr = &model.Address{
                        Id:uuid,
                        Street1: "1234 Happy Street",
                        Street2: "Poptown",
                        City:  "Smashville",
                        State: "CA",
                        Zip:    "90210",
                }


                ret_addr, err := CreateAddress(test_addr, ac)

                So(err, ShouldBeNil)
                So(ret_addr, ShouldNotBeNil)
                So(ret_addr.Street1, ShouldEqual, "1234 Happy Street")
                So(ret_addr.Street2, ShouldEqual, "Poptown")
                So(ret_addr.City, ShouldEqual, "Smashville")
                So(ret_addr.State, ShouldEqual, "CA")
                So(ret_addr.Zip, ShouldEqual, "90210")
                So(ret_addr.Created, ShouldNotBeNil)

                //check that we can't duplicate record
                ret_addr, err = CreateAddress(test_addr, ac)
                So(err, ShouldNotBeNil)

        })

}

func TestGetAddress(t *testing.T) {

        Convey("Given that we can persist the address, we should be able to retrieve it as well", t, func() {

                a, err := GetAddress(uuid, ac)

                So(err, ShouldBeNil)
                So(a, ShouldNotBeNil)
                So(a.Id, ShouldEqual, uuid)
                So(a.Id.Hex(), ShouldEqual, uuid.Hex())
                So(a.Street1, ShouldEqual, test_addr.Street1)
                So(a.Street2, ShouldEqual, test_addr.Street2)
                So(a.City, ShouldEqual, test_addr.City)
                So(a.State, ShouldEqual, test_addr.State)
                So(a.Zip, ShouldEqual, test_addr.Zip)

                a, err = GetAddress(bson.NewObjectId(), ac)
                So(err, ShouldNotBeNil)

        })

}

func TestUpdateAddress(t *testing.T) {

        Convey("Given that we can retrive the record, lets update it", t, func() {

                test_addr.Street1 = "998 Sleepy Hill"
                test_addr.Street2  = "The Sliders"
                test_addr.City =  "Pootyville"
                test_addr.State =  "VA"
                test_addr.Zip =  "25045"

                ret_addr, err := UpdateAddress(test_addr, ac)

                So(err, ShouldBeNil)
                So(ret_addr, ShouldNotBeNil)
                So(ret_addr.Street1, ShouldEqual, "998 Sleepy Hill")
                So(ret_addr.Street2, ShouldEqual, "The Sliders")
                So(ret_addr.City, ShouldEqual, "Pootyville")
                So(ret_addr.State, ShouldEqual, "VA")
                So(ret_addr.Zip, ShouldEqual, "25045")
                So(ret_addr.Updated.After(ret_addr.Created), ShouldBeTrue)
                So(ret_addr.Created.Equal(ret_addr.Created), ShouldBeTrue)


        })

}


func TestCreateUser(t *testing.T) {

        // forced db cleanup
        ac.DBSession.DB(ac.DBName).DropDatabase() // cleanup.

        Convey("Given that we have a user, a profile and an address, we can persist a new user", t, func() {
                uuid = test_user.Id


                ret_user, err := CreateUser(test_user, test_fbp, test_addr, ac)

                So(err, ShouldBeNil)
                So(ret_user, ShouldNotBeNil)
                So(ret_user.Cell, ShouldEqual, "(510) 321 3877")

                p, err := GetFBProfile(ret_user.FBProfile, ac)

                So(err, ShouldBeNil)
                So(p, ShouldNotBeNil)
                So(p.Firstname, ShouldEqual, test_fbp.Firstname)
                So(p.Lastname, ShouldEqual, test_fbp.Lastname)

                a, err := GetAddress(ret_user.Address, ac)

                So(err, ShouldBeNil)
                So(a, ShouldNotBeNil)
                So(a.State, ShouldEqual, test_addr.State)
                So(a.Street1, ShouldEqual, test_addr.Street1)
                So(a.Street2, ShouldEqual, test_addr.Street2)
                So(a.City, ShouldEqual, test_addr.City)

                // check we can't duplicate the user
                ret_user, err = CreateUser(test_user, test_fbp, test_addr, ac)
                So(err, ShouldNotBeNil)

        })

}

func TestCreateUserSimple(t *testing.T) {

        Convey("Given that we have a user, we can persist a new user without the required profiles", t, func() {


                test_user =&model.User{ }
                ret_user, err := CreateUserSimple(test_user, ac)

                So(err, ShouldBeNil)
                So(ret_user, ShouldNotBeNil)

                ret_user, err = CreateUserSimple(test_user, ac)

                So(err, ShouldNotBeNil)
                So(ret_user, ShouldBeNil)

        })

}

func TestGetUser(t *testing.T) {

        Convey("Given that we can persist the user, we should be able to retreive it", t, func() {

                a, err := GetUser(uuid, ac)
                So(err, ShouldBeNil)
                So(a, ShouldNotBeNil)
                So(a.Id, ShouldEqual, uuid)
                So(a.Id.Hex(), ShouldEqual, uuid.Hex())

                a, err = GetUser(bson.NewObjectId(), ac)
                So(err, ShouldNotBeNil)

        })

}


func TestUpdateUser(t *testing.T) {

        Convey("Given that we can find our users, lets verify we can update them", t, func() {


                test_user.Email = "pop@chop.com"
                test_user.Cell  = "911-999-111"

                ret_user, err := UpdateUser(test_user, ac)

                So(err, ShouldBeNil)
                So(ret_user, ShouldNotBeNil)
                So(ret_user.Cell, ShouldEqual, test_user.Cell)
                So(ret_user.Email, ShouldEqual,test_user.Email)


        })

}

func TestGetUserByEmail(t *testing.T) {

        Convey("Given that we can find the user by an ID, lets verify we can search for an email also", t, func() {

                u, err := GetUserByEmail(test_user.Email, ac)
                So(err, ShouldBeNil)
                So(u, ShouldNotBeNil)
                So(u.Cell, ShouldEqual, "911-999-111")
                So(u.Email, ShouldEqual, test_user.Email)

                u, err = GetUserByEmail("test@fail.com", ac)
                So(err, ShouldNotBeNil)

        })

}



