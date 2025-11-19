package model

type AuthInfo struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"Login"`
	Type     string `json:"type"`
}

type GoogleInfo struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type YandexUserInfo struct {
	Name  string `json:"login"`
	Email string `json:"default_email"`
}
