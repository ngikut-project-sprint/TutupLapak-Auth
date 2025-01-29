package handler

import (
	"fmt"
	"net/http"
	"tutuplapak-auth/internal/config"
	"tutuplapak-auth/internal/model"

	"github.com/labstack/echo/v4"
)

func GetProfilehandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		userId := c.Get("USER_ID") // replace YourType with the actual type that has Claims field
		if userId == "" {
			return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
		}

		db := config.DB()
		userModel := &model.User{}
		db.First(&userModel, userId)

		if userModel.ID == 0 {
			return echo.NewHTTPError(http.StatusNotFound, "User is not found")
		}

		response := model.ProfileResponse{
			Email:             userModel.Email,
			Phone:             userModel.Phone,
			FileId:            userModel.FileId,
			FileUri:           "",
			FileThumbnailUri:  "",
			BankAccountName:   userModel.BankAccountName,
			BankAccountHolder: userModel.BankAccountHolder,
			BankAccountNumber: userModel.BankAccountNumber,
		}

		return c.JSON(http.StatusOK, response)
	}
}

func PutProfilehandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		userId := c.Get("USER_ID") // replace YourType with the actual type that has Claims field
		if userId == "" {
			return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
		}

		profile := new(model.ProfileRequest)
		if err := c.Bind(profile); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
		}

		if err := c.Validate(profile); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
		}

		db := config.DB()
		userModel := &model.User{}
		db.First(&userModel, userId)

		if userModel.ID == 0 {
			return echo.NewHTTPError(http.StatusNotFound, "User is not found")
		}

		db.Model(&userModel).Updates(model.User{
			Email:             userModel.Email,
			Phone:             userModel.Phone,
			FileId:            profile.FileId,
			BankAccountName:   profile.BankAccountName,
			BankAccountHolder: profile.BankAccountHolder,
			BankAccountNumber: profile.BankAccountNumber,
		})

		response := model.ProfileResponse{
			Email:             userModel.Email,
			Phone:             userModel.Phone,
			FileId:            userModel.FileId,
			FileUri:           "",
			FileThumbnailUri:  "",
			BankAccountName:   userModel.BankAccountName,
			BankAccountHolder: userModel.BankAccountHolder,
			BankAccountNumber: userModel.BankAccountNumber,
		}

		return c.JSON(http.StatusOK, response)
	}
}

func PostLinkPhone() echo.HandlerFunc {
	return func(c echo.Context) error {
		userId := c.Get("USER_ID") // replace YourType with the actual type that has Claims field
		if userId == "" {
			return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
		}

		profile := new(model.ProfileLinkPhoneRequest)
		if err := c.Bind(profile); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
		}

		if err := c.Validate(profile); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
		}

		db := config.DB()
		userModel := &model.User{}

		db.Where("phone = ?", profile.Phone).First(&userModel)
		fmt.Println(userModel)
		if userModel.ID != 0 {
			return echo.NewHTTPError(http.StatusConflict, "Phone is taken")
		}

		db.First(&userModel, userId)

		if userModel.ID == 0 {
			return echo.NewHTTPError(http.StatusNotFound, "User is not found")
		}

		db.Model(&userModel).Update("phone", profile.Phone)

		response := model.ProfileResponse{
			Email:             userModel.Email,
			Phone:             userModel.Phone,
			FileId:            userModel.FileId,
			FileUri:           "",
			FileThumbnailUri:  "",
			BankAccountName:   userModel.BankAccountName,
			BankAccountHolder: userModel.BankAccountHolder,
			BankAccountNumber: userModel.BankAccountNumber,
		}

		return c.JSON(http.StatusOK, response)
	}
}

func PostLinkEmail() echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("masuk sini dong")
		userId := c.Get("USER_ID") // replace YourType with the actual type that has Claims field
		if userId == "" {
			return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
		}

		profile := new(model.ProfileLinkEmailRequest)
		if err := c.Bind(profile); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
		}

		if err := c.Validate(profile); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
		}

		db := config.DB()
		userModel := &model.User{}

		db.Where("email = ?", profile.Email).First(&userModel)
		fmt.Println(userModel)
		if userModel.ID != 0 {
			return echo.NewHTTPError(http.StatusConflict, "Email is taken")
		}

		db.First(&userModel, userId)

		if userModel.ID == 0 {
			return echo.NewHTTPError(http.StatusNotFound, "User is not found")
		}

		db.Model(&userModel).Update("email", profile.Email)

		response := model.ProfileResponse{
			Email:             userModel.Email,
			Phone:             userModel.Phone,
			FileId:            userModel.FileId,
			FileUri:           "",
			FileThumbnailUri:  "",
			BankAccountName:   userModel.BankAccountName,
			BankAccountHolder: userModel.BankAccountHolder,
			BankAccountNumber: userModel.BankAccountNumber,
		}

		return c.JSON(http.StatusOK, response)
	}
}
