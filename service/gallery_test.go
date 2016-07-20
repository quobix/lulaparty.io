package service

import (
        "testing"
        . "github.com/smartystreets/goconvey/convey"
        "github.com/quobix/lulaparty.io/model"

        "github.com/quobix/lulaparty.io/data"
        "fmt"
        "os"
        "gopkg.in/mgo.v2/bson"
        "github.com/quobix/lulaparty.io/gcp"
        "net/http"
        "google.golang.org/api/storage/v1"
        "time"
)

var ac          *model.AppConfig
var service     *storage.Service
var asset1      = "../test-assets/pic1.jpg"

var galItem          *model.GalleryItem
var gal              *model.Gallery

var galId bson.ObjectId
var galItemId bson.ObjectId

func TestMain(m *testing.M) {
        Setup()
        fmt.Fprintf(os.Stderr, "starting service tests!\n")
        result := m.Run()
        fmt.Fprintf(os.Stderr, "finished service tests!\n")
        Teardown()
        os.Exit(result)
}

func Setup() {
        ac = data.CreateTestSession()
        service, _ = gcp.CreateStorageService()
}

func Teardown() {

}

func TestPersistGalleryItemToStorage(t *testing.T) {
        var fuuid string

        Convey("Given we are able to store a gallery item, we should be able to persist an object and wire the references", t, func () {
                _, err := gcp.CreateBucket(model.BUCKET_GALLERY, service, ac)

                file, err := os.Open(asset1)

                _g := &model.Gallery {
                        Id: bson.NewObjectId(),
                        OwnerId: bson.NewObjectId() }

                _gi := &model.GalleryItem {
                        Id: bson.NewObjectId(),
                        GalleryId: _g.Id,
                        OwnerId: _g.OwnerId }

                fmt.Fprintf(os.Stderr, "starting gallery persist\n")

                gal, _ = data.CreateGallery(_g, ac)
                galItem = _gi
                galId = gal.Id
                galItemId = galItem.Id

                ret_gi, err := PersistGalleryItemToStorage(galItem, file, ac)

                So(err, ShouldBeNil)
                So(ret_gi, ShouldNotBeNil)

                time.Sleep(500 * time.Millisecond) // we need a second here to let GCP catch up

                gi1, gierr := data.GetGalleryItem(ret_gi.Id, ac)
                fuuid=gi1.FileUUID

                So(gierr, ShouldBeNil)
                So(gi1, ShouldNotBeNil)
                So(fuuid, ShouldEqual, galItem.FileUUID)
        })

        Convey("And given we can wire and persist, we should be able to verify the objects exist and are correctly wired", t, func () {
                resp, err := http.Get(gcp.GenerateObjectURI(model.BUCKET_GALLERY, fuuid, ac))

                So(err, ShouldBeNil)
                So(resp, ShouldNotBeNil)
                So(resp.StatusCode, ShouldNotEqual, 403)
                So(resp.StatusCode, ShouldEqual, 200)

                _g, err := data.GetGallery(galId, ac)
                So(err, ShouldBeNil)
                So(_g.ContainsItem(galItemId), ShouldBeTrue)
        })

}

func TestRemoveGalleryItemFromStorage(t *testing.T) {

        Convey("Given we have a valid persisted gallery item, we should be able to tear it down, remove and verify", t, func () {

                _g, err := data.GetGallery(galId, ac)
                So(err, ShouldBeNil)
                So(_g.ContainsItem(galItemId), ShouldBeTrue)

                gi1, gierr := data.GetGalleryItem(galItemId, ac)
                So(gierr, ShouldBeNil)
                So(gi1, ShouldNotBeNil)
                So(gi1.FileUUID, ShouldNotBeNil)
                So(gi1.Id, ShouldEqual, galItemId)
                So(gi1.GalleryId, ShouldEqual, galId)

                err = RemoveGalleryItemFromStorage(gi1, ac)
                So(err, ShouldBeNil)

                _g, err = data.GetGallery(_g.Id, ac)
                So(err, ShouldBeNil)
                So(_g.ContainsItem(gi1.Id), ShouldBeFalse)

                err = RemoveGalleryItemFromStorage(gi1, ac)
                So(err, ShouldNotBeNil)

                err = RemoveGalleryItemFromStorage(&model.GalleryItem { Id: bson.NewObjectId() }, ac)
                So(err, ShouldNotBeNil)

                err = data.DeleteGallery(_g, ac)
                So(err, ShouldBeNil)

                err = gcp.DeleteBucket(model.BUCKET_GALLERY, service, ac)
                So(err, ShouldBeNil)

        })
}
