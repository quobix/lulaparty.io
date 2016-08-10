package service

import (
    "github.com/quobix/lulaparty.io/model"
    "github.com/goinggo/tracelog"
    "github.com/quobix/lulaparty.io/data"
    "time"
)

func CheckForOrCreateExistingProfile(ter *model.TokenExchangeRequest,
    at *model.AccessToken, ac *model.AppConfig) (*model.User, error) {
    u, err := data.GetUserByEmail(ter.Email, ac)
    if (err != nil) {
        tracelog.Trace("service", "AuthenticateUser", "New user creation: [" + ter.Email + "]")
        u, err = CreateNewUserFromTex(ter, at, ac)
        if (err != nil) {
            return nil, err
        }
    }
    return u, nil
}

func CreateNewUserFromTex(ter *model.TokenExchangeRequest, at *model.AccessToken,
    ac *model.AppConfig) (*model.User, error) {
    fbp := &model.FBProfile {
        Firstname:  ter.Firstname,
        Lastname:   ter.Lastname,
        Name:       ter.Firstname + " " + ter.Lastname,
        Email:      ter.Email,
    }
    tracelog.Trace("service", "CreateNewUserFromTex", "Creating new FB profile for new user [" + ter.Email + "]")
    ret_fbp, err := data.CreateFBProfile(fbp, ac)
    if(err!=nil) {
        return nil, err
    }
    tracelog.Trace("service", "CreateNewUserFromTex", "Adding (at) for new user [" + ter.Email + "]")
    ret_at, err := data.AddAccessTokenToFBProfile(ret_fbp, at, ac)
    if(err!=nil) {
        return nil, err
    }
    fbp.AccessToken = ret_at.Id; // we have to wire this
    tracelog.Trace("service", "CreateNewUserFromTex", "Updating FB(p) to reference (at) [" + ter.Email + "]")
    ret_fbp, err  = data.UpdateFBProfile(fbp, ac)
    if(err!=nil) {
        return nil, err
    }
    u := &model.User {
        FBProfile:      ret_fbp.Id,
        Email:          ter.Email,
        LastAuth:       time.Now(),
        TrialExpired:   false,
        FBAuthToken:    ret_at.Id,
    }
    tracelog.Trace("service", "CreateNewUserFromTex", "Creating actual new user [" + ter.Email + "]")
    ret_u, err := data.CreateUserSimple(u, ac)
    if(err!=nil) {
        return nil, err
    }
    return ret_u, nil
}