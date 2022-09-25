package rpc_service

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes"
	common_pb "sparrow/api/protobuf_spec/common"
	order_pb "sparrow/api/protobuf_spec/order"
	user_pb "sparrow/api/protobuf_spec/user"
	"sparrow/internal/app/order-server-demo/grpc_client"
	"sparrow/internal/pkg/serror"
)

func NewOrderService() *OrderService {
	return &OrderService{}
}

type OrderService struct {
}

func (*OrderService) GetOrderInfo(ctx context.Context, orderReq *order_pb.GetOrderInfoReq) (*common_pb.Response, error) {
	//call user-server-demo first
	userCli, err := grpc_client.GetGRPCClient(context.Background(), "user-server-demo.default")
	if err != nil {
		fmt.Println(err.Error())
		return serror.ERR_UNFOUND_CLI.Response(), nil
	}
	userResp, err := userCli.(user_pb.UserClient).GetUserInfo(ctx, &user_pb.GetUserInfoRequest{Id: 1})
	if err != nil {
		return serror.ERR_CALL.Response(), nil
	}
	if userResp.Code != 0 {
		return serror.ERR_BIZ.Response(), nil
	}
	user := &user_pb.UserInfo{}
	err = ptypes.UnmarshalAny(userResp.Data, user)
	if err != nil {
		return serror.ERR_MARSHAL_ANY.Response(), nil
	}

	orderInfo := &order_pb.OrderInfo{
		Price:       110,
		Address:     "dream house",
		ProductName: "wash",
		User:        user,
	}

	return serror.SUCCESS.ResponseWithData(orderInfo), nil

}
