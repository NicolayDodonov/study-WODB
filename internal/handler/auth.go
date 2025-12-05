package handler

import (
	"net/http"
	"study-WODB/internal/logger"
	"study-WODB/internal/model"
	"study-WODB/internal/services"
)

type Auth struct {
	log *logger.Logger
	svc *services.AuthServices
}

func NewAuth(log *logger.Logger, svc *services.AuthServices) *Auth {
	return &Auth{
		log: log,
		svc: svc,
	}
}

func (s *Auth) Login(w http.ResponseWriter, r *http.Request) {
	// todo

}

// Внутренняя функция аутентификации пользователя в системе.
// Проверяет если ли пользователь с таким логином в системе.
// Если его нет - регистрирует его с таким логином и паролем.
// Если есть - проверяет его пароль.
func (a *Auth) login(userData *model.AuthInfo) error {
	a.log.Debug("- login")
	// Проверяем существование пользователя в системе
	isExist, err := a.svc.CheckUser(userData.Email)
	if err != nil {
		return err
	}
	if isExist {
		// Пользователь есть, проверяем его email и пароль
		err = a.svc.CheckUserPassword(userData.Email, userData.Password)
		if err != nil {
			return err
		}
		return nil
	} else {
		err = a.svc.AddUser(userData)
		if err != nil {
			return err
		}
		return nil
	}
}
