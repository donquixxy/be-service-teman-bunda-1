package entity

import (
	"time"
)

type User struct {
	Id                      string          `gorm:"primaryKey;column:id;"`
	IdFamilyMembers         string          `gorm:"column:id_family_members;"`
	FamilyMembers           FamilyMembers   `gorm:"foreignKey:IdFamilyMembers"`
	BalancePoint            BalancePoint    `gorm:"foreignKey:IdUser"`
	Username                string          `gorm:"column:username;"`
	Password                string          `gorm:"column:password;"`
	RefreshToken            string          `gorm:"column:refresh_token;"`
	IsActive                int             `gorm:"column:is_active;"`
	ReferalCode             string          `gorm:"column:referal_code;"`
	RegistrationReferalCode string          `gorm:"column:registration_referal_code;"`
	IdRole                  string          `gorm:"column:id_role;"`
	IdLevelMember           int             `gorm:"column:id_level_member;"`
	PasswordResetCode       string          `gorm:"column:password_reset_code;"`
	UserLevelMember         UserLevelMember `gorm:"foreignKey:IdLevelMember"`
	CreatedDate             time.Time       `gorm:"column:created_at;"`
}

func (User) TableName() string {
	return "users"
}
