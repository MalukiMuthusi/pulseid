package models

import "gorm.io/gorm"

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

type AuthHeader struct {
	Authorization string `json:"Authorization" header:"Authorization" binding:"required"`
}
