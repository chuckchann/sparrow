package serror

import (
	"fmt"
	"github.com/golang/protobuf/ptypes"
	"sparrow/internal/pkg/entity"
	"sparrow/internal/pkg/slog"

	//"github.com/golang/protobuf/proto"

	"github.com/golang/protobuf/proto"
	common_pb "sparrow/api/protobuf_spec/common"
)

type SError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (e *SError) String() string {
	return fmt.Sprintf("code:%d  msg:%s ", e.Code, e.Msg)
}

func (e *SError) PbResponse() *common_pb.Response {
	resp := &common_pb.Response{
		Code: int64(e.Code),
		Msg:  e.Msg,
	}
	return resp
}

//transfer serror to common response
func (e *SError) Response() *entity.Response {
	resp := &entity.Response{
		Code: e.Code,
		Msg:  e.Msg,
	}
	return resp
}

func (e *SError) ReplaceMsg(newMsg string) *SError {
	e.Msg = newMsg
	return e
}

func (e *SError) ReplaceCode(newCode int) *SError {
	e.Code = newCode
	return e

}

func (e *SError) ResponseWithData(m proto.Message) *common_pb.Response {
	any, err := ptypes.MarshalAny(m)
	if err != nil {
		slog.Warn("ResponseWithData: MarshalAny err ", err.Error(), "proto message:", m.String())
		return ERR_MARSHAL_ANY.Response()
	}
	resp := &common_pb.Response{
		Code: int64(e.Code),
		Msg:  e.Msg,
		Data: any,
	}
	return resp
}

func (e *SError) Error() string {
	return e.String()
}

func NewError(e error) *SError {
	if e == nil {
		return ERR_SYSTEM
	} else {
		return ERR_SYSTEM.ReplaceMsg(e.Error())
	}
}
