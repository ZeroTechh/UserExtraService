package handler

import (
	"context"

	proto "github.com/ZeroTechh/VelocityCore/proto/UserExtraService"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"github.com/ZeroTechh/UserExtraService/core/extraDB"
	"github.com/ZeroTechh/UserExtraService/core/types"
)

func toProto(data types.Extra) *proto.Data {
	r := proto.Data{}
	copier.Copy(&r, &data)
	return &r
}

func toData(request *proto.Data) (data types.Extra) {
	copier.Copy(&data, &request)
	return
}

// New returns a new service handler
func New() *Handler {
	h := Handler{}
	h.init()
	return &h
}

// Handler handles all user extra service functions
type Handler struct {
	extra *extraDB.ExtraDB
}

// init initializes
func (h *Handler) init() {
	h.extra = extraDB.New()
}

// Add adds user extra data into db
func (h Handler) Add(ctx context.Context, request *proto.Data) (*proto.AddResponse, error) {
	data := toData(request)
	msg, err := h.extra.Create(ctx, data)
	err = errors.Wrap(err, "Error while creating user extra data")
	return &proto.AddResponse{Message: msg}, err
}

// Get returns user extra data
func (h Handler) Get(ctx context.Context, request *proto.GetRequest) (*proto.Data, error) {
	data, err := h.extra.Get(ctx, request.UserID)
	err = errors.Wrap(err, "Error while getting user extra data")
	response := toProto(data)
	return response, err
}

// Update updates user extra data
func (h Handler) Update(ctx context.Context, request *proto.UpdateRequest) (*proto.UpdateResponse, error) {
	update := toData(request.Update)
	msg, err := h.extra.Update(ctx, request.UserID, update)
	err = errors.Wrap(err, "Error while updating user extra data")
	return &proto.UpdateResponse{Message: msg}, err
}

// Validate validates user extra data
func (h Handler) Validate(ctx context.Context, request *proto.Data) (*proto.ValidateResponse, error) {
	valid := extraDB.Valid(toData(request))
	return &proto.ValidateResponse{Valid: valid}, nil
}
