package entity

import (
	"time"

	"gopkg.in/guregu/null.v4"
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
	VerificationDueDate     time.Time       `gorm:"column:verification_due_date;"`
	VerificationDate        null.Time       `gorm:"column:verification_date;"`
	NotVerification         int             `gorm:"column:not_verification;"`
	OtpCode                 string          `gorm:"column:otp_code;"`
	OtpCodeExpiredDueDate   null.Time       `gorm:"column:otp_code_expired_due_date;"`
	OtpLimitPhone           int             `gorm:"column:otp_limit_phone;"`
	OtpLimitResetDate       null.Time       `gorm:"column:otp_limit_reset_date;"`
	CreatedDate             time.Time       `gorm:"column:created_at;"`
	TokenDevice             string          `gorm:"column:token_device;"`
	IsDelete                int             `gorm:"column:is_delete;"`
	// New Timegap
	IsTimegap             	int          	`gorm:"column:is_timegap;"`
	TimegapData             string          `gorm:"column:timegap_data;"`
}

func (User) TableName() string {
	return "users"
}
