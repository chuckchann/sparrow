package persistence

import "sparrow/internal/app/user-server-demo/service_biz/domain/entity"

//will realize in infrastructure layer
type UserRepo interface {
	SaveUser(*entity.User) (*entity.User, error)
	GetUser(uint64) (*entity.User, error)
	GetUsers() ([]entity.User, error)
	GetUserByUnionId(string) (*entity.User, error)
	UpdateUser(*entity.User) error
}
