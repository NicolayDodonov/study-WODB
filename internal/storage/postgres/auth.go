package postgres

import "study-WODB/internal/model"

func (s *Storage) CheckUserByEmail(email string) (bool, error) {
	
	return false, nil
}

func (s *Storage) CheckUserByEmailAndPassword(email, password string) error {
	return nil
}

func (s *Storage) AddUser(userdata *model.AuthInfo) error {
	return nil
}
