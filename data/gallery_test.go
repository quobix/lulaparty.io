package data

import (
        "testing"
        . "github.com/smartystreets/goconvey/convey"
        "github.com/quobix/lulaparty.io/model"
        "gopkg.in/mgo.v2/bson"

)



func TestCreateGallery(t *testing.T) {

        if(ac == nil) {
                ac = CreateTestSession()
        }

        Convey("Given that we have gallery, we should be able to persist it", t, func() {
                uuid = g1.Id
                g := &model.Gallery {
                        Id:g1.Id,
                        OwnerId: test_user.Id,
                }


                ret_g, err := CreateGallery(g, ac)

                So(err, ShouldBeNil)
                So(ret_g.Created, ShouldNotBeNil)

                //check that we can't duplicate record
                g, err = CreateGallery(g, ac)
                So(err, ShouldNotBeNil)



        })

}

func TestGetGallery(t *testing.T) {

        Convey("Given that we can persist the gallery, we should be able to retreive it", t, func() {

                g, err := GetGallery(g1.Id, ac)
                So(err, ShouldBeNil)
                So(g, ShouldNotBeNil)
                So(g.Id, ShouldEqual, uuid)
                So(g.Id.Hex(), ShouldEqual, uuid.Hex())

                g, err = GetGallery(bson.NewObjectId(), ac)
                So(err, ShouldNotBeNil)

        })


}

func TestUpdateGallery(t *testing.T) {

        Convey("Given that we can retrive the record, lets update it", t, func() {

                g1.Name = "Dave's Gallery"
                ret_g, err := UpdateGallery(g1, ac)

                So(err, ShouldBeNil)
                So(ret_g, ShouldNotBeNil)
                So(ret_g.Name, ShouldEqual, g1.Name)

                // double check
                g, err := GetGallery(g1.Id, ac)
                So(err, ShouldBeNil)
                So(g, ShouldNotBeNil)
                So(g.Name, ShouldEqual, g1.Name)

                bad, err := GetGallery(bson.NewObjectId(), ac)
                So(bad, ShouldBeNil)
                So(err, ShouldNotBeNil)

                bad, err = UpdateGallery(&model.Gallery{}, ac)
                So(err, ShouldNotBeNil)
                So(bad, ShouldBeNil)


        })

}


func TestAddGalleryItemToGallery(t *testing.T) {

        //make sure our gallery items have an owner!
        gItem1.OwnerId = test_user.Id
        gItem2.OwnerId = test_user.Id
        gItem3.OwnerId = test_user.Id

        gItem1.Name="Gallery Item 1"
        gItem2.Name="Gallery Item 2"
        gItem3.Name="Gallery Item 3"

        gItem1.GalleryId = g1.Id;
        gItem2.GalleryId = g1.Id;
        gItem3.GalleryId = g1.Id;



        Convey("Given that we have a gallery, we should be able to add an item without persisting it", t, func() {

                g, err := AddItemToGallery(g1.Id, gItem1, ac)
                So(err, ShouldBeNil)
                So(g, ShouldNotBeNil)
                So(len(g.GalleryItems), ShouldEqual, 1)

        })

        Convey("We should be able to add another couple of items to the gallery", t, func() {

                g, err := AddItemToGallery(g1.Id, gItem2, ac)
                So(err, ShouldBeNil)
                So(g, ShouldNotBeNil)
                So(len(g.GalleryItems), ShouldEqual, 2)

                g, err = AddItemToGallery(g1.Id, gItem3, ac)
                So(err, ShouldBeNil)
                So(g, ShouldNotBeNil)
                So(len(g.GalleryItems), ShouldEqual, 3)

        })

        Convey("Trying to add a duplicate gallery item should result in an error", t, func() {

                it, err := AddItemToGallery(g1.Id, gItem2, ac)
                So(err, ShouldNotBeNil)
                So(it, ShouldBeNil)

        })

}

func TestGetGalleryItem(t *testing.T) {

        Convey("Given that we can persist gallery items, we should be able to retrive them", t, func() {

                gi, err := GetGalleryItem(gItem1.Id, ac)
                So(err, ShouldBeNil)
                So(gi, ShouldNotBeNil)
                So(gi.Id, ShouldEqual, gItem1.Id)
                So(gi.Id.Hex(), ShouldEqual, gItem1.Id.Hex())

                gi, err = GetGalleryItem(bson.NewObjectId(), ac)
                So(err, ShouldNotBeNil)

        })
}


func TestUpdateGalleryItem(t *testing.T) {

        Convey("Given that we can retrive the gallery item, we should be able to update it", t, func() {

                gItem1.Caption ="rum in your bum"
                gItem2.Caption ="bum in your rum"
                gItem3.Caption ="plum in my bum"
                gItem1.FileUUID = "aft2345"

                ret_g1, err1 := UpdateGalleryItem(gItem1, ac)
                ret_g2, err2 := UpdateGalleryItem(gItem2, ac)
                ret_g3, err3 := UpdateGalleryItem(gItem3, ac)

                So(err1, ShouldBeNil)
                So(err2, ShouldBeNil)
                So(err3, ShouldBeNil)
                So(ret_g1, ShouldNotBeNil)
                So(ret_g1.Caption, ShouldEqual, "rum in your bum")
                So(ret_g2.Caption, ShouldEqual, "bum in your rum")
                So(ret_g3.Caption, ShouldEqual, "plum in my bum")

                // double check
                g, err := GetGalleryItem(gItem1.Id, ac)
                So(err, ShouldBeNil)
                So(g, ShouldNotBeNil)
                So(g.Caption, ShouldEqual, "rum in your bum")
                So(g.FileUUID, ShouldEqual, "aft2345")


                g, err = GetGalleryItem(gItem2.Id, ac)
                So(err, ShouldBeNil)
                So(g, ShouldNotBeNil)
                So(g.Caption, ShouldEqual, "bum in your rum")

                bad, err := UpdateGalleryItem(&model.GalleryItem{}, ac)
                So(err, ShouldNotBeNil)
                So(bad, ShouldBeNil)


        })

}

func TestGetUserGalleryItems(t *testing.T) {

        Convey("Given that we can add items to a users gallery, we should be able to pull them all our by user", t, func() {

                g, err := GetGalleryItemsByUserId(test_user.Id, model.MODEL_JSON_CAPTION, false, ac)
                So(err, ShouldBeNil)
                So(g, ShouldNotBeNil)
                So(len(g), ShouldEqual, 3)
                So(g[0].Caption, ShouldEqual, "bum in your rum")
                So(g[1].Caption, ShouldEqual, "plum in my bum")
                So(g[2].Caption, ShouldEqual, "rum in your bum")

        })

        Convey("Given that we can add items to a users gallery, we should be able to pull them all our by user in reverse", t, func() {

                g, err := GetGalleryItemsByUserId(test_user.Id, model.MODEL_JSON_CAPTION, true, ac)
                So(err, ShouldBeNil)
                So(g, ShouldNotBeNil)
                So(len(g), ShouldEqual, 3)
                So(g[0].Caption, ShouldEqual, "rum in your bum")
                So(g[1].Caption, ShouldEqual, "plum in my bum")
                So(g[2].Caption, ShouldEqual, "bum in your rum")


        })

        Convey("Given that we can pull a users gallery items, we should be able to also get an galleries items.", t, func() {

                g, err := GetGalleryItems(g1, model.MODEL_JSON_NAME, false, ac)
                So(err, ShouldBeNil)
                So(g, ShouldNotBeNil)
                So(len(g), ShouldEqual, 3)
                So(g[0].Name, ShouldEqual, "Gallery Item 1")
                So(g[1].Name, ShouldEqual, "Gallery Item 2")
                So(g[2].Name, ShouldEqual, "Gallery Item 3")


        })

}

func TestRemoveItemFromGallery(t *testing.T) {

        Convey("We can pull our items, We should be able to remove an item from the gallery as well", t, func() {

                g, err := RemoveItemFromGallery(g1.Id, gItem3, ac)
                So(err, ShouldBeNil)
                So(g, ShouldNotBeNil)
                So(len(g.GalleryItems), ShouldEqual, 2)


        })

        Convey("If can remove a single item, We should be able to remove all of our items from our gallery", t, func() {

                g, err := RemoveItemFromGallery(g1.Id, gItem2, ac)
                So(err, ShouldBeNil)
                So(g, ShouldNotBeNil)
                So(len(g.GalleryItems), ShouldEqual, 1)

                g, err = RemoveItemFromGallery(g1.Id, gItem1, ac)
                So(err, ShouldBeNil)
                So(g, ShouldNotBeNil)
                So(len(g.GalleryItems), ShouldEqual, 0)

        })

}

func TestDeleteGallery(t *testing.T) {

        Convey("We should be able to delete a gallery once created", t, func() {

                err := DeleteGallery(g1, ac)
                So(err, ShouldBeNil)

                g, err := GetGallery(g1.Id, ac)
                So(err, ShouldNotBeNil)
                So(g, ShouldBeNil)

        })

}

func TestDeleteGalleryItem(t *testing.T) {

        Convey("We should be able to delete a gallery item once created", t, func() {

                err := DeleteGalleryItem(gItem1, ac)
                So(err, ShouldBeNil)

                g, err := GetInventoryItem(gItem1.Id, ac)
                So(err, ShouldNotBeNil)
                So(g, ShouldBeNil)

                err = DeleteGalleryItem(gItem2, ac)
                So(err, ShouldBeNil)

                err = DeleteGalleryItem(gItem3, ac)
                So(err, ShouldBeNil)

        })

}
