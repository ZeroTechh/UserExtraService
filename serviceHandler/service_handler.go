package serviceHandler

import (
	"context"

	proto "github.com/ZeroTechh/VelocityCore/proto/UserExtraService"
	"github.com/jinzhu/copier"

	"github.com/ZeroTechh/UserExtraService/core/extraDB"
	"github.com/ZeroTechh/UserExtraService/core/types"
)

func dataToProto(data types.Extra) *proto.Data {
	request := proto.Data{}
	copier.Copy(&request, &data)
	return &request
}

func protoToData(request *proto.Data) (data types.Extra) {
	copier.Copy(&data, &request)
	return
}

// New returns a new service handler
func New() *Handler {
	handler := Handler{}
	handler.init()
	return &handler
}

// Handler is used to handle all user extra service functions
type Handler struct {
	extraDB *extraDB.ExtraDB
}

// Init is used to initialize
func (handler *Handler) init() {
	handler.extraDB = extraDB.New()
}

// Add is used to handle Add function
func (handler Handler) Add(ctx context.Context, request *proto.Data) (*proto.AddResponse, error) {
	data := protoToData(request)
	msg := handler.extraDB.Create(data)
	return &proto.AddResponse{Message: msg}, nil
}

// Get is used to handle Get function
func (handler Handler) Get(ctx context.Context, request *proto.GetRequest) (*proto.Data, error) {
	data := handler.extraDB.Get(request.UserID)
	response := dataToProto(data)
	return response, nil
}

// Update is used to handler Update function
func (handler Handler) Update(ctx context.Context, request *proto.UpdateRequest) (*proto.UpdateResponse, error) {
	update := protoToData(request.Update)
	msg := handler.extraDB.Update(request.UserID, update)
	return &proto.UpdateResponse{Message: msg}, nil
}

// Validate is used to handle Validate function
func (handler Handler) Validate(ctx context.Context, request *proto.Data) (*proto.ValidateResponse, error) {
	valid := extraDB.IsDataValid(protoToData(request))
	return &proto.ValidateResponse{Valid: valid}, nil
}
