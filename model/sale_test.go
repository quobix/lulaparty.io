package model

import (
        "testing"
        . "github.com/smartystreets/goconvey/convey"
)

var sale = &Sale{}

func TestSetCreated_Sale(t *testing.T) {

        Convey("Given we have a sale, we should be able to test the creation date is correct", t, func() {
                helperSetCreated(sale)
        })
}

func TestUpdate_Sale(t *testing.T) {

        Convey("Given we sale, we should be able to test the last updated date is correct", t, func () {
                helperUpdated(sale)
        })
}

func TestGetId_Sale(t *testing.T) {

        Convey("Given we have a party inventory there should be an appropriate ID", t, func () {
                helperGetId(sale)
        })
}

func TestSetId_Sale(t *testing.T) {

        Convey("Given we have a persisted entity, we should be able to set an ID ", t, func () {
                helperSetId(&Sale{})
        })
}