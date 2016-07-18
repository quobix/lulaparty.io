package model

import (
        "gopkg.in/mgo.v2"
        "gopkg.in/mgo.v2/bson"
        "time"
        "reflect"
)

const (
        COLLECTION_TEST_POSTFIX         = "_test"
        COLLECTION_FBPROFILE            = "fbprofile"
        COLLECTION_ACCESSTOKEN          = "accesstoken"
        COLLECTION_ADDRESS              = "address"
        COLLECTION_USER                 = "user"
        COLLECTION_CUSTOMER_PROFILE     = "customer_profile"
        COLLECTION_HOSTESS_PROFILE      = "hostess_profile"
        COLLECTION_PROVIDER_PROFILE     = "provider_profile"
        COLLECTION_INVENTORY            = "inventory"
        COLLECTION_INVENTORY_ITEM       = "inventory_item"
        COLLECTION_INVENTORY_ITEM_SIZES = "inventory_item_sizes"
        COLLECTION_PARTY                = "party"
        COLLECTION_PARTY_INVENTORY      = "party_inventory"
        COLLECTION_PARTY_INVITE         = "party_invite"
        COLLECTION_REWARD               = "reward"
        COLLECTION_GALLERY              = "gallery"
        COLLECTION_GALLERY_ITEM         = "gallery_item"
        MODEL_JSON_OWNERID              = "owner_id"
        MODEL_JSON_ID                   = "id"
        MODEL_JSON_CREATED              = "created"
        MODEL_JSON_UPDATED              = "updated"
        MODEL_JSON_NAME                 = "name"
        MODEL_JSON_CAPTION              = "caption"
        MODEL_JSON_SHORTCODE            = "shortcode"
        BUCKET_GALLERY                  = "gallery"
)

type PersistedEntity interface {
        SetCreated()                    time.Time
        Update()                        time.Time
        GetId()                         bson.ObjectId
        SetId(bson.ObjectId)
}

type AppConfig struct {
        DBHost                          string          `json:"dbhost"`
        DBUser                          string          `json:"dbuser"`
        DBPassword                      string          `json:"-"`
        DBName                          string          `json:"dbname"`
        DBPort                          int             `json:"dbport"`
        DBSession                       *mgo.Session    `json:"-"`
        TestMode                        bool
        JTWSecret                       string          `json:"-"`
}

type ServiceResponse struct {
        Error   	                bool		`json:"error"`
        Message 	                string		`json:"message"`
}

type TokenExchangeRequest struct {
        AccessToken   	                string        	`json:"accessToken"`
        ExpiresIn     	                int64        	`json:"expiresIn"`
        SignedRequest 	                string        	`json:"signedRequest"`
        UserId        	                string        	`json:"userId"`
        Email		                string		`json:"email"`
}

func (c *AppConfig) CopyDBSession() (*mgo.Session) {
        if(c.DBSession == nil) {
                panic("no DB session in app config, unable to create a new copy")
        }
        if(c.DBSession == nil) { return nil }
        return c.DBSession.Copy();
}

func getTime() (time.Time) {
        return time.Now();
}

func GenerateMessage(l string, m interface{}) string {
        return l + " " + reflect.TypeOf(m).String() + ": %s"
}
