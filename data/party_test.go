package data

import (
        "testing"
        . "github.com/smartystreets/goconvey/convey"
        "lulaparty.io/model"
        "gopkg.in/mgo.v2/bson"
        "time"
        "lulaparty.io/util"
        "github.com/shopspring/decimal"
)

func TestCreateParty(t *testing.T) {

        if(ac == nil) {
                ac = CreateTestSession()
        }

        // forced db cleanup
        //ac.DBSession.DB(ac.DBName).DropDatabase() // cleanup.

        Convey("Given that we have a party, we should be able to persist it", t, func() {

                p1 = &model.Party{
                        Id:bson.NewObjectId(),
                        OwnerId: test_user.Id,
                        Name: "Super Party 1" }

                ret_p, err := CreateParty(p1, ac)

                So(err, ShouldBeNil)
                So(ret_p.Created, ShouldNotBeNil)

                //check that we can't duplicate record
                ret_p, err = CreateParty(p1, ac)
                So(err, ShouldNotBeNil)

        })
}

func TestGetParty(t *testing.T) {

        Convey("Given that we can persist a party, we should be able to retrieve it", t, func() {

                p1, err := GetParty(p1.Id, ac)

                So(err, ShouldBeNil)
                So(p1.Created, ShouldNotBeNil)
                So(p1.Name, ShouldEqual, "Super Party 1")

        })
}

func TestUpdateParty(t *testing.T) {

        Convey("Given that we can retrieve a party, we should be able to update it", t, func() {
                p1.Description = "The best party you can have"
                ret, err := UpdateParty(p1, ac)
                So(err, ShouldBeNil)
                So(ret, ShouldNotBeNil)

                ret, err = GetParty(p1.Id, ac) // check with a fresh fetch
                So(ret.Description, ShouldEqual, "The best party you can have")

                ret, err = UpdateParty(&model.Party{}, ac)
                So(err, ShouldNotBeNil)
                So(ret, ShouldBeNil)

        })
}

func TestDeleteParty(t *testing.T) {

        Convey("Given that we can update a party, we should be able to remove it", t, func() {

                err := DeleteParty(p1, ac)
                So(err, ShouldBeNil)

                rp1, err := GetParty(p1.Id, ac) // check we can't find it.

                So(err, ShouldNotBeNil)
                So(rp1, ShouldBeNil)

                err = DeleteParty( &model.Party{}, ac)
                So(err, ShouldNotBeNil)
        })
}

func TestAddProviderToParty(t *testing.T) {
        Convey("Given that we have a party we should be able to add an a provider with out having to persist the provider", t, func() {

                p2 = &model.Party{
                        Id:bson.NewObjectId(),
                        OwnerId: test_user.Id,
                        Name: "Super Party 2" }

                _, err := CreateParty(p2, ac)


                p, err := AddProviderToParty(p2.Id, prov1, ac)
                So(err, ShouldBeNil)
                So(p, ShouldNotBeNil)
                So(len(p.Providers), ShouldEqual, 1)

        })

        Convey("We should be able to add another couple of providers to the party", t, func() {

                p, err := AddProviderToParty(p2.Id, prov2, ac)
                So(err, ShouldBeNil)
                So(p, ShouldNotBeNil)
                So(len(p.Providers), ShouldEqual, 2)

                p, err = AddProviderToParty(p2.Id, prov3, ac)
                So(err, ShouldBeNil)
                So(p, ShouldNotBeNil)
                So(len(p.Providers), ShouldEqual, 3)

        })

        Convey("When Trying to add a duplicate provider to a party we should result in an error", t, func() {

                p, err := AddProviderToParty(p2.Id, prov3, ac)
                So(err, ShouldNotBeNil)
                So(p, ShouldBeNil)

        })
}

func TestRemoveProviderFromParty(t *testing.T) {
        Convey("We should be able to remove providers from a party", t, func() {

                p, err := RemoveProviderFromParty(p2.Id, prov1, ac)
                So(err, ShouldBeNil)
                So(p,ShouldNotBeNil)
                So(len(p.Providers), ShouldEqual, 2)

        })

        Convey("We should be able to remove all of the party providers", t, func() {

                p, err := RemoveProviderFromParty(p2.Id, prov2, ac)
                So(err, ShouldBeNil)
                So(p,ShouldNotBeNil)
                So(len(p.Providers), ShouldEqual, 1)

                p, err = RemoveProviderFromParty(p2.Id, prov3, ac)
                So(err, ShouldBeNil)
                So(p,ShouldNotBeNil)
                So(len(p.Providers), ShouldEqual, 0)

        })
}

func TestAddHostessToParty(t *testing.T) {
        Convey("Given that we have a party we should be able to add a hostess with out having to persist the provider", t, func() {

                p, err := AddHostessToParty(p2.Id, hostess1, ac)
                So(err, ShouldBeNil)
                So(p, ShouldNotBeNil)
                So(len(p.Hostesses), ShouldEqual, 1)

        })

        Convey("We should be able to add another couple of hostesses to the party", t, func() {

                p, err := AddHostessToParty(p2.Id, hostess2, ac)
                So(err, ShouldBeNil)
                So(p, ShouldNotBeNil)
                So(len(p.Hostesses), ShouldEqual, 2)

                p, err = AddHostessToParty(p2.Id, hostess3, ac)
                So(err, ShouldBeNil)
                So(p, ShouldNotBeNil)
                So(len(p.Hostesses), ShouldEqual, 3)

        })

        Convey("When Trying to add a duplicate hostess to a party, we should result in an error", t, func() {

                p, err := AddHostessToParty(p2.Id, hostess3, ac)
                So(err, ShouldNotBeNil)
                So(p, ShouldBeNil)

        })
}

func TestRemoveHostessFromParty(t *testing.T) {
        Convey("We should be able to remove hostesses from a party", t, func() {

                p, err := RemoveHostessFromParty(p2.Id, hostess2, ac)
                So(err, ShouldBeNil)
                So(p,ShouldNotBeNil)
                So(len(p.Hostesses), ShouldEqual, 2)

        })

        Convey("We should be able to remove all of the party hostesses", t, func() {

                p, err := RemoveHostessFromParty(p2.Id, hostess1, ac)
                So(err, ShouldBeNil)
                So(p,ShouldNotBeNil)
                So(len(p.Hostesses), ShouldEqual, 1)

                p, err = RemoveHostessFromParty(p2.Id, hostess3, ac)
                So(err, ShouldBeNil)
                So(p,ShouldNotBeNil)
                So(len(p.Hostesses), ShouldEqual, 0)

                p, err = RemoveHostessFromParty(p2.Id, hostess3, ac)
                So(err, ShouldNotBeNil)
                So(p,ShouldBeNil)
        })
}

func TestCreatePartyInventory(t *testing.T) {

        Convey("Given that we have an party inventory, we should be able to persist it", t, func() {
                pi1 = &model.PartyInventory{
                        Id:bson.NewObjectId(),
                        OwnerId: test_user.Id }

                ret_inv, err := CreatePartyInventory(pi1, ac)

                So(err, ShouldBeNil)
                So(ret_inv.Created, ShouldNotBeNil)

                //check that we can't duplicate record
                _, err = CreatePartyInventory(pi1, ac)
                So(err, ShouldNotBeNil)
        })
}

func TestGetPartyInventory(t *testing.T) {

        Convey("Given that we can persist the party inventory, we should be able to retreive it", t, func() {

                i, err := GetPartyInventory(pi1.Id, ac)
                So(err, ShouldBeNil)
                So(i, ShouldNotBeNil)
                So(i.Id, ShouldEqual, pi1.Id)
                So(i.Id.Hex(), ShouldEqual, pi1.Id.Hex())

                i, err = GetPartyInventory(bson.NewObjectId(), ac)
                So(err, ShouldNotBeNil)

        })
}

func TestUpdatePartyInventory(t *testing.T) {

        Convey("Given that we can retrive the record, lets update it", t, func() {

                pi1.Name = "Dave's Party Inventory"
                ret_inv, err := UpdatePartyInventory(pi1, ac)

                So(err, ShouldBeNil)
                So(ret_inv, ShouldNotBeNil)
                So(ret_inv.Name, ShouldEqual, pi1.Name)

                // double check
                i, err := GetPartyInventory(pi1.Id, ac)
                So(err, ShouldBeNil)
                So(i, ShouldNotBeNil)
                So(i.Name, ShouldEqual, pi1.Name)

                bad, err := GetPartyInventory(bson.NewObjectId(), ac)
                So(bad, ShouldBeNil)
                So(err, ShouldNotBeNil)

                bad, err = UpdatePartyInventory(&model.PartyInventory{}, ac)
                So(err, ShouldNotBeNil)
                So(bad, ShouldBeNil)
        })
}



func TestAddItemToPartyInventory(t *testing.T) {

        //make sure our inventory items have an owner!
        invItem1.OwnerId = test_user.Id
        invItem2.OwnerId = test_user.Id
        invItem3.OwnerId = test_user.Id

        invItem1.Name="PItem1"
        invItem2.Name="PItem2"
        invItem3.Name="PItem3"

        invItem1.InventoryId = pi1.Id;
        invItem2.InventoryId = pi1.Id;
        invItem3.InventoryId = pi1.Id;



        Convey("Given that we have an party inventory, we should be able to add an item without persisting it", t, func() {

                it, err := AddItemToPartyInventory(pi1.Id, invItem1, ac)
                So(err, ShouldBeNil)
                So(it, ShouldNotBeNil)
                So(len(it.InventoryItems), ShouldEqual, 1)

        })

        Convey("We should be able to add another couple of items to the inventory", t, func() {

                it, err := AddItemToPartyInventory(pi1.Id, invItem2, ac)
                So(err, ShouldBeNil)
                So(it, ShouldNotBeNil)
                So(len(it.InventoryItems), ShouldEqual, 2)

                it, err = AddItemToPartyInventory(pi1.Id, invItem3, ac)
                So(err, ShouldBeNil)
                So(it, ShouldNotBeNil)
                So(len(it.InventoryItems), ShouldEqual, 3)

        })

        Convey("Trying to add a duplicate inventory item should result in an error", t, func() {

                it, err := AddItemToPartyInventory(pi1.Id, invItem2, ac)
                So(err, ShouldNotBeNil)
                So(it, ShouldBeNil)

        })
}

func TestRemoveItemFromPartyInventory(t *testing.T) {

        Convey("We should be able to remove an item from the party inventory as well", t, func() {

                it, err := RemoveItemFromPartyInventory(pi1.Id, invItem3, ac)
                So(err, ShouldBeNil)
                So(it, ShouldNotBeNil)
                So(len(it.InventoryItems), ShouldEqual, 2)


        })

        Convey("We should be able to remove all of our items from our party inventory", t, func() {

                it, err := RemoveItemFromPartyInventory(pi1.Id, invItem2, ac)
                So(err, ShouldBeNil)
                So(it, ShouldNotBeNil)
                So(len(it.InventoryItems), ShouldEqual, 1)

                it, err = RemoveItemFromPartyInventory(pi1.Id, invItem1, ac)
                So(err, ShouldBeNil)
                So(it, ShouldNotBeNil)
                So(len(it.InventoryItems), ShouldEqual, 0)

        })
}

func TestDeletePartyInventory(t *testing.T) {

        Convey("We should be able to delete an party inventory once created", t, func() {

                err := DeletePartyInventory(pi1, ac)
                So(err, ShouldBeNil)

                i, err := GetPartyInventory(pi1.Id, ac)
                So(err, ShouldNotBeNil)
                So(i, ShouldBeNil)

                err = DeletePartyInventory(pi1, ac)
                So(err, ShouldNotBeNil)

        })

}
func TestCreatePartyInvite(t *testing.T) {
        Convey("Given that we have a party invite, we should be able to persist it", t, func() {
                sc1 := util.RandStringRunes(5)
                p3 = &model.Party{
                        Id:bson.NewObjectId(),
                        OwnerId: test_user.Id,
                        Name: "Super Party 3" }

                _, err := CreateParty(p3, ac)

                ti := time.Now().UTC()

                pinv1 = &model.PartyInvite {
                        Id:bson.NewObjectId(),
                        OwnerId: test_user.Id,
                        Accepted: false,
                        PartyId: p3.Id,
                        Expires: time.Date(ti.Year(), ti.Month(), ti.Day(),
                                ti.Hour()+5, ti.Minute(), ti.Second(), ti.Nanosecond(), time.UTC),
                        ShortCode: sc1 }

                ret_pinv, err := CreatePartyInvite(pinv1, ac)

                So(err, ShouldBeNil)
                So(ret_pinv.Created, ShouldNotBeNil)
                So(ret_pinv.ExpiresIn(), ShouldEqual, "5 Hours")
                So(ret_pinv.HasExpired(), ShouldBeFalse)

                pinv2 = &model.PartyInvite {
                        Id:bson.NewObjectId(),
                        OwnerId: test_user.Id,
                        Accepted: false,
                        PartyId: p3.Id,
                        Expires: time.Date(ti.Year(), ti.Month(), ti.Day(),
                                ti.Hour()+9, ti.Minute()+2, ti.Second(), ti.Nanosecond(), time.UTC),
                        ShortCode: sc1 }

                So(CheckInviteShortCodeIsUnique(sc1, ac), ShouldBeFalse)

                ret_pinv, err = CreatePartyInvite(pinv2, ac)

                So(err, ShouldNotBeNil)

                pinv2 = &model.PartyInvite {
                        Id:bson.NewObjectId(),
                        OwnerId: test_user.Id,
                        Accepted: false,
                        PartyId: p3.Id,
                        Expires: time.Date(ti.Year(), ti.Month(), ti.Day(),
                                ti.Hour()+12, ti.Minute()+2, ti.Second(), ti.Nanosecond(), time.UTC),
                        ShortCode: util.RandStringRunes(5) }

                ret_pinv, err = CreatePartyInvite(pinv2, ac)

                So(err, ShouldBeNil)
                So(ret_pinv.Created, ShouldNotBeNil)
                So(ret_pinv.ExpiresIn(), ShouldEqual, "12 Hours, 2 Mins")
                So(ret_pinv.HasExpired(), ShouldBeFalse)

        })
}

func TestGetPartyInvite(t *testing.T) {

        Convey("Given that we can persist a party invite, we should be able to retrieve it", t, func() {

                _pinv, err := GetPartyInvite(pinv1.Id, ac)

                So(err, ShouldBeNil)
                So(_pinv.Created, ShouldNotBeNil)


        })
}

func TestUpdatePartyInvite(t *testing.T) {

        Convey("Given that we can retrieve a party invite, we should be able to update it", t, func() {

                _pinv, err := GetPartyInvite(pinv1.Id, ac)
                _pinv.Accepted = true

                ret, err := UpdatePartyInvite(_pinv, ac)
                So(err, ShouldBeNil)
                So(ret, ShouldNotBeNil)

                ret, err = GetPartyInvite(_pinv.Id, ac) // check with a fresh fetch
                So(ret.Accepted, ShouldBeTrue)
                So(ret.HasExpired(), ShouldBeTrue)
                So(ret.ExpiresIn(), ShouldEqual, model.PARTY_INVITE_ACCEPTED)

                ret, err = UpdatePartyInvite(&model.PartyInvite{}, ac)
                So(err, ShouldNotBeNil)
                So(ret, ShouldBeNil)

        })
}

func TestGetPartyInviteByShortCode(t *testing.T) {

        Convey("Given that we can retrieve a party invite, we should be able to find by code", t, func() {

                _pinv, err := GetPartyInviteByShortCode("abc123", ac)
                So(err, ShouldNotBeNil)
                So(_pinv, ShouldBeNil)

                _pinv, err = GetPartyInviteByShortCode(pinv1.ShortCode, ac)
                So(err, ShouldBeNil)
                So(_pinv, ShouldNotBeNil)
                So(_pinv.ShortCode, ShouldEqual, pinv1.ShortCode)

                _pinv, err = GetPartyInviteByShortCode(pinv2.ShortCode, ac)
                So(err, ShouldBeNil)
                So(_pinv, ShouldNotBeNil)
                So(_pinv.ShortCode, ShouldEqual, pinv2.ShortCode)

        })
}



func TestCreateReward(t *testing.T) {

        Convey("Given that we have a reward, we should be able to persist it", t, func() {

                rew1 = &model.Reward {
                        PartyInviteId: pinv1.Id }

                price1, _ := decimal.NewFromString("45.22")
                rew1.CreditValue = model.Price{price1,model.MODEL_JSON_USD}

                _, err := CreateReward(rew1, ac)
                So(err, ShouldBeNil)
                So(rew1.Id, ShouldNotBeNil)

                _, err = CreateReward(rew1, ac)
                So(err, ShouldNotBeNil)


        })
}

func TestGetReward(t *testing.T) {

        Convey("Given that we can persist a reward, we should be able to retrieve it", t, func() {

                r, err := GetReward(rew1.Id, ac)

                So(err, ShouldBeNil)
                So(r.Created, ShouldNotBeNil)
                So(r.CreditValue, ShouldNotBeNil)

                So(r.CreditValue.Value.String(), ShouldEqual, "45.22")


        })
}

func TestUpdateReward(t *testing.T) {

        Convey("Given that we can retrieve a reward, we should be able to update it", t, func() {

                price1, _ := decimal.NewFromString("233443.23")
                rew1.CreditValue = model.Price{price1,model.MODEL_JSON_USD}


                ret, err := UpdateReward(rew1, ac)
                So(err, ShouldBeNil)
                So(ret, ShouldNotBeNil)
                So(ret.CreditValue.Value.String(), ShouldEqual, "233443.23")

                ret, err = GetReward(rew1.Id, ac) // check with a fresh fetch
                So(err, ShouldBeNil)
                So(ret, ShouldNotBeNil)
                So(ret.CreditValue.Value.String(), ShouldEqual, "233443.23")

                ret, err = UpdateReward(&model.Reward{}, ac)
                So(err, ShouldNotBeNil)
                So(ret, ShouldBeNil)

        })
}

func TestDeleteReward(t *testing.T) {

        Convey("Given that we can update a reward, we should be able to remove it", t, func() {

                err := DeleteReward(rew1, ac)
                So(err, ShouldBeNil)

                rp1, err := GetReward(rew1.Id, ac) // check we can't find it.

                So(err, ShouldNotBeNil)
                So(rp1, ShouldBeNil)

                err = DeleteReward( &model.Reward{}, ac)
                So(err, ShouldNotBeNil)
        })
}

func TestDeletePartyInvite(t *testing.T) {

        Convey("Given that we can update a party invite, we should be able to remove it", t, func() {

                err := DeletePartyInvite(pinv1, ac)
                So(err, ShouldBeNil)

                err = DeletePartyInvite(pinv2, ac)
                So(err, ShouldBeNil)

                rp1, err := GetPartyInvite(pinv1.Id, ac) // check we can't find it.

                So(err, ShouldNotBeNil)
                So(rp1, ShouldBeNil)

                rp2, err := GetPartyInvite(pinv2.Id, ac) // check we can't find it.

                So(err, ShouldNotBeNil)
                So(rp2, ShouldBeNil)


                err = DeletePartyInvite( &model.PartyInvite{}, ac)
                So(err, ShouldNotBeNil)
        })
}