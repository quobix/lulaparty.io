package model

import (
        "testing"
        . "github.com/smartystreets/goconvey/convey"
        "gopkg.in/mgo.v2/bson"
        "os"
        "path"
        "github.com/quobix/lulaparty.io/util"
)

var ga = &Gallery{}

func TestSetCreated_Gallery(t *testing.T) {

        Convey("Given we have a gallery, we should be able to test the creation date is correct", t, func() {
                helperSetCreated(ga)
        })
}

func TestUpdate_Gallery(t *testing.T) {

        Convey("Given we a gallery, we should be able to test the last updated date is correct", t, func () {
                helperUpdated(ga)
        })
}

func TestGetId_Gallery(t *testing.T) {

        Convey("Given we have a gallery there should be an appropriate ID", t, func () {
                helperGetId(ga)
        })
}

func TestSetId_Gallery(t *testing.T) {

        Convey("Given we have a persisted entity, we should be able to set an ID ", t, func () {
                helperSetId(&Gallery{})
        })
}


var gi = &GalleryItem{}

func TestSetCreated_GalleryItem(t *testing.T) {

        Convey("Given we have a gallery item, we should be able to test the creation date is correct", t, func() {
                helperSetCreated(gi)
        })
}

func TestUpdate_GalleryItem(t *testing.T) {

        Convey("Given we a style, we should be able to test the last updated date is correct", t, func () {
                helperUpdated(gi)
        })
}

func TestGetId_GalleryItem(t *testing.T) {

        Convey("Given we have a style there should be an appropriate ID", t, func () {
                helperGetId(gi)
        })
}

func TestSetId_GalleryItem(t *testing.T) {

        Convey("Given we have a persisted entity, we should be able to set an ID ", t, func () {
                helperSetId(&GalleryItem{})
        })
}

func TestGallery_ContainsItem(t *testing.T) {
        Convey("Given we have a have a valid gallery, check an ID can be found in the collection", t, func () {

                gi := bson.NewObjectId()
                gi_nope:= bson.NewObjectId()
                ga := &Gallery { Id: bson.NewObjectId() }

                ga.GalleryItems = append(ga.GalleryItems, gi)

                So(ga.ContainsItem(gi), ShouldBeTrue)
                So(ga.ContainsItem(gi_nope), ShouldBeFalse)


        })

}

func TestGenerateGalleryItemUUID(t *testing.T) {
        Convey("Given we have access to the OS, we should be able to read in a file " +
        "from the test-assets and validate the read was a success", t, func() {

                var asset1 = "../test-assets/pic1.jpg"
                file, err := os.Open(asset1)
                So(err, ShouldBeNil)
                So(file, ShouldNotBeNil)
                So(file.Name(), ShouldEqual, asset1)
                s, _ := file.Stat();
                So(s, ShouldNotBeNil)
                So(s.Size(), ShouldEqual, 345248)

                var u                   *User
                var g                   *Gallery
                var gi                  *GalleryItem


                u = &User {
                        Id: bson.NewObjectId(),
                }
                g = &Gallery {
                        Id: bson.NewObjectId(),
                }
                gi = &GalleryItem {
                        Id: bson.NewObjectId(),
                }

                Convey("We should then be able to generate an expected gallery item UUID for storage in gcp", func() {

                        ext     :=path.Ext(file.Name())
                        fn      :=path.Base(file.Name())
                        So(ext, ShouldEqual, ".jpg")
                        So(fn, ShouldEqual, "pic1.jpg")

                        gi_uuid := GenerateGalleryItemUUID(u, g, gi, file)

                        var expected = u.Id.Hex() + util.FILE_UUID_FSSEP + g.Id.Hex() +
                        util.FILE_UUID_FSSEP + gi.Id.Hex() + util.FILE_UUID_EXT + fn

                        So(expected, ShouldEqual, gi_uuid)
                })
        })
}
