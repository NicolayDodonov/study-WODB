package handler

import (
	"fmt"
	"net/http"
	"os"
	"study-WODB/internal/config"
	"study-WODB/internal/services"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type GoogleAuth struct {
	Config *oauth2.Config
	State  string
	svc    *services.AuthServices
}

// New возвращает готовую конфигурацию
func NewGoogleAuth(cnf *config.Config) *GoogleAuth {
	AuthConfig := &oauth2.Config{
		RedirectURL:  makeRedirectUrl(cnf),       //адрес переадресации назад
		ClientID:     os.Getenv("Client_ID"),     //получить из переменных окружений
		ClientSecret: os.Getenv("Client_Secret"), //получить из переменных окружений
		Scopes: []string{ // Список получаемых данных
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}
	return &GoogleAuth{
		Config: AuthConfig,
		State:  cnf.State,
		svc:    services.NewAuthServices(),
	}
}

// GoogleCall формирует url и переадресует понему к google oAuth2 сервис.
func (auth *GoogleAuth) GoogleCall(w http.ResponseWriter, r *http.Request) {
	url := auth.Config.AuthCodeURL(auth.State)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

// GoogleBack принимает запрос от google oAuth2 сервиса и обрабатывает полученные данные.
func (auth *GoogleAuth) GoogleBack(w http.ResponseWriter, r *http.Request) {
	// проверяем правильность State токена
	if r.FormValue("state") != auth.State {
		http.Error(w, "state did not match", http.StatusBadRequest)
		return
	}
	code := r.FormValue("code")
	// запрашиваем токен
	token, err := auth.Config.Exchange(oauth2.NoContext, code) //token
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// получаем информацию о пользователе
	user, err := auth.svc.ParseGoogleData(token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = auth.svc.AuthUser(user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

func makeRedirectUrl(cnf *config.Config) string {
	return fmt.Sprintf("http://%s:%d/auth/google-callback", cnf.HttpServer.Addr, cnf.HttpServer.Port)
}
