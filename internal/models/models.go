package models

type TokenParameter struct {
	Token string `uri:"token" binding:"required"`
}

type BasicError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
