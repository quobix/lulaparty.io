package gcp

import (
        "testing"
        . "github.com/smartystreets/goconvey/convey"
)

func TestCreateClient(t *testing.T) {

        Convey("Given that we have configured everything right, we should be able to create a GCP client.", t, func() {

                client, err := CreateClient(storageScope)
                So(err, ShouldBeNil)
                So(client, ShouldNotBeNil)
        })

}
