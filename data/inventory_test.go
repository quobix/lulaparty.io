package data

import (
        "testing"
        . "github.com/smartystreets/goconvey/convey"
        "lulaparty.io/model"

        "gopkg.in/mgo.v2/bson"
        "github.com/shopspring/decimal"
)



func TestCreateInventory(t *testing.T) {

        if(ac == nil) {
                ac = CreateTestSession()
        }

        Convey("Given that we have an inventory, we should be able to persist it", t, func() {
                uuid = inv1.Id
                inv := &model.Inventory{
                        Id:inv1.Id,
                        OwnerId: test_user.Id,
                }


                ret_inv, err := CreateInventory(inv, ac)

                So(err, ShouldBeNil)
                So(ret_inv.Created, ShouldNotBeNil)

                //check that we can't duplicate record
                inv, err = CreateInventory(inv, ac)
                So(err, ShouldNotBeNil)



        })

}

func TestGetInventory(t *testing.T) {

        Convey("Given that we can persist the inventory, we should be able to retreive it", t, func() {

                i, err := GetInventory(inv1.Id, ac)
                So(err, ShouldBeNil)
                So(i, ShouldNotBeNil)
                So(i.Id, ShouldEqual, uuid)
                So(i.Id.Hex(), ShouldEqual, uuid.Hex())

                i, err = GetInventory(bson.NewObjectId(), ac)
                So(err, ShouldNotBeNil)

        })


}

func TestUpdateInventory(t *testing.T) {

        Convey("Given that we can retrive the record, lets update it", t, func() {

                inv1.Name = "Dave's Inventory"
                ret_inv, err := UpdateInventory(inv1, ac)

                So(err, ShouldBeNil)
                So(ret_inv, ShouldNotBeNil)
                So(ret_inv.Name, ShouldEqual, inv1.Name)

                // double check
                i, err := GetInventory(inv1.Id, ac)
                So(err, ShouldBeNil)
                So(i, ShouldNotBeNil)
                So(i.Name, ShouldEqual, inv1.Name)

                bad, err := GetInventory(bson.NewObjectId(), ac)
                So(bad, ShouldBeNil)
                So(err, ShouldNotBeNil)

                bad, err = UpdateInventory(&model.Inventory{}, ac)
                So(err, ShouldNotBeNil)
                So(bad, ShouldBeNil)


        })

}


func TestAddItemToInventory(t *testing.T) {

        //make sure our inventory items have an owner!
        invItem1.OwnerId = test_user.Id
        invItem2.OwnerId = test_user.Id
        invItem3.OwnerId = test_user.Id

        invItem1.Name="Item1"
        invItem2.Name="Item2"
        invItem3.Name="Item3"

        invItem1.InventoryId = inv1.Id;
        invItem2.InventoryId = inv1.Id;
        invItem3.InventoryId = inv1.Id;



        price1, err := decimal.NewFromString("45.22")
        if err != nil {
                panic(err)
        }
        price2, err := decimal.NewFromString("8.95")
        if err != nil {
                panic(err)
        }
        price3, err := decimal.NewFromString("72.44")
        if err != nil {
                panic(err)
        }
        invItem1.Price = model.Price{price1,model.MODEL_JSON_USD}
        invItem2.Price = model.Price{price2, model.MODEL_JSON_USD}
        invItem3.Price = model.Price{price3, model.MODEL_JSON_USD}


        Convey("Given that we have an inventory, we should be able to add an item without persisting it", t, func() {

                it, err := AddItemToInventory(inv1.Id, invItem1, ac)
                So(err, ShouldBeNil)
                So(it, ShouldNotBeNil)
                So(len(it.InventoryItems), ShouldEqual, 1)

        })

        Convey("We should be able to add another couple of items to the inventory", t, func() {

                it, err := AddItemToInventory(inv1.Id, invItem2, ac)
                So(err, ShouldBeNil)
                So(it, ShouldNotBeNil)
                So(len(it.InventoryItems), ShouldEqual, 2)

                it, err = AddItemToInventory(inv1.Id, invItem3, ac)
                So(err, ShouldBeNil)
                So(it, ShouldNotBeNil)
                So(len(it.InventoryItems), ShouldEqual, 3)

        })

        Convey("Trying to add a duplicate inventory item should result in an error", t, func() {

                it, err := AddItemToInventory(inv1.Id, invItem2, ac)
                So(err, ShouldNotBeNil)
                So(it, ShouldBeNil)

        })

}

func TestGetInventoryItem(t *testing.T) {

        Convey("Given that we can persist inventory items, we should be able to retrive them", t, func() {

                i, err := GetInventoryItem(invItem1.Id, ac)
                So(err, ShouldBeNil)
                So(i, ShouldNotBeNil)
                So(i.Id, ShouldEqual, invItem1.Id)
                So(i.Id.Hex(), ShouldEqual, invItem1.Id.Hex())

                i, err = GetInventoryItem(bson.NewObjectId(), ac)
                So(err, ShouldNotBeNil)

        })
}


func TestUpdateInventoryItem(t *testing.T) {

        Convey("Given that we can retrive the inventory item, we should be able to update it", t, func() {

                invItem1.CareInfo ="no wash wash"
                invItem2.CareInfo ="yes wash wash"
                ret_i1, err1 := UpdateInventoryItem(invItem1, ac)
                ret_i2, err2 := UpdateInventoryItem(invItem2, ac)

                So(err1, ShouldBeNil)
                So(err2, ShouldBeNil)
                So(ret_i1, ShouldNotBeNil)
                So(ret_i1.CareInfo, ShouldEqual, "no wash wash")
                So(ret_i2.CareInfo, ShouldEqual, "yes wash wash")

                // double check
                i, err := GetInventoryItem(invItem1.Id, ac)
                So(err, ShouldBeNil)
                So(i, ShouldNotBeNil)
                So(i.CareInfo, ShouldEqual, "no wash wash")

                i, err = GetInventoryItem(invItem2.Id, ac)
                So(err, ShouldBeNil)
                So(i, ShouldNotBeNil)
                So(i.CareInfo, ShouldEqual, "yes wash wash")

                bad, err := UpdateInventoryItem(&model.InventoryItem{}, ac)
                So(err, ShouldNotBeNil)
                So(bad, ShouldBeNil)


        })

}

func TestGetUserInventoryItems(t *testing.T) {

        Convey("Given that we can add items to a users inventory, we should be able to pull them all our by user", t, func() {

                it, err := GetUserInventoryItems(test_user, false, ac)
                So(err, ShouldBeNil)
                So(it, ShouldNotBeNil)
                So(len(it), ShouldEqual, 3)
                So(it[0].Name, ShouldEqual, "Item1")
                So(it[1].Name, ShouldEqual, "Item2")

        })

        Convey("Given that we can add items to a users inventory, we should be able to pull them all our by user in reverse", t, func() {

                it, err := GetUserInventoryItems(test_user, true, ac)
                So(err, ShouldBeNil)
                So(it, ShouldNotBeNil)
                So(len(it), ShouldEqual, 3)
                So(it[0].Name, ShouldEqual, "Item3")
                So(it[1].Name, ShouldEqual, "Item2")

        })

        Convey("Given that we can pull a users inventory items, we should be able to also get an inventories items.", t, func() {

                it, err := GetInventoryItems(inv1, model.MODEL_JSON_NAME, false, ac)
                So(err, ShouldBeNil)
                So(it, ShouldNotBeNil)
                So(len(it), ShouldEqual, 3)
                So(it[0].Name, ShouldEqual, "Item1")
                So(it[1].Name, ShouldEqual, "Item2")
                So(it[2].Name, ShouldEqual, "Item3")

                // check prices
                So(it[0].Price.Value.Equals(invItem1.Price.Value),ShouldBeTrue)
                So(it[1].Price.Value.Equals(invItem2.Price.Value),ShouldBeTrue)
                So(it[2].Price.Value.Equals(invItem3.Price.Value),ShouldBeTrue)

        })

}

func TestRemoveItemFromInventory(t *testing.T) {

        Convey("We should be able to remove an item from the inventory as well", t, func() {

                it, err := RemoveItemFromInventory(inv1.Id, invItem3, ac)
                So(err, ShouldBeNil)
                So(it, ShouldNotBeNil)
                So(len(it.InventoryItems), ShouldEqual, 2)


        })

        Convey("We should be able to remove all of our items from our inventory", t, func() {

                it, err := RemoveItemFromInventory(inv1.Id, invItem2, ac)
                So(err, ShouldBeNil)
                So(it, ShouldNotBeNil)
                So(len(it.InventoryItems), ShouldEqual, 1)

                it, err = RemoveItemFromInventory(inv1.Id, invItem1, ac)
                So(err, ShouldBeNil)
                So(it, ShouldNotBeNil)
                So(len(it.InventoryItems), ShouldEqual, 0)

        })

}

func TestDeleteInventory(t *testing.T) {

        Convey("We should be able to delete an inventory once created", t, func() {

                err := DeleteInventory(inv1, ac)
                So(err, ShouldBeNil)

                i, err := GetInventory(inv1.Id, ac)
                So(err, ShouldNotBeNil)
                So(i, ShouldBeNil)

                err = DeleteInventory(inv1, ac)
                So(err, ShouldNotBeNil)

        })

}

func TestDeleteInventoryItem(t *testing.T) {

        Convey("We should be able to delete an inventory item once created", t, func() {

                err := DeleteInventoryItem(invItem1, ac)
                So(err, ShouldBeNil)

                i, err := GetInventoryItem(invItem1.Id, ac)
                So(err, ShouldNotBeNil)
                So(i, ShouldBeNil)

                err = DeleteInventoryItem(invItem1, ac)
                So(err, ShouldNotBeNil)

        })

}