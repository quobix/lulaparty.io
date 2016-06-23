package model

import (
        "gopkg.in/mgo.v2/bson"
        "time"
)

const (
        MODEL_JSON_GALLERYID  = "gallery_id"
)

type Gallery struct {
        Id                      bson.ObjectId   `json:"id" bson:"_id,omitempty"`
        OwnerId                 bson.ObjectId   `json:"owner_id" bson:"owner_id,omitempty"`
        Created                 time.Time       `json:"created"`
        Updated                 time.Time       `json:"updated"`
        Caption                 string          `json:"caption"`
        Description             string          `json:"description"`
        Name                    string          `json:"name"`
        MasterGalleryItem       bson.ObjectId   `json:"master_id" bson:"master_id,omitempty"`
        Visible                 bool            `json:"visible`
        GalleryItems            []bson.ObjectId `json:"items" bson:"items,omitempty"`
}

type GalleryItem struct {
        Id                      bson.ObjectId   `json:"id" bson:"_id,omitempty"`
        OwnerId                 bson.ObjectId   `json:"owner_id" bson:"owner_id,omitempty"`
        GalleryId               bson.ObjectId   `json:"gallery_id" bson:"gallery_id,omitempty"`
        Created                 time.Time       `json:"created"`
        Updated                 time.Time       `json:"updated"`
        Caption                 string          `json:"caption"`
        Description             string          `json:"description"`
        FileUUID                string          `json:"file_uuid`
        Visible                 bool            `json:"visible`
        Name                    string          `json:"name"`
}

func (n *GalleryItem) SetCreated() (time.Time) {
        n.Created = getTime()
        return n.Created
}

func (n *GalleryItem) Update() (time.Time) {
        n.Updated = getTime()
        return n.Updated
}

func (n *GalleryItem) GetId() bson.ObjectId {
        return n.Id
}

func (n *GalleryItem) SetId(id bson.ObjectId) {
        n.Id = id
}

func (n *Gallery) SetCreated() (time.Time) {
        n.Created = getTime()
        return n.Created
}

func (n *Gallery) Update() (time.Time) {
        n.Updated = getTime()
        return n.Updated
}

func (n *Gallery) GetId() bson.ObjectId {
        return n.Id
}

func (n *Gallery) SetId(id bson.ObjectId) {
        n.Id = id
}

func (n *Gallery) ContainsItem(i bson.ObjectId) bool {
        for _, v := range n.GalleryItems {
                if (v == i) {
                        return true;
                }
        }
        return false;
}