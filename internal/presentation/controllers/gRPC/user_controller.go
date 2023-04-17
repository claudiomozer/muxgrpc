package grpc

import (
	muxgrpc "github.com/claudiomozer/muxgrpc/internal/presentation/controllers/gRPC/proto"
	controllers_protocols "github.com/claudiomozer/muxgrpc/internal/presentation/controllers/protocols"
)

type UserController struct {
	create_user           controllers_protocols.CreateUserUseCase
	find_user_by_document controllers_protocols.FindUserByDocumentUseCase
	server                muxgrpc.UserServiceServer
}

func NewUserController(
	create_user controllers_protocols.CreateUserUseCase,
	find_user_by_document controllers_protocols.FindUserByDocumentUseCase,
) *UserController {
	return &UserController{
		create_user:           create_user,
		find_user_by_document: find_user_by_document,
	}
}

func (controller *UserController) SetServer(server muxgrpc.UserServiceServer) {
	controller.server = server
}
