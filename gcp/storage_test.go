package gcp

import (
        "testing"
        . "github.com/smartystreets/goconvey/convey"

        "github.com/quobix/lulaparty.io/model"
        "github.com/quobix/lulaparty.io/util"
        "fmt"
        "os"
        "google.golang.org/api/storage/v1"
        "net/http"

)

var bid         string
var ac          *model.AppConfig
var service     *storage.Service

var asset1      = "../test-assets/pic1.jpg"
var asset1fn    string

func TestMain(m *testing.M) {
        Setup()
        fmt.Fprintf(os.Stderr, "starting storage tests!\n")
        result := m.Run()
        fmt.Fprintf(os.Stderr, "finished storage tests!\n")
        Teardown()
        os.Exit(result)
}

func Setup() {

        ac = &model.AppConfig{ TestMode: true }

        bid = util.GenerateUUID()

}

func Teardown() {

}


func TestCreateStorageService(t *testing.T) {

        Convey("Given that we can configure a gcp client, we should be able to create a storage session.", t, func() {
                srv, err := CreateStorageService()
                So(err, ShouldBeNil)
                So(srv, ShouldNotBeNil)
                service = srv
        })

}

func TestCreateBucket(t *testing.T) {

        Convey("Given that we can create a storage session, we should be able to create a test bucket.", t, func() {

                bu, err := CreateBucket(bid,service, ac)

                So(err, ShouldBeNil)
                So(bu, ShouldNotBeNil)
                So(bu.Name, ShouldEqual, bid + "_test")

                //should result in the same test but it skips the create
                bu, err = CreateBucket(bid,service, ac)

                So(err, ShouldBeNil)
                So(bu, ShouldNotBeNil)
                So(bu.Name, ShouldEqual, bid + "_test")


        })

}

func TestListBuckets(t *testing.T) {

        Convey("Given that we can create a bucket, we should be able to get a list of them!", t, func() {
                bucks, err := ListBuckets(service, ac)
                So(bucks, ShouldNotBeNil)
                So(err, ShouldBeNil)
                So(len(bucks), ShouldBeGreaterThanOrEqualTo, 1);
                So(bucks[0].Name, ShouldEqual, bid + "_test")


        })

}

func TestUploadObjectToBucket(t *testing.T) {

        Convey("Given that we can create buckets, we should be able to upload objects to them!", t, func() {


                file, err := os.Open(asset1)
                So(err, ShouldBeNil)
                So(file, ShouldNotBeNil)

                ob, err := UploadObjectToBucket(bid, file, service, ac)
                So(err, ShouldBeNil)
                So(ob, ShouldNotBeNil)
                asset1fn = ob.Name


                ob, err = UploadObjectToBucket(bid, nil, service, ac) // no file, should error.
                So(err, ShouldNotBeNil)
                So(ob, ShouldBeNil)

        })

}


func TestGetObjectFromBucket(t *testing.T) {

        Convey("Given that we can upload objects to a bucket, we should be able to retrieve it", t, func() {

                ob, err := GetObjectFromBucket(bid, asset1fn, service, ac)

                So(err, ShouldBeNil)
                So(ob, ShouldNotBeNil)
                So(ob.Name, ShouldEqual, asset1fn)

                ob, err = GetObjectFromBucket(bid, "no-asset-here", service, ac)

                So(err, ShouldNotBeNil)
                So(ob, ShouldBeNil)

        })
}


func TestMakeObjectPublicReadable(t *testing.T) {

        Convey("Given that we can upload objects to a bucket, we should be able to make them public", t, func() {

                ob, err := GetObjectFromBucket(bid, asset1fn, service, ac)
                So(err, ShouldBeNil)
                So(ob, ShouldNotBeNil)

                uri :=GenerateObjectURI(bid, asset1fn, ac)

                So(uri, ShouldEqual, storageURIBase + "/" + bucketNameFilter(bid, ac) + "/" + asset1fn)
                // http://storage.googleapis.com/bid/asset1fn
                resp, err := http.Get(uri)

                So(err, ShouldBeNil)
                So(resp, ShouldNotBeNil)
                So(resp.StatusCode, ShouldEqual, 403) // forbidden
                fmt.Println("uri is " + GenerateObjectURI(bid, asset1fn, ac))


                err = MakeObjectPublicReadable(bid, asset1fn, service, ac)

                So(err, ShouldBeNil)

                resp, err = http.Get(GenerateObjectURI(bid, asset1fn, ac))

                So(err, ShouldBeNil)
                So(resp, ShouldNotBeNil)
                So(resp.StatusCode, ShouldNotEqual, 403)
                So(resp.StatusCode, ShouldEqual, 200)


                err = MakeObjectPublicReadable(bid, "no-asset-here", service, ac) // should fail
                So(err, ShouldNotBeNil)

        })
}


func TestDeleteObjectInBucket(t *testing.T) {

        Convey("Given that we can upload and publicly view the object, we should be able to delete it!", t, func() {

                err := DeleteObjectInBucket(bid, asset1fn, service, ac)

                So(err, ShouldBeNil)

                ob, err := GetObjectFromBucket(bid, asset1fn, service, ac)

                So(err, ShouldNotBeNil)
                So(ob, ShouldBeNil)

                err = DeleteObjectInBucket(bid, asset1fn, service, ac) // re-delete should fail
                So(err, ShouldNotBeNil)

        })
}



func TestDeleteBucket(t *testing.T) {

        Convey("Given that we can create a bucket, we should be able to delete it!.", t, func() {

                err := DeleteBucket(bid, service, ac)
                So(err, ShouldBeNil)

                dbu, err := GetBucket(bid, service, ac)
                So(dbu, ShouldBeNil)
                So(err, ShouldNotBeNil)

                err = DeleteBucket(bid, service, ac) // try and re-delete should throw an error
                So(err, ShouldNotBeNil)


        })

}
