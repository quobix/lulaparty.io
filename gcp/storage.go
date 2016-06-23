package gcp

import (
        storage "google.golang.org/api/storage/v1"
        "github.com/quobix/lulaparty.io/model"

        "fmt"
        "os"
        "path"
        "github.com/quobix/lulaparty.io/util"
)

const (
        // This scope allows the application full control over resources in Google Cloud Storage
        storageScope    = storage.DevstorageFullControlScope
        storageURIBase  = "https://storage.googleapis.com"
)

func GenerateObjectURI(b string, o string, ac *model.AppConfig, ) string {
        return storageURIBase + util.FILE_UUID_FSSEP + bucketNameFilter(b, ac) + util.FILE_UUID_FSSEP + o
}

func bucketNameFilter(bn string, ac *model.AppConfig) string {
        if(ac.TestMode) {
                bn += model.COLLECTION_TEST_POSTFIX
        }
        return bn
}

func getProjectID(ac *model.AppConfig) string {
        if(ac.TestMode) {
                return PROJECT_ID_TEST
        }
        return PROJECT_ID
}

func CreateStorageService() (*storage.Service , error) {
        cl, err := CreateClient(storageScope)
        service, err := storage.New(cl)
        if err != nil {
                return nil, err
        }
        return service, nil
}

func CreateBucket(b string, s *storage.Service, ac *model.AppConfig) (*storage.Bucket, error) {
        buck, err := GetBucket(b, s, ac)
        if(err!=nil) { // if the get throws an error... we can create, otherwise skip.
                if res, err := s.Buckets.Insert(getProjectID(ac),
                                &storage.Bucket{Name: bucketNameFilter(b, ac)}).Do(); err == nil {
                        return res, nil

                } else {
                       return nil, err
                }
        }
        return buck, nil
}

func GetBucket(b string, s *storage.Service, ac *model.AppConfig) (*storage.Bucket, error) {
        buck, err := s.Buckets.Get(bucketNameFilter(b, ac)).Do();
        if(err!=nil) { return nil, err}
        return buck, nil
}

func DeleteBucket(b string, s *storage.Service, ac *model.AppConfig) error {
        if err := s.Buckets.Delete(bucketNameFilter(b, ac)).Do(); err != nil {
                return fmt.Errorf("unable to delete bucket: %v", err)
        }
        return nil
}

func DeleteObjectInBucket(b string, objectName string, s *storage.Service, ac *model.AppConfig) error {

        if err := s.Objects.Delete(bucketNameFilter(b, ac), objectName).Do(); err != nil {
                return fmt.Errorf("unable to delete object: [%v], %v", objectName, err)
        }
        return nil
}

func ListBuckets(s *storage.Service, ac *model.AppConfig) ([]*storage.Bucket, error) {
        if res, err := s.Buckets.List(getProjectID(ac)).Do(); err == nil {
                return res.Items, nil
        } else {
                return nil, err
        }
}

func UploadObjectToBucketUsingName(bucket string, file *os.File,
        name string, s *storage.Service, ac *model.AppConfig) (*storage.Object, error) {
        if(file==nil) {
                return nil, fmt.Errorf("can't upload, missing file pointer!")
        }
        obj := &storage.Object { Name: name }
        if res, err := s.Objects.Insert(bucketNameFilter(bucket, ac), obj).Media(file).Do(); err == nil {
               return res, nil
        } else {
               return nil, fmt.Errorf("unable to insert new object to bucket: %v", err)
        }
}

func UploadObjectToBucket(bucket string, file *os.File,
               s *storage.Service, ac *model.AppConfig) (*storage.Object, error) {
        if(file==nil) {
                return nil, fmt.Errorf("can't upload, missing file pointer!")
        }
        return UploadObjectToBucketUsingName(bucket, file, path.Base(file.Name()), s, ac)
}

func MakeObjectPublicReadable(bucket string, objectName string, s *storage.Service, ac *model.AppConfig) error {
        _, err :=s.ObjectAccessControls.Update(bucketNameFilter(bucket, ac), objectName,
                "allUsers", &storage.ObjectAccessControl{Role: "READER", Object: objectName}).Do();
        if(err != nil){
                return err
        }
        return nil
}

func GetObjectFromBucket(bucket string, objectName string, s *storage.Service, ac *model.AppConfig) (*storage.Object, error) {
        res, err := s.Objects.Get(bucketNameFilter(bucket, ac), objectName).Do();
        if(err!=nil) {
                return nil, err
        }
        return res, nil
}