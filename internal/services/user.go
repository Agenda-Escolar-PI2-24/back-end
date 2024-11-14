package services

import (
	"agenda-escolar/internal/domain"
	"agenda-escolar/internal/storage/repository"
	"crypto/md5"
	"encoding/hex"
	"errors"
)

var userRepository repository.UserRepository

type UserService struct {
}

func (us *UserService) Register(user domain.User) error {
	exists, err := us.Exists(user.Username)
	if exists || err != nil {
		errMsg := "user already exists"
		if err != nil {
			errMsg = err.Error()
		}
		return errors.New(errMsg)
	}

	user.Password = encriptToMd5(user.Password)
	return userRepository.Register(user)
}

func (*UserService) Login(user *domain.User) (domain.User, error) {
	user.Password = encriptToMd5(user.Password)

	return userRepository.Auth(user)
}

func (*UserService) Exists(username string) (bool, error) {
	return userRepository.Exists(username)
}

// return str encrypted using md5 algorithm
func encriptToMd5(str string) string {
	md5ID := md5.Sum([]byte(str))
	hash := hex.EncodeToString(md5ID[:])

	return hash
}
