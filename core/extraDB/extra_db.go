package extraDB

import (
	"context"

	"github.com/ZeroTechh/VelocityCore/utils"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/ZeroTechh/UserExtraService/core/types"
)

// New returns a new extraDB handler struct
func New() *ExtraDB {
	extraDB := ExtraDB{}
	extraDB.init()
	return &extraDB
}

// ExtraDB is used to handle user extra data
type ExtraDB struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection
}

// init initializes client and database
func (extraDB *ExtraDB) init() {
	extraDB.client = utils.CreateMongoDB(dbConfig.Str("address"), log)
	extraDB.database = extraDB.client.Database(dbConfig.Str("db"))
	extraDB.collection = extraDB.database.Collection(dbConfig.Str("collection"))
}

// Create is used to add new data into database
func (extraDB ExtraDB) Create(data types.Extra) string {
	if !IsDataValid(data) {
		return messages.Str("invalidUserData")
	}

	extraDB.collection.InsertOne(context.TODO(), data)
	return ""
}

// Get is used to a users data
func (extraDB ExtraDB) Get(userID string) (data types.Extra) {
	filter := types.Extra{UserID: userID}
	extraDB.collection.FindOne(context.TODO(), filter).Decode(&data)
	return
}

// Update updates user's extraDB data
func (extraDB ExtraDB) Update(userID string, update types.Extra) string {
	if !IsUpdateValid(update) {
		return messages.Str("invalidUserData")
	}

	extraDB.collection.UpdateOne(
		context.TODO(),
		types.Extra{UserID: userID},
		map[string]types.Extra{"$set": update},
	)
	return ""
}
