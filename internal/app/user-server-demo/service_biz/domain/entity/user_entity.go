package entity

import (
	user_pb "sparrow/api/protobuf_spec/user"
	"time"
)

type User struct {
	ID         uint64     `gorm:"primary_key;auto_increment" json:"id"`
	FirstName  string     `gorm:"size:100;not null;" json:"first_name"`
	LastName   string     `gorm:"size:100;not null;" json:"last_name"`
	Mobile     string     `gorm:"size:100;not null;" json:"mobile"`
	Email      string     `gorm:"size:100;not null;unique" json:"email"`
	Job        string     `gorm:"size:100;not null;unique" json:"job"`
	Password   string     `gorm:"size:100;not null;" json:"password"`
	Sex        int        `gorm:"sex" json:"sex"`
	UnionId    string     `gorm:"union_id" json:"unionId"`
	HeadImgUrl string     `gorm:"head_img_url" json:"headImgUrl"`
	CreatedAt  time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at,omitempty"`
}

//usually for list item
type ApiUser struct {
	FirstName string `gorm:"size:100;not null;" json:"first_name"`
	Email     string `gorm:"size:100;not null;unique" json:"email"`
}

func (*User) TableName() string {
	return "t_user"
}

func (u *User) ToPbUser() *user_pb.UserInfo {
	return &user_pb.UserInfo{
		Id:     int64(u.ID),
		Name:   u.FirstName,
		Email:  u.Email,
		Mobile: u.Mobile,
		Job:    u.Job,
	}
}

type LoginReq struct {
	Code string `json:"code"`
}
