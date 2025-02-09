package entity

import "time"

type User struct {
	ID         int64     `gorm:"primaryKey;column:id;type:int(10) unsigned;not null"`
	Username   string    `gorm:"unique;column:user_name;type:varchar(128);not null;default:''"`
	Password   string    `gorm:"column:password;type:varchar(255);not null;"`
	Sex        int8      `gorm:"column:sex;type:tinyint unsigned;not null;default:0"`
	Age        int8      `gorm:"column:age;type:tinyint unsigned;not null;default:0"`
	Phone      string    `gorm:"column:phone;type:varchar(128);not null;default:''"`
	IsNumber   int8      `gorm:"column:is_number;type:tinyint unsigned;not null;default:0"`
	MemberEnd  time.Time `gorm:"column:member_end;type:datetime;not null;default:CURRENT_TIMESTAMP"` // 会员到期时间
	UserStatus int8      `gorm:"column:user_status;type:tinyint unsigned;not null;default:0"`
	Email      string    `gorm:"column:email;type:varchar(128);not null;default:''"`
	Avatar     string    `gorm:"column:avatar;type:varchar(255);not null;default:''"`
	CreatedAt  time.Time `gorm:"column:created_at;type:datetime;not null;default:CURRENT_TIMESTAMP"` // 创建时间
	UpdatedAt  time.Time `gorm:"column:updated_at;type:datetime;not null;default:CURRENT_TIMESTAMP"` // 修改时间
}

type RegisterUserReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginUserReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ChangePwd struct {
	OldPassword string `json:"oldPassword" binding:"required"`
	NewPassword string `json:"newPassword" binding:"required"`
}

type UpdateUserReq struct {
	Username string `json:"username"`
	Sex      int8   `json:"sex"`
	Age      int8   `json:"age"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
}
