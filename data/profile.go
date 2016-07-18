package data

import (
        "github.com/quobix/lulaparty.io/model"
        "gopkg.in/mgo.v2/bson"
        "time"
        "fmt"
)

func CreateCustomerProfile(profile *model.CustomerProfile, ac *model.AppConfig) (*model.CustomerProfile, error) {
        _, err := createPersistedEntity(ac, profile, model.COLLECTION_CUSTOMER_PROFILE)
        if(err !=nil ) {
                return nil, err
        }
        return profile, nil
}

func GetCustomerProfile(id bson.ObjectId, ac *model.AppConfig) (*model.CustomerProfile, error) {
        p := &model.CustomerProfile{}
        err := getHelper(id, p, ac, model.COLLECTION_CUSTOMER_PROFILE)
        if err != nil {
                return nil, err
        }
        return p, nil
}

func UpdateCustomerProfile(profile *model.CustomerProfile, ac *model.AppConfig) (*model.CustomerProfile, error) {
        err := updateHelper(profile, ac, model.COLLECTION_CUSTOMER_PROFILE)
        if err != nil {
                return nil, err
        }
        return profile, nil
}

func CreateFBProfile(profile *model.FBProfile, ac *model.AppConfig) (*model.FBProfile, error) {
        _, err := createPersistedEntity(ac, profile, model.COLLECTION_FBPROFILE)
        if(err !=nil ) {
                return nil, err
        }
        return profile, nil
}

func CreateAccessToken(at *model.AccessToken, ac *model.AppConfig) (*model.AccessToken, error) {
        _, err := createPersistedEntity(ac, at, model.COLLECTION_ACCESSTOKEN)
        if(err !=nil ) {
                return nil, err
        }
        return at, nil
}


func GetAccessToken(id bson.ObjectId, ac *model.AppConfig) (*model.AccessToken, error) {
        p := &model.AccessToken{}
        err := getHelper(id, p, ac, model.COLLECTION_ACCESSTOKEN)
        if err != nil {
                return nil, err
        }
        return p, nil
}

func AddAccessTokenToFBProfile(profile *model.FBProfile, at *model.AccessToken, ac *model.AppConfig) (*model.FBProfile, error) {
        sess := ac.CopyDBSession()
        defer sess.Close()

        // first of all we need to create the profile and the address documents
        at, aErr := CreateAccessToken(at, ac)
        if(aErr !=nil ) {
                return nil, fmt.Errorf(
                        model.GenerateMessage(model.ERROR_MODEL_CREATE_FAILED, at), aErr)
        }

        profile.AccessToken = at.Id
        tn := time.Now()
        profile.Updated = tn

        _, err := UpdateFBProfile(profile, ac)
        if err != nil {
                return nil, fmt.Errorf(
                        model.GenerateMessage(model.ERROR_MODEL_UPDATE_FAILED, profile), aErr)
        }
        return profile, nil;
}


func GetFBProfile(id bson.ObjectId, ac *model.AppConfig) (*model.FBProfile, error) {
        p := &model.FBProfile{}
        err := getHelper(id, p, ac, model.COLLECTION_FBPROFILE)
        if err != nil {
                return nil, err
        }
        return p, nil
}

func UpdateFBProfile(profile *model.FBProfile, ac *model.AppConfig) (*model.FBProfile, error) {
        err := updateHelper(profile, ac, model.COLLECTION_FBPROFILE)
        if err != nil {
                return nil, err
        }
        return profile, nil
}

func CreateHostessProfile(profile *model.HostessProfile, ac *model.AppConfig) (*model.HostessProfile, error) {
        _, err := createPersistedEntity(ac, profile, model.COLLECTION_HOSTESS_PROFILE)
        if(err !=nil ) {
                return nil, err
        }
        return profile, nil
}

func GetHostessProfile(id bson.ObjectId, ac *model.AppConfig) (*model.HostessProfile, error) {
        p := &model.HostessProfile{}
        err := getHelper(id, p, ac, model.COLLECTION_HOSTESS_PROFILE)
        if err != nil {
                return nil, err
        }
        return p, nil
}

func UpdateHostessProfile(profile *model.HostessProfile, ac *model.AppConfig) (*model.HostessProfile, error) {
        err := updateHelper(profile, ac, model.COLLECTION_HOSTESS_PROFILE)
        if err != nil {
                return nil, err
        }
        return profile, nil
}

func CreateProviderProfile(profile *model.ProviderProfile, ac *model.AppConfig) (*model.ProviderProfile, error) {
        _, err := createPersistedEntity(ac, profile, model.COLLECTION_PROVIDER_PROFILE)
        if(err !=nil ) {
                return nil, err
        }
        return profile, nil
}

func GetProviderProfile(id bson.ObjectId, ac *model.AppConfig) (*model.ProviderProfile, error) {
        p := &model.ProviderProfile{}
        err := getHelper(id, p, ac, model.COLLECTION_PROVIDER_PROFILE)
        if err != nil {
                return nil, err
        }
        return p, nil
}

func UpdateProviderProfile(profile *model.ProviderProfile, ac *model.AppConfig) (*model.ProviderProfile, error) {
        err := updateHelper(profile, ac, model.COLLECTION_PROVIDER_PROFILE)
        if err != nil {
                return nil, err
        }
        return profile, nil
}

