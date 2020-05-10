package extraDB

import (
	"context"

	"github.com/ZeroTechh/VelocityCore/utils"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/ZeroTechh/UserExtraService/core/types"
)

// New returns a new extraDB handler struct
func New() *ExtraDB {
	e := ExtraDB{}
	e.init()
	return &e
}

// ExtraDB is used to handle user extra data
type ExtraDB struct {
	coll *mongo.Collection
}

// init initializes client and database
func (e *ExtraDB) init() {
	c := utils.CreateMongoDB(dbConfig.Str("address"), log)
	db := c.Database(dbConfig.Str("db"))
	e.coll = db.Collection(dbConfig.Str("collection"))
}

// Create adds new user extra data into db
func (e ExtraDB) Create(ctx context.Context, data types.Extra) (string, error) {
	if !valid(data) {
		return messages.Str("invalidUserData"), nil
	}
	_, err := e.coll.InsertOne(ctx, data)
	return "", errors.Wrap(err, "Error while inserting into db")
}

// Get returns user extra data
func (e ExtraDB) Get(ctx context.Context, userID string) (data types.Extra, err error) {
	err = e.coll.FindOne(ctx, types.Extra{UserID: userID}).Decode(&data)
	err = errors.Wrap(err, "Error while finding from db")
	return
}

// Update updates user extra data
func (e ExtraDB) Update(
	ctx context.Context, userID string, update types.Extra) (string, error) {

	if !updateValid(update) {
		return messages.Str("invalidUserData"), nil
	}

	_, err := e.coll.UpdateOne(
		ctx,
		types.Extra{UserID: userID},
		map[string]types.Extra{"$set": update},
	)

	return "", errors.Wrap(err, "Error while updating db")
}
