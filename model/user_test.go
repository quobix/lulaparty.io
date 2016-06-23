package model

import (
        "testing"
        . "github.com/smartystreets/goconvey/convey"
)

var u = &User{}

func TestSetCreated_User(t *testing.T) {

        Convey("Given we have a user, we should be able to test the creation date is correct", t, func() {
                helperSetCreated(u)
        })

}

func TestUpdate_User(t *testing.T) {

        Convey("Given we have a user, we should be able to test the last updated date is correct", t, func () {
                helperUpdated(u)
        })
}

func TestGetId_User(t *testing.T) {

        Convey("Given we have a user, there should be an appropriate ID", t, func () {
                helperGetId(u)
        })
}

func TestSetId_User(t *testing.T) {

        Convey("Given we have a persisted entity, we should be able to set an ID ", t, func () {
                helperSetId(&User{})
        })
}



var a = &Address{}

func TestSetCreated_Address(t *testing.T) {

        Convey("Given we have an address, we should be able to test the creation date is correct", t, func() {
                helperSetCreated(a)
        })

}

func TestUpdate_Address(t *testing.T) {

        Convey("Given we have an address, we should be able to test the last updated date is correct", t, func () {
                helperUpdated(a)
        })
}

func TestGetId_Address(t *testing.T) {

        Convey("Given we have an address there should be an appropriate ID", t, func () {
                helperGetId(a)
        })
}

func TestSetId_Address(t *testing.T) {

        Convey("Given we have a persisted entity, we should be able to set an ID ", t, func () {
                helperSetId(&Address{})
        })
}