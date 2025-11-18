package auth

import (
	"io/ioutil"
	"net/http"
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
		RedirectURL:  cnf.AuthRedirectURL, //адрес переадресации назад
		ClientID:     "",                  //todo: получить из переменных окружений
		ClientSecret: "",                  //todo: получить из переменных окружений
		Scopes:       []string{},
		Endpoint:     google.Endpoint,
	}
	return &Auth{
		Config: AuthConfig,
		State:  cnf.State,
	}
}

func (auth *Auth) GoogleCall(w http.ResponseWriter, r *http.Request) {
	url := auth.Config.AuthCodeURL("state")
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
