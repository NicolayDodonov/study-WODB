package services

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"study-WODB/internal/model"
	"study-WODB/internal/storage/postgres"

	"golang.org/x/oauth2"
)

type AuthServices struct {
	//todo: дать подключение к СУБД PostgresSQL
	storage *postgres.Storage
}

func NewAuthServices(storage *postgres.Storage) *AuthServices {
	return &AuthServices{storage: storage}
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
	var UserInfo model.GoogleInfo
	err = json.Unmarshal(rawData, &UserInfo)
	if err != nil {
		return nil, err
	}

	// Получаем формат стандартной информации аутентификации
	User := model.AuthInfo{
		Type:     model.Google,
		Email:    UserInfo.Email,
		Name:     UserInfo.Name,
		Password: "",
	}
	return &User, nil
}

func (s *AuthServices) ParseYandexData(token *oauth2.Token) (*model.AuthInfo, error) {
	resp, err := http.Get("https://login.yandex.ru/info?&oauth_token=" + token.AccessToken)
	if err != nil {
		return nil, err
	}
	// декодируем полученные данные о пользователе
	rawData, err := io.ReadAll(resp.Body) //userData
	if err != nil {
		return nil, err
	}
	var UserInfo model.YandexUserInfo
	err = json.Unmarshal(rawData, &UserInfo)
	if err != nil {
		return nil, err
	}

	// Получаем формат стандартной информации аутентификации
	User := model.AuthInfo{
		Type:     model.Yandex,
		Email:    UserInfo.Email,
		Name:     UserInfo.Name,
		Password: "",
	}
	return &User, nil
}

func (s *AuthServices) CheckUser(email string) (bool, error) {
	if email == "" {
		return false, errors.New("email is empty")
	}
	IsExist, err := s.storage.CheckUserByEmail(email)
	return IsExist, err
}

func (s *AuthServices) CheckUserPassword(email, password string) error {
	if email == "" {
		return errors.New("email is empty")
	}
	err := s.storage.CheckUserByEmailAndPassword(email, password)
	if err != nil {
		return err
	}
	return nil
}

func (s *AuthServices) AddUser(userdata *model.AuthInfo) error {
	if userdata.Email == "" {
		return errors.New("email is empty")
	}
	err := s.storage.AddUser(userdata)
	if err != nil {
		return err
	}
	return nil
}
