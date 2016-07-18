package model

import (
        "gopkg.in/mgo.v2/bson"
        "time"
)

type User struct {
        Id                      bson.ObjectId   `json:"id" bson:"_id,omitempty"`
        FBProfile               bson.ObjectId   `json:"fb_pid" bson:"fb_pid,omitempty"`
        CustomerProfile         bson.ObjectId   `json:"cust_pid" bson:"cust_pid,omitempty"`
        HostessProfile          bson.ObjectId   `json:"host_pid" bson:"host_pid,omitempty"`
        ProviderProfile         bson.ObjectId   `json:"provider_pid" bson:"profile_id,omitempty"`
        Address                 bson.ObjectId   `json:"address_id" bson:"address_id,omitempty"`
        Created                 time.Time       `json:"created"`
        Updated                 time.Time       `json:"updated"`
        Cell                    string          `json:"cell"`
        Email                   string          `json:"email"`
        ShipPurchases           bool            `json:"ship_purchases"`
        LastAuth                time.Time       `json:"last_auth"`
        TrialExpired            bool            `json:"trial_expired"`
        Sales                   []bson.ObjectId `json:"sales" bson:"sales,omitempty"`
        FBAuthToken             string          `json:"fbauth_token"`
}

type Address struct {
        Id              bson.ObjectId   `json:"id" bson:"_id,omitempty"`
        OwnerId         bson.ObjectId   `json:"owner_id" bson:"owner_id,omitempty"`
        Street1         string          `json:"street1"`
        Street2         string          `json:"street2"`
        City            string          `json:"city"`
        State           string          `json:"state"`
        Zip             string          `json:"zip"`
        Created         time.Time       `json:"created"`
        Updated         time.Time       `json:"updated"`
}


func (n *Address) SetCreated() (time.Time) {
        n.Created = getTime()
        return n.Created
}

func (n *Address) Update() (time.Time) {
        n.Updated = getTime()
        return n.Updated
}

func (n *Address) GetId() bson.ObjectId {
        return n.Id
}

func (n *Address) SetId(id bson.ObjectId) {
        n.Id = id
}

func (n *User) SetCreated() (time.Time){
        n.Created = getTime()
        return n.Created
}

func (n *User) Update() (time.Time){
        n.Updated = getTime()
        return n.Updated
}

func (n *User) GetId() bson.ObjectId {
        return n.Id
}

func (n *User) SetId(id bson.ObjectId) {
        n.Id = id
}
