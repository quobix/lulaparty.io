package data

import (
        "lulaparty.io/model"
        "gopkg.in/mgo.v2/bson"
        "fmt"
        "lulaparty.io/util"
)

func CreateParty(p *model.Party, ac *model.AppConfig) (*model.Party, error) {
        _, err :=createPersistedEntity(ac, p, model.COLLECTION_PARTY)
        if(err !=nil ) {
                return nil, err
        }
        return p, nil;
}


func GetParty(id bson.ObjectId, ac *model.AppConfig) (*model.Party, error) {
        p := &model.Party{}
        err := getHelper(id, p, ac, model.COLLECTION_PARTY)
        if err != nil {
                return nil, err
        }
        return p, nil
}

func UpdateParty(p *model.Party, ac *model.AppConfig) (*model.Party, error) {
        err := updateHelper(p, ac, model.COLLECTION_PARTY)
        if err != nil {
                return nil, err
        }
        return p, nil
}

func DeleteParty(p *model.Party, ac *model.AppConfig) error {
        err :=deletePersistedEntity(ac, p, model.COLLECTION_PARTY)
        if(err !=nil ) {
                return err
        }
        return nil;
}

func AddProviderToParty(id bson.ObjectId, p *model.ProviderProfile, ac *model.AppConfig) (*model.Party, error) {
        pa, err := GetParty(id, ac)
        if(err != nil) {
                return nil, err
        }
        for _, itms := range pa.Providers {
                if(itms.Hex() == p.Id.Hex()) {
                        return nil, fmt.Errorf(model.GenerateMessage(model.ERROR_MODEL_DUPLICATE, p), err)
                }
        }
        _, pErr := GetProviderProfile(p.Id, ac)
        if(pErr != nil) {
                _, pErr = CreateProviderProfile(p, ac)
                if(err != nil) {
                        return nil, pErr
                }
        }
        pa.Providers = append(pa.Providers, p.Id)
        pa, err = UpdateParty(pa, ac)
        if(err != nil) {
                return nil, err
        }
        return pa, nil
}

func RemoveProviderFromParty(id bson.ObjectId, p *model.ProviderProfile, ac *model.AppConfig) (*model.Party, error) {
        pa, err := GetParty(id, ac)
        if(err != nil) {
                return nil, err
        }
        i := 0
        found := false
        for idx, prvid := range pa.Providers {
                if(prvid.Hex() == p.Id.Hex()) {
                        i = idx
                        found = true
                }
        }
        if(!found) {
                return nil, fmt.Errorf(model.GenerateMessage(model.ERROR_MODEL_REFERENCE_MISSING, p), err)
        }
        // slice trick to delete from collection, don't care about
        pa.Providers = util.SliceHelper(pa.Providers,i)
        pa, err = UpdateParty(pa, ac)
        if(err != nil) {
                return nil, err
        }
        return pa, nil
}

func AddHostessToParty(id bson.ObjectId, p *model.HostessProfile, ac *model.AppConfig) (*model.Party, error) {
        pa, err := GetParty(id, ac)
        if(err != nil) {
                return nil, err
        }
        for _, itms := range pa.Hostesses {
                if(itms.Hex() == p.Id.Hex()) {
                        return nil, fmt.Errorf(model.GenerateMessage(model.ERROR_MODEL_DUPLICATE, p), err)
                }
        }
        _, pErr := GetHostessProfile(p.Id, ac)
        if(pErr != nil) {
                _, pErr = CreateHostessProfile(p, ac)
                if(err != nil) {
                        return nil, pErr
                }
        }
        pa.Hostesses = append(pa.Hostesses, p.Id)
        pa, err = UpdateParty(pa, ac)
        if(err != nil) {
                return nil, err
        }
        return pa, nil
}

func RemoveHostessFromParty(id bson.ObjectId, p *model.HostessProfile, ac *model.AppConfig) (*model.Party, error) {
        pa, err := GetParty(id, ac)
        if(err != nil) {
                return nil, err
        }
        i := 0
        found := false
        for idx, prvid := range pa.Hostesses {
                if(prvid.Hex() == p.Id.Hex()) {
                        i = idx
                        found = true
                }
        }
        if(!found) {
                return nil, fmt.Errorf(model.GenerateMessage(model.ERROR_MODEL_REFERENCE_MISSING, p), err)
        }
        // slice trick to delete from collection, don't care about
        pa.Hostesses = util.SliceHelper(pa.Hostesses,i)
        pa, err = UpdateParty(pa, ac)
        if(err != nil) {
                return nil, err
        }
        return pa, nil
}

func CreatePartyInventory(pi *model.PartyInventory, ac *model.AppConfig) (*model.PartyInventory, error) {
        _, err :=createPersistedEntity(ac, pi, model.COLLECTION_PARTY_INVENTORY)
        if(err !=nil ) {
                return nil, err
        }
        return pi, nil;
}

func GetPartyInventory(id bson.ObjectId, ac *model.AppConfig) (*model.PartyInventory, error) {
        p := &model.PartyInventory{}
        err := getHelper(id, p, ac, model.COLLECTION_PARTY_INVENTORY)
        if err != nil {
                return nil, err
        }
        return p, nil
}

func UpdatePartyInventory(inv *model.PartyInventory, ac *model.AppConfig) (*model.PartyInventory, error) {
        err := updateHelper(inv, ac, model.COLLECTION_PARTY_INVENTORY)
        if err != nil {
                return nil, err
        }
        return inv, nil
}

func DeletePartyInventory(i *model.PartyInventory, ac *model.AppConfig) error {
        err :=deletePersistedEntity(ac, i, model.COLLECTION_PARTY_INVENTORY)
        if(err !=nil ) {
                return err
        }
        return nil;
}

func AddItemToPartyInventory(id bson.ObjectId, item *model.InventoryItem,
                                ac *model.AppConfig) (*model.PartyInventory, error) {
        it, err := GetPartyInventory(id, ac)
        if(err != nil) {
                return nil, err
        }
        for _, itms := range it.InventoryItems {
                if(itms.Hex() == item.Id.Hex()) {
                        return nil, fmt.Errorf(model.GenerateMessage(model.ERROR_MODEL_DUPLICATE, item), err)
                }
        }
        _, itmErr := GetInventoryItem(item.Id, ac)
        if(itmErr != nil) {
                _, itmErr = CreateInventoryItem(item, ac)
                if(err != nil) {
                        return nil, itmErr
                }
        }
        it.InventoryItems = append(it.InventoryItems, item.Id)
        it, err = UpdatePartyInventory(it, ac)
        if(err != nil) {
                return nil, err
        }
        return it, nil
}

func RemoveItemFromPartyInventory(id bson.ObjectId, item *model.InventoryItem,
                                        ac *model.AppConfig) (*model.PartyInventory, error) {
        it, err := GetPartyInventory(id, ac)
        if(err != nil) {
                return nil, err
        }
        i := 0
        found := false
        for idx, itms := range it.InventoryItems {
                if(itms.Hex() == item.Id.Hex()) {
                        i = idx
                        found = true
                }
        }
        if(!found) {
                return nil, fmt.Errorf(model.GenerateMessage(model.ERROR_MODEL_REFERENCE_MISSING, item), err)
        }
        // slice trick to delete from collection, don't care about
        it.InventoryItems = util.SliceHelper(it.InventoryItems,i)
        it, err = UpdatePartyInventory(it, ac)
        if(err != nil) {
                return nil, err
        }
        return it, nil
}

func CheckInviteShortCodeIsUnique(code string, ac *model.AppConfig) (bool) {
        p := &model.PartyInvite{}
        err := queryHelperSingle(bson.M{model.MODEL_JSON_SHORTCODE: code}, p,
                        ac, model.COLLECTION_PARTY_INVITE)
        if err == nil {
                return false
        }
        return true
}


func CreatePartyInvite(p *model.PartyInvite, ac *model.AppConfig) (*model.PartyInvite, error) {
        if(!CheckInviteShortCodeIsUnique(p.ShortCode, ac)) {
                return nil, fmt.Errorf("can't create invite, shortcode [%v] is not unique!", p.ShortCode)
        }
        _, err :=createPersistedEntity(ac, p, model.COLLECTION_PARTY_INVITE)
        if(err !=nil ) {
                return nil, err
        }
        return p, nil;
}


func GetPartyInvite(id bson.ObjectId, ac *model.AppConfig) (*model.PartyInvite, error) {
        p := &model.PartyInvite{}
        err := getHelper(id, p, ac, model.COLLECTION_PARTY_INVITE)
        if err != nil {
                return nil, err
        }
        return p, nil
}

func GetPartyInviteByShortCode(code string, ac *model.AppConfig) (*model.PartyInvite, error) {
        p := &model.PartyInvite{}
        err := queryHelperSingle(bson.M{model.MODEL_JSON_SHORTCODE: code}, p,
                ac, model.COLLECTION_PARTY_INVITE)
        if err != nil {
                return nil, err
        }
        return p, nil
}

func UpdatePartyInvite(p *model.PartyInvite, ac *model.AppConfig) (*model.PartyInvite, error) {
        err := updateHelper(p, ac, model.COLLECTION_PARTY_INVITE)
        if err != nil {
                return nil, err
        }
        return p, nil
}

func DeletePartyInvite(p *model.PartyInvite, ac *model.AppConfig) error {
        err :=deletePersistedEntity(ac, p, model.COLLECTION_PARTY_INVITE)
        if(err !=nil ) {
                return err
        }
        return nil;
}

func CreateReward(p *model.Reward, ac *model.AppConfig) (*model.Reward, error) {
        _, err :=createPersistedEntity(ac, p, model.COLLECTION_REWARD)
        if(err !=nil ) {
                return nil, err
        }
        return p, nil;
}

func GetReward(id bson.ObjectId, ac *model.AppConfig) (*model.Reward, error) {
        p := &model.Reward{}
        err := getHelper(id, p, ac, model.COLLECTION_REWARD)
        if err != nil {
                return nil, err
        }
        return p, nil
}


func UpdateReward(p *model.Reward, ac *model.AppConfig) (*model.Reward, error) {
        err := updateHelper(p, ac, model.COLLECTION_REWARD)
        if err != nil {
                return nil, err
        }
        return p, nil
}

func DeleteReward(p *model.Reward, ac *model.AppConfig) error {
        err :=deletePersistedEntity(ac, p, model.COLLECTION_REWARD)
        if(err !=nil ) {
                return err
        }
        return nil;
}