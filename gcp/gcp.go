package gcp

import (
        "golang.org/x/oauth2/google"
        "golang.org/x/net/context"
        "net/http"
)

const (
        PROJECT_ID = "lulaparty-io"
        PROJECT_ID_TEST = "lulaparty-io-test"

)


func CreateClient(scope string) (*http.Client, error) {
        client, err := google.DefaultClient(context.Background(), scope)
        if err != nil { return nil, err }
        return client, nil
}