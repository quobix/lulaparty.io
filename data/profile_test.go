package data

import (
        "testing"
        . "github.com/smartystreets/goconvey/convey"
        "gopkg.in/mgo.v2/bson"
        "github.com/quobix/lulaparty.io/model"
        "time"
)

func TestCreateCustomerProfile(t *testing.T) {

        if(ac == nil) {
                ac = CreateTestSession()
        }

        // forced db cleanup
        ac.DBSession.DB(ac.DBName).DropDatabase() // cleanup.


        Convey("Given that we have a customer profile, we should be able to persist it", t, func() {
                uuid = cust1.Id

                ret_cp, err := CreateCustomerProfile(cust1, ac)

                So(err, ShouldBeNil)
                So(ret_cp, ShouldNotBeNil)
                So(ret_cp.OwnerId.Hex(), ShouldEqual, "")

                // lets try and create an error
                _, err = CreateCustomerProfile(cust1, ac)
                So(err, ShouldNotBeNil) // already existing document
        })

}

func TestGetCustomerProfile(t *testing.T) {

        Convey("Given that we can persist the profile, we should be able to retrieve it as well", t, func() {

                p, err := GetCustomerProfile(cust1.Id, ac)

                So(err, ShouldBeNil)
                So(p, ShouldNotBeNil)
                So(p.Id, ShouldEqual, uuid)
                So(p.Id.Hex(), ShouldEqual, uuid.Hex())
                So(p.OwnerId.Hex(), ShouldEqual, "")

                _, err = GetProviderProfile(prov1.Id, ac) // check polluted ID's
                So(err, ShouldNotBeNil)

                _, err = GetCustomerProfile(bson.NewObjectId(), ac)
                So(err, ShouldNotBeNil)

        })

}

func TestUpdateCustomerProfile(t *testing.T) {

        Convey("Given that we can retrive the record, lets update it", t, func() {


                cust1.OwnerId =  test_user.Id

                ret_cp, err := UpdateCustomerProfile(cust1, ac)

                So(err, ShouldBeNil)
                So(ret_cp, ShouldNotBeNil)
                So(ret_cp.OwnerId.Hex(), ShouldNotEqual, "")
                So(ret_cp.OwnerId.Hex(), ShouldEqual, test_user.Id.Hex())
                So(ret_cp.Updated.After(ret_cp.Created), ShouldBeTrue)
                So(ret_cp.Created.Equal(ret_cp.Created), ShouldBeTrue)

                p, err := GetCustomerProfile(uuid, ac) // double check
                So(err, ShouldBeNil)
                So(p, ShouldNotBeNil)
                So(p.OwnerId.Hex(), ShouldEqual, test_user.Id.Hex())

                ret, err := UpdateCustomerProfile(&model.CustomerProfile{}, ac)
                So(err, ShouldNotBeNil)
                So(ret, ShouldBeNil)
        })

}

func TestCreateFBProfile(t *testing.T) {


        Convey("Given that we have a profile, we should be able to persist it", t, func() {
                uuid = test_fbp.Id

                ret_fbp, err := CreateFBProfile(test_fbp, ac)

                So(err, ShouldBeNil)
                So(ret_fbp, ShouldNotBeNil)
                So(ret_fbp.Firstname, ShouldEqual, "John")
                So(ret_fbp.Lastname, ShouldEqual, "Appleseed")
                So(ret_fbp.Email, ShouldEqual, "john@appleseed.com")

                // lets try and create an error
                ret_fbp, err = CreateFBProfile(test_fbp, ac)
                So(err, ShouldNotBeNil) // already existing document
        })

}

func TestGetFBProfile(t *testing.T) {

        Convey("Given that we can persist the profile, we should be able to retrieve it as well", t, func() {

                p, err := GetFBProfile(uuid, ac)

                So(err, ShouldBeNil)
                So(p, ShouldNotBeNil)
                So(p.Id, ShouldEqual, uuid)
                So(p.Id.Hex(), ShouldEqual, uuid.Hex())
                So(p.Email, ShouldEqual, test_fbp.Email)

                p, err = GetFBProfile(bson.NewObjectId(), ac)
                So(err, ShouldNotBeNil)

        })

}

func TestUpdateFBProfile(t *testing.T) {

        Convey("Given that we can retrive the record, lets update it", t, func() {

                test_fbp.Firstname = "David Thomas"
                test_fbp.Lastname  = "Shanley"
                test_fbp.Email =  "dave@quobix.com"

                ret_fbp, err := UpdateFBProfile(test_fbp, ac)

                So(err, ShouldBeNil)
                So(ret_fbp, ShouldNotBeNil)
                So(ret_fbp.Firstname, ShouldEqual, "David Thomas")
                So(ret_fbp.Lastname, ShouldEqual, "Shanley")
                So(ret_fbp.Email, ShouldEqual, "dave@quobix.com")
                So(ret_fbp.Updated.After(ret_fbp.Created), ShouldBeTrue)
                So(ret_fbp.Created.Equal(ret_fbp.Created), ShouldBeTrue)

                ret, err := UpdateFBProfile(&model.FBProfile{}, ac)
                So(err, ShouldNotBeNil)
                So(ret, ShouldBeNil)


        })

}

func TestAddAccessTokenToFBProfile(t *testing.T) {

        Convey("Given that we create a FB profile, we should be able to add an access token to it.", t, func() {

                at := &model.AccessToken {
                        Token: "aabbcc",
                        ExpiryInSeconds: 12345,
                        Expires: time.Now().UTC() }


                ret_fbp, err := AddAccessTokenToFBProfile(test_fbp, at, ac)

                So(err, ShouldBeNil)
                So(ret_fbp, ShouldNotBeNil)
                So(ret_fbp.AccessToken, ShouldNotBeNil)

                // refetch to check
                p, err := GetFBProfile(uuid, ac)
                So(err, ShouldBeNil)
                So(p, ShouldNotBeNil)
                So(p.AccessToken, ShouldNotBeNil)
                So(p.AccessToken.Hex(), ShouldEqual, at.Id.Hex())

                ret_at, err := GetAccessToken(p.AccessToken, ac)
                So(err, ShouldBeNil)
                So(ret_at, ShouldNotBeNil)
                So(ret_at.Token, ShouldEqual, "aabbcc")

        })

}



func TestCreateHostessProfile(t *testing.T) {

        Convey("Given that we have a hostess profile, we should be able to persist it", t, func() {
                uuid = hostess1.Id

                ret_hp, err := CreateHostessProfile(hostess1, ac)

                So(err, ShouldBeNil)
                So(ret_hp, ShouldNotBeNil)
                So(ret_hp.OwnerId.Hex(), ShouldEqual, "")

                // lets try and create an error
                ret_hp, err = CreateHostessProfile(hostess1, ac)
                So(err, ShouldNotBeNil) // already existing document
        })

}

func TestGetHostessProfile(t *testing.T) {

        Convey("Given that we can persist the profile, we should be able to retrieve it as well", t, func() {

                p, err := GetHostessProfile(hostess1.Id, ac)

                So(err, ShouldBeNil)
                So(p, ShouldNotBeNil)
                So(p.Id, ShouldEqual, uuid)
                So(p.Id.Hex(), ShouldEqual, uuid.Hex())
                So(p.OwnerId.Hex(), ShouldEqual, "")

                p, err = GetHostessProfile(bson.NewObjectId(), ac)
                So(err, ShouldNotBeNil)

        })

}

func TestUpdateHostessProfile(t *testing.T) {

        Convey("Given that we can retrive the record, lets update it", t, func() {


                hostess1.OwnerId =  test_user.Id

                ret_hp, err := UpdateHostessProfile(hostess1, ac)

                So(err, ShouldBeNil)
                So(ret_hp, ShouldNotBeNil)
                So(ret_hp.OwnerId.Hex(), ShouldNotEqual, "")
                So(ret_hp.OwnerId.Hex(), ShouldEqual, test_user.Id.Hex())
                So(ret_hp.Updated.After(ret_hp.Created), ShouldBeTrue)
                So(ret_hp.Created.Equal(ret_hp.Created), ShouldBeTrue)

                p, err := GetHostessProfile(uuid, ac) // double check
                So(err, ShouldBeNil)
                So(p, ShouldNotBeNil)
                So(p.OwnerId.Hex(), ShouldEqual, test_user.Id.Hex())

                ret, err := UpdateHostessProfile(&model.HostessProfile{}, ac)
                So(err, ShouldNotBeNil)
                So(ret, ShouldBeNil)

        })

}

func TestCreateProviderProfile(t *testing.T) {

        Convey("Given that we have a provider profile, we should be able to persist it", t, func() {
                uuid = prov1.Id

                ret_pp, err := CreateProviderProfile(prov1, ac)

                So(err, ShouldBeNil)
                So(ret_pp, ShouldNotBeNil)
                So(ret_pp.OwnerId.Hex(), ShouldEqual, "")

                // lets try and create an error
                ret_pp, err = CreateProviderProfile(prov1, ac)
                So(err, ShouldNotBeNil) // already existing document
        })

}

func TestGetProviderProfile(t *testing.T) {

        Convey("Given that we can persist the profile, we should be able to retrieve it as well", t, func() {

                p, err := GetProviderProfile(prov1.Id, ac)

                So(err, ShouldBeNil)
                So(p, ShouldNotBeNil)
                So(p.Id, ShouldEqual, uuid)
                So(p.Id.Hex(), ShouldEqual, uuid.Hex())
                So(p.OwnerId.Hex(), ShouldEqual, "")

                _, err = GetHostessProfile(prov1.Id, ac) // check polluted ID's
                So(err, ShouldNotBeNil)

                p, err = GetProviderProfile(bson.NewObjectId(), ac)
                So(err, ShouldNotBeNil)

        })

}

func TestUpdateProviderProfile(t *testing.T) {

        Convey("Given that we can retrive the record, lets update it", t, func() {


                prov1.OwnerId =  test_user.Id

                ret_pp, err := UpdateProviderProfile(prov1, ac)

                So(err, ShouldBeNil)
                So(ret_pp, ShouldNotBeNil)
                So(ret_pp.OwnerId.Hex(), ShouldNotEqual, "")
                So(ret_pp.OwnerId.Hex(), ShouldEqual, test_user.Id.Hex())
                So(ret_pp.Updated.After(ret_pp.Created), ShouldBeTrue)
                So(ret_pp.Created.Equal(ret_pp.Created), ShouldBeTrue)

                p, err := GetProviderProfile(uuid, ac) // double check
                So(err, ShouldBeNil)
                So(p, ShouldNotBeNil)
                So(p.OwnerId.Hex(), ShouldEqual, test_user.Id.Hex())

                ret, err := UpdateProviderProfile(&model.ProviderProfile{}, ac)
                So(err, ShouldNotBeNil)
                So(ret, ShouldBeNil)

        })

}
