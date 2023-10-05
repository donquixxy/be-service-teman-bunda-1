package services

import (
	"errors"

	"github.com/sirupsen/logrus"
	"github.com/tensuqiuwulu/be-service-teman-bunda/exceptions"
	"github.com/tensuqiuwulu/be-service-teman-bunda/models/http/response"
	"github.com/tensuqiuwulu/be-service-teman-bunda/repository/mysql"
	"gorm.io/gorm"
)

type AppVersionService interface {
	FindVersionApp(requestID string, os int) response.FindSettingVerApp2
}

type appVersionService struct {
	db                   *gorm.DB
	l                    *logrus.Logger
	appVersionRepository mysql.AppVersionRepository
}

func NewAppVersionService(
	db *gorm.DB,
	l *logrus.Logger,
	appVersionRepository mysql.AppVersionRepository,
) AppVersionService {
	return &appVersionService{
		db:                   db,
		l:                    l,
		appVersionRepository: appVersionRepository,
	}
}

func (a *appVersionService) FindVersionApp(requestID string, os int) response.FindSettingVerApp2 {
	verApp := a.appVersionRepository.FindByValue(a.db, os)

	if len(verApp) == 0 {
		exceptions.PanicIfRecordNotFound(errors.New("no app version were found"), requestID, []string{"version not found"}, a.l)
	}
	response := response.ToFindSettingAppResponses(os, verApp[0].Ver, verApp[1].Ver)

	return response
}
