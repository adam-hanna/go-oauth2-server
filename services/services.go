package services

import (
	"reflect"

	"github.com/adam-hanna/go-oauth2-server/config"
	"github.com/adam-hanna/go-oauth2-server/health"
	"github.com/adam-hanna/go-oauth2-server/oauth"
	"github.com/adam-hanna/go-oauth2-server/session"
	"github.com/adam-hanna/go-oauth2-server/web"
	"github.com/gorilla/sessions"
	"github.com/jinzhu/gorm"
)

func init() {

}

var (
	HealthService  health.ServiceInterface
	OauthService   oauth.ServiceInterface
	WebService     web.ServiceInterface
	SessionService session.ServiceInterface
)

// UseHealthService sets the health service
func UseHealthService(h health.ServiceInterface) {
	HealthService = h
}

// UseOauthService sets the oAuth service
func UseOauthService(o oauth.ServiceInterface) {
	OauthService = o
}

// UseWebService sets the web service
func UseWebService(w web.ServiceInterface) {
	WebService = w
}

// UseSessionService sets the session service
func UseSessionService(s session.ServiceInterface) {
	SessionService = s
}

// InitServices starts up all services
func InitServices(cnf *config.Config, db *gorm.DB) error {
	if nil == reflect.TypeOf(HealthService) {
		HealthService = health.NewService(db)
	}

	if nil == reflect.TypeOf(OauthService) {
		OauthService = oauth.NewService(cnf, db)
	}

	if nil == reflect.TypeOf(SessionService) {
		// note: default session store is CookieStore
		SessionService = session.NewService(cnf, sessions.NewCookieStore([]byte(cnf.Session.Secret)))
	}

	if nil == reflect.TypeOf(WebService) {
		WebService = web.NewService(cnf, OauthService, SessionService)
	}

	return nil
}

// CloseServices closes any open services
func CloseServices() {
	HealthService.Close()
	OauthService.Close()
	WebService.Close()
	SessionService.Close()
}