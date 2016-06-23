package data

import (
        "github.com/quobix/lulaparty.io/model"
        "gopkg.in/mgo.v2/bson"
        "fmt"
)

func CreateGallery(g *model.Gallery, ac *model.AppConfig) (*model.Gallery, error) {
        _, err :=createPersistedEntity(ac, g, model.COLLECTION_GALLERY)
        if(err !=nil ) {
                return nil, err
        }
        return g, nil;
}


func GetGallery(id bson.ObjectId, ac *model.AppConfig) (*model.Gallery, error) {
        g := &model.Gallery{}
        err := getHelper(id, g, ac, model.COLLECTION_GALLERY)
        if err != nil {
                return nil, err
        }
        return g, nil
}

func UpdateGallery(g *model.Gallery, ac *model.AppConfig) (*model.Gallery, error) {
        err := updateHelper(g, ac, model.COLLECTION_GALLERY)
        if err != nil {
                return nil, err
        }
        return g, nil
}

func CreateGalleryItem(gi *model.GalleryItem, ac *model.AppConfig) (*model.GalleryItem, error) {
        _, err :=createPersistedEntity(ac, gi, model.COLLECTION_GALLERY_ITEM)
        if(err !=nil ) {
                return nil, err
        }
        return gi, nil;
}


func GetGalleryItem(id bson.ObjectId, ac *model.AppConfig) (*model.GalleryItem, error) {
        gi := &model.GalleryItem{}
        err := getHelper(id, gi, ac, model.COLLECTION_GALLERY_ITEM)
        if err != nil {
                return nil, err
        }
        return gi, nil
}

func UpdateGalleryItem(gi *model.GalleryItem, ac *model.AppConfig) (*model.GalleryItem, error) {
        err := updateHelper(gi, ac, model.COLLECTION_GALLERY_ITEM)
        if err != nil {
                return nil, err
        }
        return gi, nil
}

func DeleteGallery(g *model.Gallery, ac *model.AppConfig) error {
        err :=deletePersistedEntity(ac, g, model.COLLECTION_GALLERY)
        if(err !=nil ) {
                return err
        }
        return nil;
}

func DeleteGalleryItem(g *model.GalleryItem, ac *model.AppConfig) error {
        err :=deletePersistedEntity(ac, g, model.COLLECTION_GALLERY_ITEM)
        if(err !=nil ) {
                return err
        }
        return nil;
}

func AddItemToGallery(id bson.ObjectId, item *model.GalleryItem, ac *model.AppConfig) (*model.Gallery, error) {
        it, err := GetGallery(id, ac)
        if(err != nil) {
                return nil, err
        }
        for _, itms := range it.GalleryItems {
                if(itms.Hex() == item.Id.Hex()) {
                        return nil, fmt.Errorf(model.GenerateMessage(model.ERROR_MODEL_DUPLICATE, item), err)
                }
        }
        _, itmErr := GetGalleryItem(item.Id, ac)
        if(itmErr != nil) {
                _, itmErr = CreateGalleryItem(item, ac)
                if(err != nil) {
                        return nil, itmErr
                }
        }
        it.GalleryItems = append(it.GalleryItems, item.Id)
        it, err = UpdateGallery(it, ac)
        if(err != nil) {
                return nil, err
        }
        return it, nil
}

func RemoveItemFromGallery(id bson.ObjectId, item *model.GalleryItem, ac *model.AppConfig) (*model.Gallery, error) {
        it, err := GetGallery(id, ac)
        if(err != nil) {
                return nil, fmt.Errorf("no gallery found! %v", err)
        }
        i := 0
        found := false
        for idx, itms := range it.GalleryItems {
                if(itms.Hex() == item.Id.Hex()) {
                        i = idx
                        found = true
                }
        }
        if(!found) {
                return nil, fmt.Errorf(model.GenerateMessage(model.ERROR_MODEL_REFERENCE_MISSING, item), err)
        }
        // slice trick to delete from collection, don't care about
        sl := append(it.GalleryItems[:i], it.GalleryItems[i+1:]...)
        it.GalleryItems = sl // re-assign
        it, err = UpdateGallery(it, ac)
        if(err != nil) {
                return nil, err
        }
        return it, nil
}


func galleryItemHelper(i bson.ObjectId, s string, rev bool, ac *model.AppConfig, q bson.M) ([]model.GalleryItem, error) {
        sess := ac.CopyDBSession()
        defer sess.Close()

        r := []model.GalleryItem{}

        c := getEntityCollection(ac, sess, model.COLLECTION_GALLERY_ITEM)
        err := c.Find(q).Sort(createSort(rev, s)).All(&r)

        if err != nil {
                return nil, fmt.Errorf(model.GenerateMessage(model.ERROR_MODEL_QUERY_FAILED, r), q, err)
        }
        return r, nil;
}

func GetGalleryItemsByGalleryId(i bson.ObjectId, s string, rev bool, ac *model.AppConfig) ([]model.GalleryItem, error) {
       return galleryItemHelper(i, s, rev, ac, bson.M{model.MODEL_JSON_GALLERYID: i})
}

func GetGalleryItemsByUserId(i bson.ObjectId, s string, rev bool, ac *model.AppConfig) ([]model.GalleryItem, error) {
        return galleryItemHelper(i, s, rev, ac, bson.M{model.MODEL_JSON_OWNERID: i})
}

func GetGalleryItems(i *model.Gallery, s string, rev bool, ac *model.AppConfig) ([]model.GalleryItem, error) {
        return GetGalleryItemsByGalleryId(i.Id, s, rev, ac)
}