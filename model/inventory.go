package model

import (
        "gopkg.in/mgo.v2/bson"
        "time"
        "github.com/shopspring/decimal"

)

const (
        SIZE_SMALL              = "small"
        SIZE_MEDIUM             = "medium"
        SIZE_LARGE              = "large"
        SIZE_XL                 = "xl"
        SIZE_XXL                = "xxl"
        SIZE_XXXL               = "xxxl"
        MODEL_JSON_INVENTORYID  = "inventory_id"
        MODEL_JSON_PRICE        = "price"
        MODEL_JSON_USD          = "usd"
)

type Inventory struct {
        Id                      bson.ObjectId   `json:"id" bson:"_id,omitempty"`
        OwnerId                 bson.ObjectId   `json:"owner_id" bson:"owner_id,omitempty"`
        Created                 time.Time       `json:"created"`
        Updated                 time.Time       `json:"updated"`
        InventoryItems          []bson.ObjectId `json:"items" bson:"items,omitempty"`
        Name                    string          `json:"name"`
        AdminNotes              string          `json:"admin_notes"`
}

type InventoryItem struct {
        Id                      bson.ObjectId   `json:"id" bson:"_id,omitempty"`
        OwnerId                 bson.ObjectId   `json:"owner_id" bson:"owner_id,omitempty"`
        InventoryId             bson.ObjectId   `json:"inventory_id" bson:"inventory_id,omitempty"`
        StyleId                 bson.ObjectId   `json:"style_id" bson:"style_id,omitempty"`
        Created                 time.Time       `json:"created"`
        Updated                 time.Time       `json:"updated"`
        Name                    string          `json:"name"`
        Description             string          `json:"description"`
        CareInfo                string          `json:"care_info"`
        AdminNotes              string          `json:"admin_notes"`
        Price                   Price           `json:"price"`
        GalleryId               bson.ObjectId   `json:"gellery_id" bson:"gallery_id,omitempty"`
}

type InventoryItemSizes struct {
        Id                      bson.ObjectId   `json:"id" bson:"_id,omitempty"`
        OwnerId                 bson.ObjectId   `json:"owner_id" bson:"owner_id,omitempty"`
        InventoryId             bson.ObjectId   `json:"inventory_id" bson:"inventory_id,omitempty"`
        InventoryItemId         bson.ObjectId   `json:"inventory_item_id" bson:"inventory_id,omitempty"`
        Created                 time.Time       `json:"created"`
        Updated                 time.Time       `json:"updated"`
        SmallCount              uint8           `json:"small"`
        MediumCount             uint8           `json:"medium"`
        LargeCount              uint8           `json:"large"`
        XLCount                 uint8           `json:"xl"`
        XXLCount                uint8           `json:"xxl"`
        XXXLCount               uint8           `json:"xxxl"`
}

type Style struct {
        Id                      bson.ObjectId   `json:"id" bson:"_id,omitempty"`
        OwnerId                 bson.ObjectId   `json:"owner_id" bson:"owner_id,omitempty"`
        Created                 time.Time       `json:"created"`
        Updated                 time.Time       `json:"updated"`
        Name                    string          `json:"name"`
        Description             string          `json:"description"`
        GalleryItemId           bson.ObjectId   `json:"gallery_item_id" bson:"gallery_item_id,omitempty"`
}


// taken from http://stackoverflow.com/questions/30891301/handling-custom-bson-marshaling-golang-mgo
type Price struct {
        Value                   decimal.Decimal
        CurrencyCode            string
}

// GetBSON implements bson.Getter.
func (c Price) GetBSON() (interface{}, error) {
        f, _ := c.Value.Float64()
       return struct {
                Value        float64 `json:"value" bson:"value"`
                CurrencyCode string  `json:"currencyCode" bson:"currencyCode"`
        }{
                Value:        f,
                CurrencyCode: c.CurrencyCode,
        }, nil
}

// SetBSON implements bson.Setter.
func (c *Price) SetBSON(raw bson.Raw) error {
        decoded := new(struct {
                Value        float64 `json:"value" bson:"value"`
                CurrencyCode string  `json:"currencyCode" bson:"currencyCode"`
        })

        bsonErr := raw.Unmarshal(decoded)

        if bsonErr == nil {
                c.Value = decimal.NewFromFloat(decoded.Value)
                c.CurrencyCode = decoded.CurrencyCode
                return nil
        } else {
                return bsonErr
        }
}

func (n *Style) SetCreated() (time.Time) {
        n.Created = getTime()
        return n.Created
}

func (n *Style) Update() (time.Time) {
        n.Updated = getTime()
        return n.Updated
}

func (n *Style) GetId() bson.ObjectId {
        return n.Id
}

func (n *Style) SetId(id bson.ObjectId) {
        n.Id = id
}

func (n *InventoryItemSizes) SetCreated() (time.Time) {
        n.Created = getTime()
        return n.Created
}

func (n *InventoryItemSizes) Update() (time.Time) {
        n.Updated = getTime()
        return n.Updated
}

func (n *InventoryItemSizes) GetId() bson.ObjectId {
        return n.Id
}

func (n *InventoryItemSizes) SetId(id bson.ObjectId) {
        n.Id = id
}

func (n *InventoryItem) SetCreated() (time.Time) {
        n.Created = getTime()
        return n.Created
}

func (n *InventoryItem) Update() (time.Time) {
        n.Updated = getTime()
        return n.Updated
}

func (n *InventoryItem) GetId() bson.ObjectId {
        return n.Id
}

func (n *InventoryItem) SetId(id bson.ObjectId) {
        n.Id = id
}

func (n *Inventory) SetCreated() (time.Time) {
        n.Created = getTime()
        return n.Created
}

func (n *Inventory) Update() (time.Time) {
        n.Updated = getTime()
        return n.Updated
}

func (n *Inventory) GetId() bson.ObjectId {
        return n.Id
}

func (n *Inventory) SetId(id bson.ObjectId) {
        n.Id = id
}