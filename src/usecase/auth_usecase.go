package usecase

import (
	"errors"
	"project/src/helper"
	"project/src/helper/jwt"
	"project/src/repository"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

// AuthUsecase ...
type AuthUsecase struct {
	Helper   helper.Helper
	UserRepo repository.UserRepoInterface
}

// Login ...
func (t *AuthUsecase) Login(email string, password string) (string, error) {
	// GET USERS
	user, err := t.UserRepo.FindOneByEmail(email)
	if err != nil {
		return "", err
	}

	// CEK PASSWORD
	if err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(password)); err != nil {
		return "", errors.New("Email or Password not match")
	}

	jwtHelper := jwt.NewJWT()
	token, err := jwtHelper.CreateJwtToken(t.Helper.Config.GetString("jwt.secret"), strconv.Itoa(int(user.ID)))
	if err != nil {
		return "", errors.New("Failed to generate token")
	}

	return token, nil
}

// Refresh ...
func (t *AuthUsecase) Refresh(userID uint) (string, error) {
	jwtHelper := jwt.NewJWT()
	token, err := jwtHelper.CreateJwtToken(t.Helper.Config.GetString("jwt.secret"), strconv.Itoa(int(userID)))
	if err != nil {
		return "", errors.New("Failed to generate token")
	}
	return token, nil
}
