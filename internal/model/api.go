package model

type AuthInfo struct {
	Email    string `yaml:"email"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
	Type     string `yaml:"type"`
}
