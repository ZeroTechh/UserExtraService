package extraDB

import (
	"reflect"

	"github.com/ZeroTechh/blaze"
	"github.com/ZeroTechh/sentinal/v2"

	"go.uber.org/zap"
)

var schemaPaths = []string{"schemas", "../core/extraDB/schemas", "core/extraDB/schemas"}

const schemaFile = "schema.yaml"

// Valid checks if data is valid
func Valid(data interface{}) bool {
	funcLog := blaze.NewFuncLog("Valid", log, zap.Any("Data", data))
	funcLog.Started()
	valid, msg, _ := sentinal.ValidateWithYAML(
		data,
		schemaFile,
		schemaPaths,
		customFuncs,
	)
	funcLog.Completed(zap.Any("Message", msg))
	return valid
}

// updateValid checks if the update is valid
func updateValid(update interface{}) bool {
	funcLog := blaze.NewFuncLog("UpdateValid", log, zap.Any("Update", update))
	funcLog.Started()
	valid, msg, _ := sentinal.ValidateFieldsWithYAML(
		update,
		schemaFile,
		schemaPaths,
		customFuncs,
	)

	// checking if the update data is not trying to update UserID
	if reflect.ValueOf(update).FieldByName("UserID").String() != "" {
		valid = false
		msg = map[string][]string{
			"UserID": {"Tring to update UserID"},
		}
	}

	funcLog.Completed(zap.Any("Message", msg))
	return valid
}
