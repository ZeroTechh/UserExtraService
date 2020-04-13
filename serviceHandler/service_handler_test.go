package serviceHandler

import (
	"context"
	"fmt"
	"testing"

	proto "github.com/ZeroTechh/VelocityCore/proto/UserExtraService"
	"github.com/stretchr/testify/assert"

	"github.com/ZeroTechh/UserExtraService/core/types"
	"github.com/ZeroTechh/UserExtraService/core/utils"
)

func TestServiceHandler(t *testing.T) {
	assert := assert.New(t)
	handler := New()
	ctx := context.TODO()

	// Testing Add function
	mockData := utils.MockData()
	fmt.Println(mockData)
	addResponse, err := handler.Add(ctx, dataToProto(mockData))
	assert.Zero(addResponse.Message)
	assert.NoError(err)

	// Testing Get
	getResponse, err := handler.Get(
		ctx, &proto.GetRequest{UserID: mockData.UserID})
	assert.Equal(mockData, protoToData(getResponse))
	assert.NoError(err)

	// Testing Update
	mockData2 := utils.MockData()
	update := types.Extra{FirstName: mockData2.FirstName}
	updateRequest := proto.UpdateRequest{
		UserID: mockData.UserID, Update: dataToProto(update),
	}
	updateResponse, err := handler.Update(ctx, &updateRequest)
	assert.NoError(err)
	assert.Zero(updateResponse.Message)
}
