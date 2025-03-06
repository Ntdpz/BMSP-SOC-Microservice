package services

import (
	"bmsp-backend-service/utils"
	"errors"
)

func (s Services) Login(username, password string) (string, error) {

	// find user
	user, err := s.repo.FindUser(username)

	if err != nil {
		return "", err
	}

	// check password
	if !utils.CheckPasswordHash(password, user.Password) {
		return "", errors.New("invalid password")
	}

	// generate jwt
	token, err := utils.GenerateJWT(user.Username)
	if err != nil {
		return "", err
	}

	return token, nil
}
