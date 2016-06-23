package model

import (
        "gopkg.in/mgo.v2/bson"
        "time"
)

type Sale struct {
        Id                      bson.ObjectId   `json:"id" bson:"_id,omitempty"`
        OwnerId                 bson.ObjectId   `json:"owner_id" bson:"owner_id,omitempty"`
        PartyId                 bson.ObjectId   `json:"party_id" bson:"party_id,omitempty"`
        Created                 time.Time       `json:"created"`
        Updated                 time.Time       `json:"updated"`
        State                   string          `json:"state"`
        Completed               bool            `json:"completed"`
        Pending                 bool            `json:"pending"`
        Processing              bool            `json:"processing"`
        InventoryItem           bson.ObjectId   `json:"inventory_item" bson:"inventory_item,omitempty"`
        RequestedSize           string          `json:"size"`
        InventoryCount          uint8           `json:"inventory_count"`
}

func (n *Sale) SetCreated() (time.Time) {
        n.Created = getTime()
        return n.Created
}

func (n *Sale) Update() (time.Time) {
        n.Updated = getTime()
        return n.Updated
}

func (n *Sale) GetId() bson.ObjectId {
        return n.Id
}

func (n *Sale) SetId(id bson.ObjectId) {
        n.Id = id
}