package handler

import (
	"fmt"
	"net/http"
	"os"
	"study-WODB/internal/config"
	"study-WODB/internal/logger"
	"study-WODB/internal/services"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/yandex"
)

type YandexAuth struct {
	Config *oauth2.Config
	State  string
	log    *logger.Logger
	svc    *services.AuthServices
}

func NewYandexAuth(cnf *config.Config, log *logger.Logger) *YandexAuth {
	AuthConfig := &oauth2.Config{
		RedirectURL:  makeRedirectUrlYandex(cnf),
		ClientID:     os.Getenv("Yandex_Client_ID"), //получить из переменных окружений
		ClientSecret: os.Getenv("Yandex_Client_Secret"),
		Endpoint:     yandex.Endpoint,
	}
	return &YandexAuth{
		Config: AuthConfig,
		State:  cnf.State,
		log:    log,
		svc:    services.NewAuthServices(),
	}
}

// YandexCall формирует url и переадресует понему к google oAuth2 сервис.
func (auth *YandexAuth) YandexCall(w http.ResponseWriter, r *http.Request) {
	auth.log.Info("Yandex-Oauth2 call started")
	url := auth.Config.AuthCodeURL(auth.State)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

// YandexBack принимает запрос от google oAuth2 сервиса и обрабатывает полученные данные.
func (auth *YandexAuth) YandexBack(w http.ResponseWriter, r *http.Request) {

}

func makeRedirectUrlYandex(cnf *config.Config) string {
	return fmt.Sprintf("http://%s:%d/auth/ya-callback", cnf.HttpServer.Addr, cnf.HttpServer.Port)
}
