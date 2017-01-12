package shared

import (
    "github.com/da4nik/swanager/config"
    "gopkg.in/mgo.v2"
)

func GetMongoSession() *mgo.Session {
    session, err := mgo.Dial(config.MongoURL)

    if err != nil {
        panic(err)
    }
    return session
}

func GetMongoDB() *mgo.Database {
    session := GetMongoSession()
    return session.DB(config.MongoDatabase)
}
