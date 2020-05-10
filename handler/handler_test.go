package handler

import (
	"context"
	"testing"

	proto "github.com/ZeroTechh/VelocityCore/proto/UserExtraService"
	"github.com/stretchr/testify/assert"

	"github.com/ZeroTechh/UserExtraService/core/types"
	"github.com/ZeroTechh/UserExtraService/core/utils"
)

func TestServiceHandler(t *testing.T) {
	assert := assert.New(t)
	h := New()
	ctx := context.TODO()

	// Testing Add function
	data := utils.Mock()
	addResp, err := h.Add(ctx, toProto(data))
	assert.Zero(addResp.Message)
	assert.NoError(err)

	// Testing Get
	getResp, err := h.Get(ctx, &proto.GetRequest{UserID: data.UserID})
	assert.Equal(data, toData(getResp))
	assert.NoError(err)

	// Testing Update
	data2 := utils.Mock()
	update := types.Extra{FirstName: data2.FirstName}
	req := proto.UpdateRequest{UserID: data.UserID, Update: toProto(update)}
	resp, err := h.Update(ctx, &req)
	assert.NoError(err)
	assert.Zero(resp.Message)

	// Testing Validate
	data = utils.Mock()
	valid, err := h.Validate(ctx, toProto(data))
	assert.True(valid.Valid)
	assert.NoError(err)
}
