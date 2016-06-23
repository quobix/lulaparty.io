package service

import (
        "lulaparty.io/model"
        "os"
        "lulaparty.io/gcp"
        "fmt"
        "lulaparty.io/data"
        "lulaparty.io/util"
)

func PersistGalleryItemToStorage(g *model.GalleryItem, f *os.File, ac *model.AppConfig) (*model.GalleryItem, error) {
        service, err := gcp.CreateStorageService()
        if(err!=nil) {
                return nil, fmt.Errorf("unable to create a storage service! %v", err)
        }

        /* check if gallery item exists in storage */
        ret_gi, err := data.GetGalleryItem(g.Id, ac)
        if(err!=nil) {
                return nil, fmt.Errorf("can't find gallery item in storage with id: %v", g.Id)
        }

        uri := util.GenerateRawGalleryItemUUID(g.OwnerId, g.GalleryId, g.Id, f)
        _, err = gcp.UploadObjectToBucketUsingName(model.BUCKET_GALLERY, f, uri, service, ac)
        if(err!=nil) {
                return nil, fmt.Errorf("unable to persist file to bucket! %v", err)
        }


        err = gcp.MakeObjectPublicReadable(model.BUCKET_GALLERY, uri, service, ac)
        if(err!=nil) {
                return nil, fmt.Errorf("unable to make item publicly readable %v", err)
        }

        // update gallery item with persisted storage
        g.FileUUID = uri
        _, err = data.UpdateGalleryItem(g, ac);
        if(err!=nil) {
                return nil, fmt.Errorf("unable to update gallery item %v", err)
        }

        // update gallery
        gal, err := data.GetGallery(g.GalleryId, ac)
        if(err!=nil) {
                return nil, fmt.Errorf("unable to retrive gallery %v", err)
        }

        gal, err = data.AddItemToGallery(gal.Id, g, ac)
        if(err!=nil) {
                return nil, fmt.Errorf("unable to update gallery with new item %v", err)
        }

        return ret_gi, nil;
}

func RemoveGalleryItemFromStorage(g *model.GalleryItem, ac *model.AppConfig) error {
        service, err := gcp.CreateStorageService()
        if(err!=nil) {
                return fmt.Errorf("unable to create a storage service! %v", err)
        }

        /* check if gallery item exists in storage */
        _, err = data.GetGalleryItem(g.Id, ac)
        if(err!=nil) {
                return fmt.Errorf("can't find gallery item in storage with id: %v", g.Id)
        }

        err = gcp.DeleteObjectInBucket(model.BUCKET_GALLERY, g.FileUUID, service, ac)
        if(err!=nil) {
                return fmt.Errorf("unable to delete object in bucket! %v, %v", g.FileUUID, err)
        }

        _, err = data.RemoveItemFromGallery(g.GalleryId, g, ac)
        if(err!=nil) {
                return fmt.Errorf("unable to remove the item [%v]from the gallery %v", g.GalleryId,err)
        }

        err = data.DeleteGalleryItem(g, ac)
        if(err!=nil) {
                return fmt.Errorf("unable to delete the gallery item %v", err)
        }
        return nil;
}