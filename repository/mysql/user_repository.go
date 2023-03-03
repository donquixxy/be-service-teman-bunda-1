package mysql

import (
	"github.com/tensuqiuwulu/be-service-teman-bunda/config"
	"github.com/tensuqiuwulu/be-service-teman-bunda/models/entity"
	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	CreateUser(DB *gorm.DB, user entity.User) (entity.User, error)
	UpdateUser(DB *gorm.DB, idUser string, user entity.User) (entity.User, error)
	UpdateStatusActiveUser(DB *gorm.DB, idUser string, user entity.User) (entity.User, error)
	UpdatePasswordResetCodeUser(DB *gorm.DB, idUser string, user entity.User) (entity.User, error)
	UpdateOtpCodeUser(DB *gorm.DB, idUser string, user entity.User) error
	UpdateUserPassword(DB *gorm.DB, idUser string, user entity.User) (entity.User, error)
	UpdateUserTokenDevice(DB *gorm.DB, idUser string, user entity.User) error
	DeleteAccount(DB *gorm.DB, idUser string, user entity.User) error
	FindUserByUsername(DB *gorm.DB, username string) (entity.User, error)
	FindUserByEmail(DB *gorm.DB, email string) (entity.User, error)
	FindUserByPhone(DB *gorm.DB, phone string) (entity.User, error)
	FindUserByReferal(DB *gorm.DB, referalCode string) (entity.User, error)
	FindUserById(DB *gorm.DB, id string) (entity.User, error)
	CountUserByRegistrationReferal(DB *gorm.DB, referal string) (userCount int, err error)
	SaveUserRefreshToken(DB *gorm.DB, id string, refreshToken string) (int64, error)
	FindUserByUsernameAndRefreshToken(DB *gorm.DB, username string, refresh_token string) (entity.User, error)
	FindUserByReferalCode(DB *gorm.DB, referalCode string) (entity.User, error)
	// TimegapApi
	CreateUserTimeGap(DB *gorm.DB, user entity.User) (entity.User, error)
	UpdateUserTimeGap(DB *gorm.DB, idUser string, user entity.User) (entity.User, error)
}

type UserRepositoryImplementation struct {
	configurationDatabase *config.Database
}

// =================================================== TimegapAPI
func (repository *UserRepositoryImplementation) CreateUserTimeGap(DB *gorm.DB, user entity.User) (entity.User, error) {
	results := DB.Create(user)
	return user, results.Error
}
func (repository *UserRepositoryImplementation) UpdateUserTimeGap(DB *gorm.DB, idUser string, user entity.User) (entity.User, error) {
	result := DB.
		Model(entity.User{}).
		Where("id = ?", idUser).
		Updates(entity.User{
			TimegapData: user.TimegapData,
		})
	return user, result.Error
}
// ======================================================= NORMAL API
func NewUserRepository(configDatabase *config.Database) UserRepositoryInterface {
	return &UserRepositoryImplementation{
		configurationDatabase: configDatabase,
	}
}

func (repository *UserRepositoryImplementation) UpdateUserTokenDevice(DB *gorm.DB, idUser string, user entity.User) error {
	result := DB.
		Model(entity.User{}).
		Where("id = ?", idUser).
		Updates(entity.User{
			TokenDevice: user.TokenDevice,
		})
	return result.Error
}

func (repository *UserRepositoryImplementation) UpdateUser(DB *gorm.DB, idUser string, user entity.User) (entity.User, error) {
	result := DB.
		Model(entity.User{}).
		Where("id = ?", idUser).
		Updates(entity.User{
			Username: user.Username,
			Password: user.Password,
		})
	return user, result.Error
}

func (repository *UserRepositoryImplementation) DeleteAccount(DB *gorm.DB, idUser string, user entity.User) error {
	result := DB.
		Model(entity.User{}).
		Where("id = ?", idUser).
		Updates(entity.User{
			IsDelete: user.IsDelete,
		})
	return result.Error
}

func (repository *UserRepositoryImplementation) UpdateUserPassword(DB *gorm.DB, idUser string, user entity.User) (entity.User, error) {
	result := DB.
		Model(entity.User{}).
		Where("id = ?", idUser).
		Updates(entity.User{
			Password:          user.Password,
			PasswordResetCode: user.PasswordResetCode,
		})
	return user, result.Error
}

func (repository *UserRepositoryImplementation) UpdateStatusActiveUser(DB *gorm.DB, idUser string, user entity.User) (entity.User, error) {
	result := DB.
		Model(entity.User{}).
		Where("id = ?", idUser).
		Updates(entity.User{
			IsActive:          user.IsActive,
			OtpCode:           user.OtpCode,
			PasswordResetCode: user.PasswordResetCode,
			VerificationDate:  user.VerificationDate,
		})
	return user, result.Error
}

func (repository *UserRepositoryImplementation) UpdatePasswordResetCodeUser(DB *gorm.DB, idUser string, user entity.User) (entity.User, error) {
	result := DB.
		Model(entity.User{}).
		Where("id = ?", idUser).
		Updates(entity.User{
			PasswordResetCode: user.PasswordResetCode,
		})
	return user, result.Error
}

func (repository *UserRepositoryImplementation) UpdateOtpCodeUser(DB *gorm.DB, idUser string, user entity.User) error {
	updateUserOtp := make(map[string]interface{})
	updateUserOtp["otp_code"] = user.OtpCode
	updateUserOtp["otp_code_expired_due_date"] = user.OtpCodeExpiredDueDate
	updateUserOtp["otp_limit_phone"] = user.OtpLimitPhone
	updateUserOtp["otp_limit_reset_date"] = user.OtpLimitResetDate

	result := DB.
		Model(entity.User{}).
		Where("id = ?", idUser).
		Updates(&updateUserOtp)
	return result.Error
}

func (repository *UserRepositoryImplementation) CreateUser(DB *gorm.DB, user entity.User) (entity.User, error) {
	results := DB.Create(user)
	return user, results.Error
}


func (repository *UserRepositoryImplementation) FindUserByUsername(DB *gorm.DB, username string) (entity.User, error) {
	var user entity.User
	results := DB.Where("users.username = ?", username).Where("users.not_verification = ?", 0).Joins("FamilyMembers").First(&user)
	return user, results.Error
}

func (repository *UserRepositoryImplementation) FindUserByReferal(DB *gorm.DB, referalCode string) (entity.User, error) {
	var user entity.User
	results := DB.Where("referal_code = ?", referalCode).Where("users.not_verification = ?", 0).First(&user)
	return user, results.Error
}

func (repository *UserRepositoryImplementation) FindUserByEmail(DB *gorm.DB, email string) (entity.User, error) {
	var user entity.User
	results := DB.Joins("FamilyMembers").Where("users.not_verification = ?", 0).Where("users.is_delete = ?", 0).Find(&user, "FamilyMembers.email = ?", email)
	return user, results.Error
}

func (repository *UserRepositoryImplementation) FindUserByPhone(DB *gorm.DB, phone string) (entity.User, error) {
	var user entity.User
	results := DB.Joins("FamilyMembers").Where("users.not_verification = ?", 0).Where("users.is_delete = ?", 0).Find(&user, "FamilyMembers.phone = ?", phone)
	return user, results.Error
}

func (repository *UserRepositoryImplementation) FindUserById(DB *gorm.DB, id string) (entity.User, error) {
	var user entity.User
	results := DB.Where("users.id = ?", id).Where("users.not_verification = ?", 0).
		Joins("FamilyMembers").
		Joins("BalancePoint").
		Joins("UserLevelMember").
		First(&user)
	return user, results.Error
}

func (repository *UserRepositoryImplementation) FindUserByReferalCode(DB *gorm.DB, referalCode string) (entity.User, error) {
	var user entity.User
	results := DB.Where("users.referal_code = ?", referalCode).
		Joins("FamilyMembers").
		Joins("BalancePoint").
		Joins("UserLevelMember").
		First(&user)
	return user, results.Error
}

func (repository *UserRepositoryImplementation) CountUserByRegistrationReferal(DB *gorm.DB, referalCode string) (countUser int, err error) {
	var user []entity.User
	results := DB.Model(&entity.User{}).Where("registration_referal_code = ?", referalCode).Find(&user)
	return int(results.RowsAffected), results.Error
}

func (repository *UserRepositoryImplementation) FindUserByUsernameAndRefreshToken(DB *gorm.DB, username string, refresh_token string) (entity.User, error) {
	var user entity.User
	results := DB.Where("username = ?", username).Where("refresh_token = ?", refresh_token).First(&user)
	return user, results.Error
}

func (repository *UserRepositoryImplementation) SaveUserRefreshToken(DB *gorm.DB, id string, refreshToken string) (int64, error) {
	results := DB.Exec("UPDATE `users` SET refresh_token = ? WHERE id = ?", refreshToken, id)
	return results.RowsAffected, results.Error
}
