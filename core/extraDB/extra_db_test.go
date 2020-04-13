package extraDB

import (
	"testing"
	"time"

	"github.com/ZeroTechh/UserService/core/types"
	"github.com/stretchr/testify/assert"

	"github.com/ZeroTechh/UserService/core/utils"
)

func TestExtraDB(t *testing.T) {
	assert := assert.New(t)
	extraDB := New()

	// Testing Create function
	_, mockData := utils.GetMockUserData()
	mockData.UserID = mockData.FirstName
	msg := extraDB.Create(mockData)
	assert.Zero(msg)

	// Testing that Create returns invalid data message for invalid data
	assert.NotZero(extraDB.Create(types.Extra{}))

	// Testing Get
	returnedData := extraDB.Get(mockData.UserID)
	assert.Equal(mockData, returnedData)

	// Testing Update
	_, mockData2 := utils.GetMockUserData()
	update := types.Extra{FirstName: mockData2.FirstName}
	msg = extraDB.Update(mockData.UserID, update)
	assert.Zero(msg)

	returnedData = extraDB.Get(mockData.UserID)
	assert.Equal(mockData2.FirstName, returnedData.FirstName)

	// Testing Update returns message for invalid update
	update = types.Extra{UserID: "NN"}
	msg = extraDB.Update(mockData.UserID, update)
	assert.NotZero(msg)

	// Testing Create returns message for invalid age
	_, mockData = utils.GetMockUserData()
	mockData.UserID = mockData.FirstName
	mockData.BirthdayUTC = time.Now().Unix()
	msg = extraDB.Create(mockData)
	assert.NotZero(msg)
}
