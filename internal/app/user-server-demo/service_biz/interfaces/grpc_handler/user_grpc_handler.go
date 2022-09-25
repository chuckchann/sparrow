package grpc_handler

import (
	"context"
	common_pb "sparrow/api/protobuf_spec/common"
	user_pb "sparrow/api/protobuf_spec/user"
	"sparrow/internal/app/user-server-demo/service_biz/application"
	"sparrow/internal/pkg/serror"
)

type UserServer struct {
	ua application.UserAppInterface
}

func NewUserServer(ua application.UserAppInterface) *UserServer {
	return &UserServer{
		ua: ua,
	}
}

func (u *UserServer) GetUserInfo(ctx context.Context, req *user_pb.GetUserInfoRequest) (*common_pb.Response, error) {

}

//OAuth2.0 Grant Login
func (u *UserServer) Login(ctx context.Context, req *user_pb.GetUserInfoRequest) (*common_pb.Response, error) {

}
