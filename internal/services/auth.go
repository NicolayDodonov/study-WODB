package services

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"study-WODB/internal/model"

	"golang.org/x/oauth2"
)

type AuthServices struct {
	//todo: дать подключение к СУБД PostgresSQL
}

func NewAuthServices() *AuthServices {
	return &AuthServices{}
}

// AuthUser проверяет актуальность входящих данных
func (s *AuthServices) AuthUser(user *model.AuthInfo) error {
	switch user.Type {
	case model.Normal:
		return nil
	case model.Google:
		return nil
	case model.Yandex:
		return nil
	default:
		return errors.New("invalid user type")
	}
}

func (s *AuthServices) ParseGoogleData(token *oauth2.Token) (*model.AuthInfo, error) {
	// запрашиваем данные о пользователе
	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return nil, err
	}
	// декодируем полученные данные о пользователе
	rawData, err := io.ReadAll(resp.Body) //userData
	if err != nil {
		return nil, err
	}
	var UserInfo model.AuthInfo
	err = json.Unmarshal(rawData, &UserInfo)
	if err != nil {
		return nil, err
	}
	UserInfo.Type = model.Google

	return &UserInfo, nil
}
