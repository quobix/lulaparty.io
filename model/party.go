package model

import (
        "gopkg.in/mgo.v2/bson"
        "time"
        "fmt"
        "math"
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

func round(f float64) float64 {
        return math.Floor(f + .5)
}
func roundPlus(f float64, places int) (int) {
        shift := math.Pow(10, float64(places))
        return int(round(f * shift) / shift);
}

type timeDiffClosure func(t time.Time) bool

func timeDiffFormatter(h,m,s,d,mo,y int) string {
        var pl = "s"
        h1 := h-(d*24)
        if (y <= 0 && mo >= 1) {
                return fmt.Sprintf("%d " + pluralize("Month", pl, mo), mo)
        }
        if (y <= 0 && d >= 1 && h1 <=0) {
                return fmt.Sprintf("%d " + pluralize("Day", pl, d), d)
        }

        if (y <= 0 && d >= 1 && h1>=1) {
                return fmt.Sprintf("%d " + pluralize("Day", pl, d) + ", %d " +
                pluralize("Hour", pl,h1), d, h1)
        }
        if (y <= 0 && d <= 0 && h >= 1 && m <= 0) {
                return fmt.Sprintf("%d " + pluralize("Hour", pl, h), h)
        }

        if (y <= 0 && d <= 0 && h >= 1 ) {
                return fmt.Sprintf("%d " + pluralize("Hour", pl, h) + ", %d " +
                pluralize("Min", pl, m), h, m)
        }

        if (y <= 0 && h <= 0 && m >= 1 && s >= 1) {
                return fmt.Sprintf("%d " + pluralize("Minute", pl, m) + ", %d " +
                pluralize("Second", pl, s), m, s)
        }

        if (y <= 0 && h <= 0 && m >= 1) {
                return fmt.Sprintf("%d " + pluralize("Minute", pl, m), m)
        }

        if (y <= 0 && m <= 0 && s >= 1) {
                return fmt.Sprintf("%d " + pluralize("Second", pl, s), s)
        }
        return fmt.Sprintf("%d " + pluralize("Year", pl, y), y)
}

func timeDiffHelper(diff time.Time, def string, m timeDiffClosure) string {

        t := time.Now().UTC()
        d := t.Sub(diff)
        seconds := d.Seconds()
        hours := d.Hours()
        minutes := d.Minutes()

        if (m(diff)) {
                h := roundPlus(hours * -1, 0)
                m := roundPlus(minutes * -1, 0) - (h * 60)
                s := roundPlus(seconds * -1, 0) - (m * 60)
                d := roundPlus(hours * -1, 0) / 24
                mo := d / 30
                y := mo / 12
                return timeDiffFormatter(h,m,s,d,mo,y)
        }
        return def
}

func (n *Party) StartsIn() string {
        if(n.IsOpen()) {
                return PARTY_STARTED
        }
        return timeDiffHelper(n.Starts, PARTY_STARTED,
                        func(t time.Time) bool {
                                return t.After(time.Now().UTC())
                        })
}

func (n *Party) EndsIn() string {
        if(n.IsOpen()) {
                return timeDiffHelper(n.Ends, PARTY_ENDED,
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
                return timeDiffHelper(n.Expires, PARTY_INVITE_EXPIRED,
                        func(t time.Time) bool {
                                return time.Now().UTC().Before(t)
                        })
        } else {
                return PARTY_INVITE_ACCEPTED
        }
}

func pluralize(val, pl string, num int) string {
        if(num>1) { return val + pl }
        return val
}
