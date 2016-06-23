package model

import (
        "testing"
        . "github.com/smartystreets/goconvey/convey"
)

var cp = &CustomerProfile{}

func TestSetCreated_CustomerProfile(t *testing.T) {

        Convey("Given we have a customner profile, we should be able to test the creation date is correct", t, func() {
                helperSetCreated(cp)
        })
}

func TestUpdate_CustomerProfile(t *testing.T) {

        Convey("Given we have a customer profile, we should be able to test the last updated date is correct", t, func () {
                helperUpdated(cp)
        })
}

func TestGetId_CustomerProfile(t *testing.T) {

        Convey("Given we have customer profile there should be an appropriate ID", t, func () {
                helperGetId(cp)
        })
}

func TestSetId_CustomerProfile(t *testing.T) {

        Convey("Given we have a persisted entity, we should be able to set an ID ", t, func () {
                helperSetId(&CustomerProfile{})
        })
}

var hp = &HostessProfile{}

func TestSetCreated_HostessProfile(t *testing.T) {

        Convey("Given we have a hostess profile, we should be able to test the creation date is correct", t, func() {
                helperSetCreated(hp)
        })
}

func TestUpdate_HostessProfile(t *testing.T) {

        Convey("Given we have a hostess profile, we should be able to test the last updated date is correct", t, func () {
                helperUpdated(hp)
        })
}

func TestGetId_HostessProfile(t *testing.T) {

        Convey("Given we have hostess profile there should be an appropriate ID", t, func () {
                helperGetId(hp)
        })
}

func TestSetId_HostessProfile(t *testing.T) {

        Convey("Given we have a persisted entity, we should be able to set an ID ", t, func () {
                helperSetId(&HostessProfile{})
        })
}


func TestSetId_ProviderProfile(t *testing.T) {

        Convey("Given we have a persisted entity, we should be able to set an ID ", t, func () {
                helperSetId(&ProviderProfile{})
        })
}

var pp = &ProviderProfile{}

func TestSetCreated_ProviderProfile(t *testing.T) {

        Convey("Given we have a providerprofile, we should be able to test the creation date is correct", t, func() {
                helperSetCreated(pp)
        })
}

func TestUpdate_ProviderProfile(t *testing.T) {

        Convey("Given we have a provider profile, we should be able to test the last updated date is correct", t, func () {
                helperUpdated(pp)
        })
}

func TestGetId_ProviderProfile(t *testing.T) {

        Convey("Given we have provider profile there should be an appropriate ID", t, func () {
                helperGetId(pp)
        })
}


var fbp = &FBProfile{}

func TestSetCreated_FBProfile(t *testing.T) {

        Convey("Given we have a fbprofile, we should be able to test the creation date is correct", t, func() {
                helperSetCreated(fbp)
        })

}

func TestUpdate_FBProfile(t *testing.T) {

        Convey("Given we have a fbprofile, we should be able to test the last updated date is correct", t, func () {
                helperUpdated(fbp)
        })
}

func TestGetId_FBProfile(t *testing.T) {

        Convey("Given we have fbprofile there should be an appropriate ID", t, func () {
                helperGetId(fbp)
        })
}

func TestSetId_FBProfile(t *testing.T) {

        Convey("Given we have a persisted entity, we should be able to set an ID ", t, func () {
                helperSetId(&FBProfile{})
        })
}

