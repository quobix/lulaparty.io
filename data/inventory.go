package data

import (
        "github.com/quobix/lulaparty.io/model"
        "gopkg.in/mgo.v2/bson"
        "fmt"
)

func CreateInventory(addr *model.Inventory, ac *model.AppConfig) (*model.Inventory, error) {
        _, err :=createPersistedEntity(ac, addr, model.COLLECTION_INVENTORY)
        if(err !=nil ) {
                return nil, err
        }
        return addr, nil;
}


func GetInventory(id bson.ObjectId, ac *model.AppConfig) (*model.Inventory, error) {
        p := &model.Inventory{}
        err := getHelper(id, p, ac, model.COLLECTION_INVENTORY)
        if err != nil {
                return nil, err
        }
        return p, nil
}

func UpdateInventory(inv *model.Inventory, ac *model.AppConfig) (*model.Inventory, error) {
        err := updateHelper(inv, ac, model.COLLECTION_INVENTORY)
        if err != nil {
                return nil, err
        }
        return inv, nil
}

func DeleteInventory(i *model.Inventory, ac *model.AppConfig) error {
        err :=deletePersistedEntity(ac, i, model.COLLECTION_INVENTORY)
        if(err !=nil ) {
                return err
        }
        return nil;
}

func CreateInventoryItem(item *model.InventoryItem, ac *model.AppConfig) (*model.InventoryItem, error) {
        _, err :=createPersistedEntity(ac, item, model.COLLECTION_INVENTORY_ITEM)
        if(err !=nil ) {
                return nil, err
        }
        return item, nil;
}

func GetInventoryItem(id bson.ObjectId, ac *model.AppConfig) (*model.InventoryItem, error) {
        p := &model.InventoryItem{}
        err := getHelper(id, p, ac, model.COLLECTION_INVENTORY_ITEM)
        if err != nil {
                return nil, err
        }
        return p, nil
}

func UpdateInventoryItem(inv *model.InventoryItem, ac *model.AppConfig) (*model.InventoryItem, error) {
        err := updateHelper(inv, ac, model.COLLECTION_INVENTORY_ITEM)
        if err != nil {
                return nil, err
        }
        return inv, nil
}

func DeleteInventoryItem(i *model.InventoryItem, ac *model.AppConfig) error {
        err :=deletePersistedEntity(ac, i, model.COLLECTION_INVENTORY_ITEM)
        if(err !=nil ) {
                return err
        }
        return nil;
}

func AddItemToInventory(id bson.ObjectId, item *model.InventoryItem, ac *model.AppConfig) (*model.Inventory, error) {
        it, err := GetInventory(id, ac)
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
        it, err = UpdateInventory(it, ac)
        if(err != nil) {
                return nil, err
        }
        return it, nil
}

func RemoveItemFromInventory(id bson.ObjectId, item *model.InventoryItem, ac *model.AppConfig) (*model.Inventory, error) {
        it, err := GetInventory(id, ac)
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
        sl := append(it.InventoryItems[:i], it.InventoryItems[i+1:]...)
        it.InventoryItems = sl // re-assign
        it, err = UpdateInventory(it, ac)
        if(err != nil) {
                return nil, err
        }
        return it, nil
}


func GetUserInventoryItems(u *model.User, rev bool, ac *model.AppConfig) ([]model.InventoryItem, error) {
        sess := ac.CopyDBSession()
        defer sess.Close()

        r := []model.InventoryItem{}
        q := bson.M{model.MODEL_JSON_OWNERID: u.Id}

        c := getEntityCollection(ac, sess, model.COLLECTION_INVENTORY_ITEM)
        err := c.Find(q).Sort(createSort(rev, model.MODEL_JSON_CREATED)).All(&r)

        if err != nil {
                return nil, fmt.Errorf(model.GenerateMessage(model.ERROR_MODEL_QUERY_FAILED, r), q, err)
        }
        return r, nil;

}

func GetInventoryItemsById(i bson.ObjectId, s string, rev bool, ac *model.AppConfig) ([]model.InventoryItem, error) {
        sess := ac.CopyDBSession()
        defer sess.Close()

        r := []model.InventoryItem{}
        q := bson.M{model.MODEL_JSON_INVENTORYID: i}

        c := getEntityCollection(ac, sess, model.COLLECTION_INVENTORY_ITEM)
        err := c.Find(q).Sort(createSort(rev, s)).All(&r)

        if err != nil {
                return nil, fmt.Errorf(model.GenerateMessage(model.ERROR_MODEL_QUERY_FAILED, r), q, err)
        }
        return r, nil;
}

func GetInventoryItems(i *model.Inventory, s string, rev bool, ac *model.AppConfig) ([]model.InventoryItem, error) {
        return GetInventoryItemsById(i.Id, s, rev, ac)
}