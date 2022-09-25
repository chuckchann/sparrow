package persistence

import (
	"errors"
	"github.com/jinzhu/gorm"
	"sparrow/internal/app/user-server-demo/service_biz/domain/entity"
)

type UserRepoImpl struct {
	db *gorm.DB
}

func NewUserRepoImpl(db *gorm.DB) *UserRepoImpl {
	return &UserRepoImpl{db}
}

var _ UserRepo = &UserRepoImpl{}

func (r *UserRepoImpl) SaveUser(user *entity.User) (*entity.User, error) {
	err := r.db.Debug().Create(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepoImpl) GetUser(id uint64) (*entity.User, error) {
	var user entity.User
	err := r.db.Debug().Where("id = ?", id).Take(&user).Error
	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("user not found")
	}
	return &user, nil
}

func (r *UserRepoImpl) GetUsers() ([]entity.User, error) {
	var users []entity.User
	err := r.db.Debug().Find(&users).Error
	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("user not found")
	}
	return users, nil
}

func (r *UserRepoImpl) GetUserByUnionId(unionId string) (*entity.User, error) {
	var user entity.User
	err := r.db.Debug().Where("union_id = ?", unionId).Take(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepoImpl) UpdateUser(u *entity.User) error {

}
