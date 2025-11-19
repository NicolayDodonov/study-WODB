package handler

import (
	"fmt"
	"net/http"
	"os"
	"study-WODB/internal/config"
	"study-WODB/internal/logger"
	"study-WODB/internal/services"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type GoogleAuth struct {
	Config *oauth2.Config
	State  string
	log    *logger.Logger
	svc    *services.AuthServices
}

// New возвращает готовую конфигурацию
func NewGoogleAuth(cnf *config.Config, log *logger.Logger) *GoogleAuth {
	AuthConfig := &oauth2.Config{
		RedirectURL:  makeRedirectUrlGoogle(cnf),        //адрес переадресации назад
		ClientID:     os.Getenv("Google_Client_ID"),     //получить из переменных окружений
		ClientSecret: os.Getenv("Google_Client_Secret"), //получить из переменных окружений
		Scopes: []string{ // Список получаемых данных
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}
	return &GoogleAuth{
		Config: AuthConfig,
		State:  cnf.State,
		log:    log,
		svc:    services.NewAuthServices(),
	}
}

// GoogleCall формирует url и переадресует понему к google oAuth2 сервис.
func (auth *GoogleAuth) GoogleCall(w http.ResponseWriter, r *http.Request) {
	auth.log.Info("Google-Oauth2 call started")
	url := auth.Config.AuthCodeURL(auth.State)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

// GoogleBack принимает запрос от google oAuth2 сервиса и обрабатывает полученные данные.
func (auth *GoogleAuth) GoogleBack(w http.ResponseWriter, r *http.Request) {
	// проверяем правильность State токена
	if r.FormValue("state") != auth.State {
		auth.log.Error("Google-Oauth2 error: state mismatch")
		http.Error(w, "state mismatch", http.StatusBadRequest)
		return
	}
	code := r.FormValue("code")
	// запрашиваем токен
	token, err := auth.Config.Exchange(oauth2.NoContext, code) //token
	if err != nil {
		auth.log.Error("Oauth2 error:" + err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// получаем информацию о пользователе
	user, err := auth.svc.ParseGoogleData(token)
	if err != nil {
		auth.log.Error("Oauth2 error:" + err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = auth.svc.AuthUser(user); err != nil {
		auth.log.Error("Oauth2 error:" + err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	auth.log.Info(fmt.Sprintf("Google user %s logged", user.Email))
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

func makeRedirectUrlGoogle(cnf *config.Config) string {
	return fmt.Sprintf("http://%s:%d/auth/google-callback", cnf.HttpServer.Addr, cnf.HttpServer.Port)
}
