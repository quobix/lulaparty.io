package model

import (
        "gopkg.in/mgo.v2/bson"
        "time"
)

type CustomerProfile struct {
        Id                      bson.ObjectId   `json:"id" bson:"_id,omitempty"`
        OwnerId                 bson.ObjectId   `json:"owner_id" bson:"owner_id,omitempty"`
        HostParty               bool            `json:"host_party"`
        BecomeConsultant        string          `json:"last_name"`
        Created                 time.Time       `json:"created"`
        Updated                 time.Time       `json:"updated"`
}

type ProviderProfile struct {
        Id                      bson.ObjectId   `json:"id" bson:"_id,omitempty"`
        OwnerId                 bson.ObjectId   `json:"owner_id" bson:"owner_id,omitempty"`
        UseSquare               bool            `json:"use_square"`
        Created                 time.Time       `json:"created"`
        Updated                 time.Time       `json:"updated"`
}

type HostessProfile struct {
        Id                      bson.ObjectId   `json:"id" bson:"_id,omitempty"`
        OwnerId                 bson.ObjectId   `json:"owner_id" bson:"owner_id,omitempty"`
        Created                 time.Time       `json:"created"`
        Updated                 time.Time       `json:"updated"`
}


type FBProfile struct {
        Id                      bson.ObjectId   `json:"id" bson:"_id,omitempty"`
        OwnerId                 bson.ObjectId   `json:"owner_id" bson:"owner_id,omitempty"`
        Firstname               string          `json:"first_name"`
        Lastname                string          `json:"last_name"`
        Email                   string          `json:"email"`
        Created                 time.Time       `json:"created"`
        Updated                 time.Time       `json:"updated"`
        AccessToken             bson.ObjectId   `json:"token_id" bson:"token_id,omitempty"`
}


type AccessToken struct {
        Id                      bson.ObjectId   `json:"id" bson:"_id,omitempty"`
        Token                   string          `json:"token"`
        ExpiryInSeconds         int             `json:"expires_seconds"`
        Expires                 time.Time       `json:"expires"`
        Created                 time.Time       `json:"created"`
        Updated                 time.Time       `json:"updated"`
}


func (n *FBProfile) SetCreated() (time.Time) {
        n.Created = getTime()
        return n.Created
}

func (n *FBProfile) Update() (time.Time) {
        n.Updated =  getTime()
        return n.Updated
}

func (n *FBProfile) GetId() bson.ObjectId {
        return n.Id
}

func (n *FBProfile) SetId(id bson.ObjectId) {
        n.Id = id
}

func (n *AccessToken) SetCreated() (time.Time) {
        n.Created = getTime()
        return n.Created
}

func (n *AccessToken) Update() (time.Time) {
        n.Updated =  getTime()
        return n.Updated
}

func (n *AccessToken) GetId() bson.ObjectId {
        return n.Id
}

func (n *AccessToken) SetId(id bson.ObjectId) {
        n.Id = id
}

func (n *HostessProfile) SetCreated() (time.Time) {
        n.Created = getTime()
        return n.Created
}

func (n *HostessProfile) Update() (time.Time) {
        n.Updated = getTime()
        return n.Updated
}

func (n *HostessProfile) GetId() bson.ObjectId {
        return n.Id
}

func (n *HostessProfile) SetId(id bson.ObjectId) {
        n.Id = id
}

func (n *CustomerProfile) SetCreated() (time.Time) {
        n.Created = getTime()
        return n.Created
}

func (n *CustomerProfile) Update() (time.Time) {
        n.Updated = getTime()
        return n.Updated
}

func (n *CustomerProfile) GetId() bson.ObjectId {
        return n.Id
}

func (n *CustomerProfile) SetId(id bson.ObjectId) {
        n.Id = id
}

func (n *ProviderProfile) SetCreated() (time.Time){
        n.Created = getTime()
        return n.Created
}

func (n *ProviderProfile) Update() (time.Time) {
        n.Updated = getTime()
        return n.Updated
}

func (n *ProviderProfile) GetId() bson.ObjectId {
        return n.Id
}

func (n *ProviderProfile) SetId(id bson.ObjectId) {
        n.Id = id
}
