package data

import (
        "lulaparty.io/model"
        "gopkg.in/mgo.v2/bson"
        "time"
        "fmt"
)

func CreateUser(u *model.User, p *model.FBProfile, a *model.Address, ac *model.AppConfig) (*model.User, error) {
        sess := ac.CopyDBSession()
        defer sess.Close()
        var t = GenerateCollectionName(ac, model.COLLECTION_USER)
        u.Created = time.Now()
        c := sess.DB(ac.DBName).C(t)

        // first of all we need to create the profile and the address documents
        a, aErr := CreateAddress(a, ac)
        p, pErr := CreateFBProfile(p, ac)
        if(aErr !=nil ) {
                return nil, fmt.Errorf(
                        model.GenerateMessage(model.ERROR_MODEL_CREATE_FAILED, a), aErr)
        }
        if(pErr !=nil ) {
                return nil, fmt.Errorf(
                        model.GenerateMessage(model.ERROR_MODEL_CREATE_FAILED, p), pErr)
        }

        // lets maps the properties back
        u.FBProfile = p.Id
        u.Address = a.Id

        // some timestamps
        tn := time.Now()
        u.Created = tn
        u.Update()

        err := c.Insert(u)
        if err != nil {
                return nil, fmt.Errorf(
                        model.GenerateMessage(model.ERROR_MODEL_CREATE_FAILED, u), aErr)
        }
        return u, nil;
}

func CreateUserSimple(u *model.User, ac *model.AppConfig) (*model.User, error) {
        _, err :=createPersistedEntity(ac, u, model.COLLECTION_USER)
        if(err !=nil ) {
                return nil, err
        }
        return u, nil;
}

func GetUser(id bson.ObjectId, ac *model.AppConfig) (*model.User, error) {
        p := &model.User{}
        err := getHelper(id, p, ac, model.COLLECTION_USER)
        if err != nil {
                return nil, err
        }
        return p, nil
}

func GetUserByEmail(email string, ac *model.AppConfig) (*model.User, error) {
        p := &model.User{}
        err := queryHelperSingle(bson.M{"email": email}, p, ac, model.COLLECTION_USER)
        if err != nil {
                return nil, err
        }
        return p, nil
}

func UpdateUser(user *model.User, ac *model.AppConfig) (*model.User, error) {
        err := updateHelper(user, ac, model.COLLECTION_USER)
        if err != nil {
                return nil, err
        }
        return user, nil
}

func CreateAddress(addr *model.Address, ac *model.AppConfig) (*model.Address, error) {
        _, err :=createPersistedEntity(ac, addr, model.COLLECTION_ADDRESS)
        if(err !=nil ) {
                return nil, err
        }
        return addr, nil;
}

func GetAddress(id bson.ObjectId, ac *model.AppConfig) (*model.Address, error) {
        p := &model.Address{}
        err := getHelper(id, p, ac, model.COLLECTION_ADDRESS)
        if err != nil {
                return nil, err
        }
        return p, nil
}

func UpdateAddress(address *model.Address, ac *model.AppConfig) (*model.Address, error) {
        err := updateHelper(address, ac, model.COLLECTION_ADDRESS)
        if err != nil {
                return nil, err
        }
        return address, nil
}