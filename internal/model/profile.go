package model

import (
	_ "github.com/go-playground/validator/v10"
)

type ProfileRequest struct {
	FileId            string `json:"fileId"`
	BankAccountName   string `json:"bankAccountName" validate:"required,min=4,max=32"`
	BankAccountHolder string `json:"bankAccountHolder" validate:"required,min=4,max=32"`
	BankAccountNumber string `json:"bankAccountNumber" validate:"required,min=4,max=32"`
}

type ProfileLinkPhoneRequest struct {
	Phone string `json:"phone" validate:"required,e164"`
}

type ProfileLinkEmailRequest struct {
	Email string `json:"email" validate:"required,email"`
}

type ProfileResponse struct {
	Email             string `json:"email"`
	Phone             string `json:"phone"`
	FileId            string `json:"fileId"`
	FileUri           string `json:"fileUri"`
	FileThumbnailUri  string `json:"fileThumbnailUri"`
	BankAccountName   string `json:"bankAccountName"`
	BankAccountHolder string `json:"bankAccountHolder"`
	BankAccountNumber string `json:"bankAccountNumber"`
}
