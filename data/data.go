package data

import (
        "gopkg.in/mgo.v2"
        "strconv"
        "github.com/quobix/lulaparty.io/model"
        "fmt"
        "gopkg.in/mgo.v2/bson"
        "os"
)

const (
        NoEntityFound                   string = "No document with the Id [%s] could be found in collection [%s]: %s"
        UnableToCreateEntity            string = "Unable to create document with Id [%s] in collection [%s]: %s"
        UnableToRetrieveEntity          string = "Unable to updated document with iD [%s] in collection [%s]: %s"
)


func CreateAppConfig(testMode bool) *model.AppConfig {

        port, _ := strconv.Atoi(os.Getenv("LLP_TEST_DB_PORT"))
        return &model.AppConfig{
                DBName: os.Getenv("LLP_TEST_DB"),
                DBUser: os.Getenv("LLP_TEST_DB_USER"),
                DBPassword: os.Getenv("LLP_TEST_DB_PASS"),
                DBPort: port,
                DBHost: os.Getenv("LLP_TEST_DB_HOST"),
                DBSession: ConnectDB(os.Getenv("LLP_TEST_DB_USER"), os.Getenv("LLP_TEST_DB_PASS"),
                        os.Getenv("LLP_TEST_DB_HOST"), port,
                        os.Getenv("LLP_TEST_DB")),
                TestMode: testMode }

}

func GenerateURI(user string, pass string, host string, port int, db string) string {
        return  "mongodb://" + user + ":" + pass + "@" + host + ":" + strconv.Itoa(port) + "/" + db
}

func ConnectDB(user string, pass string, host string, port int, db string) (*mgo.Session){
        uri :=GenerateURI(user, pass, host, port, db)
        session, err := mgo.Dial(uri)
        if err != nil {
                panic(err)
        }

        // Optional. Switch the session to a monotonic behavior.
        session.SetMode(mgo.Monotonic, true)
        return session
}

func GetDB(sess *mgo.Session, db string) (*mgo.Database) {
        return sess.DB(db);
}

func GetCollection(sess *mgo.Session, db string, collection string) (*mgo.Collection) {
        return GetDB(sess, db).C(collection)
}

func GenerateCollectionName(ac *model.AppConfig, collection string) string {
        if(ac.TestMode) {
                collection = collection + model.COLLECTION_TEST_POSTFIX
        }
        return collection
}

func CreateTestSession() (*model.AppConfig) {


        ac := CreateAppConfig(true)

        ac.DBSession.DB(ac.DBName).DropDatabase() // cleanup.
        return ac

}

func createPersistedEntity(ac *model.AppConfig,
e model.PersistedEntity, collection string) (*model.PersistedEntity, error){
        sess := ac.CopyDBSession()
        defer sess.Close()
        var t = GenerateCollectionName(ac, collection)

        e.SetCreated()
        c := sess.DB(ac.DBName).C(t)

        // in case we don't have an id assigned yet.
        if(e.GetId().Hex()=="") {
                e.SetId(bson.NewObjectId())
        }

        err := c.Insert(e)

        if err != nil {
                return nil, fmt.Errorf(model.GenerateMessage(model.ERROR_MODEL_CREATE_FAILED, e), err)
        }
        return &e, nil
}

func deletePersistedEntity(ac *model.AppConfig, e model.PersistedEntity, collection string) error {
        sess := ac.CopyDBSession()
        defer sess.Close()
        var t = GenerateCollectionName(ac, collection)


        c := sess.DB(ac.DBName).C(t)
        err := c.RemoveId(e.GetId())

        if err != nil {
                return fmt.Errorf(model.GenerateMessage(model.ERROR_MODEL_DELETE_FAILED, e), err)
        }
        return nil
}



/* get the collection, but don't manage a session, let the handler deal with it. */
func getEntityCollection(ac *model.AppConfig, sess *mgo.Session, collection string) *mgo.Collection {
        var t = GenerateCollectionName(ac, collection)
        return sess.DB(ac.DBName).C(t)
}

/**
CRUD Operations helpers.
 */
func updateHelper(e model.PersistedEntity, ac *model.AppConfig, coll string) (error) {
        sess := ac.CopyDBSession()
        defer sess.Close()
        c := getEntityCollection(ac, sess, coll)

        e.Update()
        err := c.UpdateId(e.GetId(), &e)

        if err != nil {
                return fmt.Errorf(model.GenerateMessage(model.ERROR_MODEL_UPDATE_FAILED, e), err)
        }
        return nil
}

func getHelper(id bson.ObjectId, e model.PersistedEntity, ac *model.AppConfig, coll string) (error) {
        sess := ac.CopyDBSession()
        defer sess.Close()
        c := getEntityCollection(ac, sess, coll)

        err := c.FindId(id).One(e)

        if err != nil {
                return fmt.Errorf(model.GenerateMessage(model.ERROR_MODEL_GET_FAILED, e), err)
        }
        return nil

}

func queryHelperSingle(query bson.M, e model.PersistedEntity, ac *model.AppConfig, coll string) (error) {
        sess := ac.CopyDBSession()
        defer sess.Close()
        c := getEntityCollection(ac, sess, coll)

        err := c.Find(query).One(e)

        if err != nil {
                return fmt.Errorf(model.GenerateMessage(model.ERROR_MODEL_QUERY_FAILED, e), query, err)
        }
        return nil
}

func createSort(rev bool, s string) string {
        if (rev) {
                s = "-" + s// reverse sort
        }
        return s
}
