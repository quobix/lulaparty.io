package model

import (
        "testing"
        . "github.com/smartystreets/goconvey/convey"
        "github.com/shopspring/decimal"

        "gopkg.in/mgo.v2/bson"
)

var inv = &Inventory{}

func TestSetCreated_Inventory(t *testing.T) {

        Convey("Given we have an inventory, we should be able to test the creation date is correct", t, func() {
                helperSetCreated(inv)
        })
}

func TestUpdate_Inventory(t *testing.T) {

        Convey("Given we have an inventory, we should be able to test the last updated date is correct", t, func () {
                helperUpdated(inv)
        })
}

func TestGetId_Inventory(t *testing.T) {

        Convey("Given we haven an inventory there should be an appropriate ID", t, func () {
                helperGetId(inv)
        })
}

func TestSetId_Inventory(t *testing.T) {

        Convey("Given we have a persisted entity, we should be able to set an ID ", t, func () {
                helperSetId(&Inventory{})
        })
}

var invis = &InventoryItemSizes{}

func TestSetCreated_InventoryItemSizes(t *testing.T) {

        Convey("Given we have nventory item sizes, we should be able to test the creation date is correct", t, func() {
                helperSetCreated(invis)
        })
}

func TestUpdate_InventoryItemSizes(t *testing.T) {

        Convey("Given we have inventory item sizes, we should be able to test the last updated date is correct", t, func () {
                helperUpdated(invis)
        })
}

func TestGetId_InventoryItemSizes(t *testing.T) {

        Convey("Given we haven inventory item sizes there should be an appropriate ID", t, func () {
                helperGetId(invis)
        })
}

func TestSetId_InventoryItemSizes(t *testing.T) {

        Convey("Given we have a persisted entity, we should be able to set an ID ", t, func () {
                helperSetId(&InventoryItemSizes{})
        })
}

var invi = &InventoryItem{}

func TestSetCreated_InventoryItem(t *testing.T) {

        Convey("Given we have an inventory item, we should be able to test the creation date is correct", t, func() {
                helperSetCreated(invi)
        })
}

func TestUpdate_InventoryItem(t *testing.T) {

        Convey("Given we have an inventor item, we should be able to test the last updated date is correct", t, func () {
                helperUpdated(invi)
        })
}

func TestGetId_InventoryItem(t *testing.T) {

        Convey("Given we haven an inventory item there should be an appropriate ID", t, func () {
                helperGetId(invi)
        })
}

func TestSetId_InventoryItem(t *testing.T) {

        Convey("Given we have a persisted entity, we should be able to set an ID ", t, func () {
                helperSetId(&InventoryItem{})
        })
}

var style = &Style{}

func TestSetCreated_Style(t *testing.T) {

        Convey("Given we have a style, we should be able to test the creation date is correct", t, func() {
                helperSetCreated(style)
        })
}

func TestUpdate_Style(t *testing.T) {

        Convey("Given we a style, we should be able to test the last updated date is correct", t, func () {
                helperUpdated(style)
        })
}

func TestGetId_Style(t *testing.T) {

        Convey("Given we have a style there should be an appropriate ID", t, func () {
                helperGetId(style)
        })
}

func TestSetId_Style(t *testing.T) {

        Convey("Given we have a persisted entity, we should be able to set an ID ", t, func () {
                helperSetId(&Style{})
        })
}

func TestPrice_GetBSON(t *testing.T) {

        Convey("We should be able to return the BSON after serialization", t, func () {
                d, _ := decimal.NewFromString("123.22")
                p := &Price { Value: d}
                b, _ := p.GetBSON()

                So(b, ShouldNotBeNil)

        })
}

func TestPrice_SetBSON(t *testing.T) {

        Convey("We should be able to serialize the BSON data", t, func () {
                d, _ := decimal.NewFromString("123.22")
                p := &Price { Value: d}
                err :=  p.SetBSON(bson.Raw{})

                So(err, ShouldNotBeNil)


        })
}