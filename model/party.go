package model

import (
        "gopkg.in/mgo.v2/bson"
        "time"
        "github.com/quobix/lulaparty.io/util"
)

type Party struct {
        Id                      bson.ObjectId   `json:"id" bson:"_id,omitempty"`
        OwnerId                 bson.ObjectId   `json:"owner_id" bson:"owner_id,omitempty"`
        PartyInventoryId        bson.ObjectId   `json:"party_inv_id" bson:"inv_id,omitempty"`
        Created                 time.Time       `json:"created"`
        Updated                 time.Time       `json:"updated"`
        Name                    string          `json:"name"`
        Description             string          `json:"description"`
        GalleryId               bson.ObjectId   `json:"gallery_id" bson:"gallery_id,omitempty"`
        Starts                  time.Time       `json:"starts"`
        Ends                    time.Time       `json:"ends"`
        Enabled                 bool            `json:"enabled"`
        ShortCode               string          `json:"shortcode"`
        Providers               []bson.ObjectId `json:"providers" bson:"providers,omitempty"`
        Hostesses               []bson.ObjectId `json:"hostesses" bson:"hostesses,omitempty"`
}

type PartyInventory struct {
        Id                      bson.ObjectId   `json:"id" bson:"_id,omitempty"`
        OwnerId                 bson.ObjectId   `json:"owner_id" bson:"owner_id,omitempty"`
        PartyId                 bson.ObjectId   `json:"party_id" bson:"party_id,omitempty"`
        Created                 time.Time       `json:"created"`
        Updated                 time.Time       `json:"updated"`
        GalleryId               bson.ObjectId   `json:"gallery_id" bson:"gallery_id,omitempty"`
        MultiplePurchases       bool            `json:"mutiple_purchases"`
        InventoryItems          []bson.ObjectId `json:"inventory_items" bson:"inventory_items,omitempty"`
        Name                    string          `json:"name"`
        Description             string          `json:"description"`
}

type PartyInvite struct {
        Id                      bson.ObjectId   `json:"id" bson:"_id,omitempty"`
        OwnerId                 bson.ObjectId   `json:"owner_id" bson:"owner_id,omitempty"`
        PartyId                 bson.ObjectId   `json:"party_id" bson:"party_id,omitempty"`
        Created                 time.Time       `json:"created"`
        Updated                 time.Time       `json:"updated"`
        Expires                 time.Time       `json:"expires"`
        Accepted                bool            `json:"accepted"`
        ShortCode               string          `json:"shortcode"`
        Viewed                  bool            `json:"viewed"`
        Sender                  bson.ObjectId   `json:"sender" bson:"sender,omitempty"`
        Recipient               bson.ObjectId   `json:"recipient" bson:"recipient,omitempty"`
}

type Reward struct {
        Id                      bson.ObjectId   `json:"id" bson:"_id,omitempty"`
        OwnerId                 bson.ObjectId   `json:"owner_id" bson:"owner_id,omitempty"`
        PartyInviteId           bson.ObjectId   `json:"party_invite_id" bson:"party_invite_id,omitempty"`
        Created                 time.Time       `json:"created"`
        Updated                 time.Time       `json:"updated"`
        CreditReward            bool            `json:"credit_reward"`
        ItemReward              bool            `json:"item_reward"`
        CreditValue             Price           `json:"credit_value"`
        ValueCeiling            uint8           `json:"value_ceiling"`
        ValueDelta              uint8           `json:"value_delta"`
}

func (n *Reward) SetCreated() (time.Time) {
        n.Created = getTime()
        return n.Created
}

func (n *Reward) Update() (time.Time) {
        n.Updated = getTime()
        return n.Updated
}

func (n *Reward) GetId() bson.ObjectId {
        return n.Id
}

func (n *Reward) SetId(id bson.ObjectId) {
        n.Id = id
}

func (n *PartyInvite) SetCreated() (time.Time) {
        n.Created = getTime()
        return n.Created
}

func (n *PartyInvite) Update() (time.Time) {
        n.Updated = getTime()
        return n.Updated
}

func (n *PartyInvite) GetId() bson.ObjectId {
        return n.Id
}

func (n *PartyInvite) SetId(id bson.ObjectId) {
        n.Id = id
}

func (n *PartyInventory) SetCreated() (time.Time) {
        n.Created = getTime()
        return n.Created
}

func (n *PartyInventory) Update() (time.Time) {
        n.Updated = getTime()
        return n.Updated
}

func (n *PartyInventory) GetId() bson.ObjectId {
        return n.Id
}

func (n *PartyInventory) SetId(id bson.ObjectId) {
        n.Id = id
}

func (n *Party) SetCreated() (time.Time) {
        n.Created = getTime()
        return n.Created
}

func (n *Party) Update() (time.Time) {
        n.Updated = getTime()
        return n.Updated
}

func (n *Party) GetId() bson.ObjectId {
        return n.Id
}

func (n *Party) SetId(id bson.ObjectId) {
        n.Id = id
}

func (n *Party) IsOpen() bool {
        t := time.Now()
        if(!n.Enabled) { return false }
        if(n.Starts.Before(t) && t.Before(n.Ends)) { return true }
        return false
}





func (n *Party) StartsIn() string {
        if(n.IsOpen()) {
                return PARTY_STARTED
        }
        return util.TimeDiffHelper(n.Starts, PARTY_STARTED,
                        func(t time.Time) bool {
                                return t.After(time.Now().UTC())
                        })
}

func (n *Party) EndsIn() string {
        if(n.IsOpen()) {
                return util.TimeDiffHelper(n.Ends, PARTY_ENDED,
                        func(t time.Time) bool {
                                return time.Now().UTC().Before(t)
                        })
        } else {
                return PARTY_CLOSED
        }
}

func (n *PartyInvite) HasExpired() bool {
        t := time.Now().UTC()
        if(n.Accepted) { return true }
        if(t.Before(n.Expires)) { return false }
        return true
}

func (n *PartyInvite) ExpiresIn() string {
        if(!n.Accepted) {
                return util.TimeDiffHelper(n.Expires, PARTY_INVITE_EXPIRED,
                        func(t time.Time) bool {
                                return time.Now().UTC().Before(t)
                        })
        } else {
                return PARTY_INVITE_ACCEPTED
        }
}


