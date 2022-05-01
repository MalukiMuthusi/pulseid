package models

import (
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"

	"github.com/MalukiMuthusi/pulseid/internal/logger"
	"gorm.io/gorm"
)

type TokenParameter struct {
	Token string `uri:"token" binding:"required"`
}

type BasicError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type Token struct {
	gorm.Model

	Token string `json:"token"`

	Recalled bool `json:"recalled"`
}

func NewToken() (*Token, error) {
	tokenCode, err := GenerateRandomString()
	if err != nil {
		return nil, err
	}

	return &Token{Token: tokenCode}, nil
}

func (t *Token) CheckValidity() bool {

	if t.Recalled {
		return false
	}

	createdAt := t.CreatedAt

	now := time.Now()

	duration, err := time.ParseDuration(fmt.Sprintf("%ds", 24*7*60*60))
	if err != nil {
		logger.Log.Error(err)
	}

	expiryDate := createdAt.Add(duration)
	logger.Log.Info(expiryDate.String())

	return now.Before(expiryDate)

}

type AuthHeader struct {
	Authorization string `json:"Authorization" header:"Authorization" binding:"required"`
}

func GenerateRandomString() (string, error) {
	src := rand.New(rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, 12/2)

	if _, err := src.Read(b); err != nil {
		return "", nil
	}

	return hex.EncodeToString(b)[:12], nil
}
