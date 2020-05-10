package extraDB

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/ZeroTechh/UserExtraService/core/types"
	"github.com/ZeroTechh/UserExtraService/core/utils"
)

func TestExtraDB(t *testing.T) {
	assert := assert.New(t)
	e := New()
	ctx := context.TODO()

	// Testing Create function
	data := utils.Mock()
	msg, err := e.Create(ctx, data)
	assert.Zero(msg)
	assert.NoError(err)

	// Testing that Create returns invalid data message for invalid data
	msg, err = e.Create(ctx, types.Extra{})
	assert.NoError(err)
	assert.NotZero(msg)

	// Testing Get
	d, err := e.Get(ctx, data.UserID)
	assert.NoError(err)
	assert.Equal(data, d)

	// Testing Update
	data2 := utils.Mock()
	update := types.Extra{FirstName: data2.FirstName}
	msg, err = e.Update(ctx, data.UserID, update)
	assert.Zero(msg)
	assert.NoError(err)

	d, err = e.Get(ctx, data.UserID)
	assert.NoError(err)
	assert.Equal(data2.FirstName, d.FirstName)

	// Testing Update returns message for invalid update
	update = types.Extra{UserID: "NN"}
	msg, err = e.Update(ctx, data.UserID, update)
	assert.NotZero(msg)
	assert.NoError(err)

	// Testing Create returns message for invalid age
	data = utils.Mock()
	data.BirthdayUTC = time.Now().Unix()
	msg, err = e.Create(ctx, data)
	assert.NotZero(msg)
	assert.NoError(err)
}
