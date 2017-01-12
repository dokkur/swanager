package entities
import(
    "github.com/da4nik/swanager/shared"
    "labix.org/v2/mgo/bson"
)

const collection string = "users"

type User struct {
    Id              bson.ObjectId `bson:"_id,omitempty"`
    Applications    []Application
    Token           string
}

func FindUser(criterias bson.M) (*User, error) {
    entity := User{}
    col := shared.GetMongoDB().C(collection)
    err := col.Find(criterias).One(&entity)
    return &entity, err
}

func CreateUser(entity *User) (*User, error) {
    col := shared.GetMongoDB().C(collection)
    entity.Id = bson.NewObjectId()
    err := col.Insert(entity)
    return entity, err
}

func RemoveUser(entity *User) error {
    col := shared.GetMongoDB().C(collection)
    err := col.Remove(bson.M{"_id": entity.Id})
    return err
}

func FindUserByToken(token string) (*User, error) {
    return FindUser(bson.M{"token": token})
}
