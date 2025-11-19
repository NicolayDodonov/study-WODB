package auth

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"study-WODB/internal/config"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type Auth struct {
	Config *oauth2.Config
	State  string
}

// New возвращает готовую конфигурацию
func New(cnf *config.Config) *Auth {
	AuthConfig := &oauth2.Config{
		RedirectURL:  makeRedirectUrl(cnf),       //адрес переадресации назад
		ClientID:     os.Getenv("Client_ID"),     //todo: получить из переменных окружений
		ClientSecret: os.Getenv("Client_Secret"), //todo: получить из переменных окружений
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}
	return &Auth{
		Config: AuthConfig,
		State:  cnf.State,
	}
}

func (auth *Auth) GoogleCall(w http.ResponseWriter, r *http.Request) {
	url := auth.Config.AuthCodeURL(auth.State)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func (auth *Auth) GoogleBack(w http.ResponseWriter, r *http.Request) {
	if r.FormValue("state") != auth.State {
		http.Error(w, "state did not match", http.StatusBadRequest)
	}
	code := r.FormValue("code")

	_, err := auth.Config.Exchange(oauth2.NoContext, code) //token
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	_, err = ioutil.ReadAll(resp.Body) //userData
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func makeRedirectUrl(cnf *config.Config) string {
	return fmt.Sprintf("http://%s:%d/auth/google-callback", cnf.HttpServer.Addr, cnf.HttpServer.Port)
}
