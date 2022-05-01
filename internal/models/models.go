package models

import (
	"encoding/hex"
	"math/rand"
	"time"

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
}

func NewToken() (*Token, error) {
	tokenCode, err := GenerateRandomString()
	if err != nil {
		return nil, err
	}

	return &Token{Token: tokenCode}, nil
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
