package handler

import (
	"fmt"
	"net/http"
	"time"
	"tutuplapak-auth/internal/config"
	"tutuplapak-auth/internal/model"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type jwtCustomClaims struct {
	Email string `json:"email"`
	ID    int    `json:"id"`
	jwt.RegisteredClaims
}

func AuthEmailLogin() echo.HandlerFunc {
	return func(c echo.Context) error {
		user := new(model.AuthEmailRequest)

		if c.Request().Header.Get("Content-Type") != "application/json" {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid content type")
		}

		if err := c.Bind(user); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
		}

		if err := c.Validate(user); err != nil {
			fmt.Println(err)
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
		}

		db := config.DB()
		userModel := &model.User{}
		db.First(&userModel, "email = ?", user.Email)

		if userModel.ID == 0 {
			return echo.NewHTTPError(http.StatusNotFound, "Email is not found")
		}

		if err := bcrypt.CompareHashAndPassword([]byte(userModel.Password), []byte(user.Password)); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Validation error")
		}

		claims := &jwtCustomClaims{
			userModel.Email,
			userModel.ID,
			jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid input")
		}

		response := model.AuthResponse{
			Email: user.Email,
			Token: t,
		}

		return c.JSON(http.StatusOK, response)
	}
}

func AuthEmailRegister() echo.HandlerFunc {
	return func(c echo.Context) error {
		user := new(model.AuthEmailRequest)

		if c.Request().Header.Get("Content-Type") != "application/json" {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid content type")
		}

		if err := c.Bind(user); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
		}

		if err := c.Validate(user); err != nil {
			fmt.Println(err)
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
		}

		db := config.DB()
		userModel := &model.User{}

		db.First(&userModel, "email = ?", user.Email)

		fmt.Println(`dua `, userModel)

		if userModel.ID != 0 {
			return echo.NewHTTPError(http.StatusConflict, "Email is exist")
		}

		userModel = &model.User{
			Email:    user.Email,
			Password: string(hashedPassword),
		}

		if err := db.Create(&userModel).Error; err != nil {
			fmt.Println("err ", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
		}

		claims := &jwtCustomClaims{
			userModel.Email,
			userModel.ID,
			jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid input")
		}

		response := model.AuthResponse{
			Email: user.Email,
			Token: t,
		}

		return c.JSON(http.StatusOK, response)
	}
}

func AuthPhoneLogin() echo.HandlerFunc {
	return func(c echo.Context) error {
		user := new(model.AuthPhoneRequest)

		if c.Request().Header.Get("Content-Type") != "application/json" {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid content type")
		}

		if err := c.Bind(user); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
		}

		if err := c.Validate(user); err != nil {
			fmt.Println(err)
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
		}

		db := config.DB()
		userModel := &model.User{}
		db.First(&userModel, "phone = ?", user.Phone)

		if userModel.ID == 0 {
			return echo.NewHTTPError(http.StatusNotFound, "Phone is not found")
		}

		if err := bcrypt.CompareHashAndPassword([]byte(userModel.Password), []byte(user.Password)); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Validation error")
		}

		claims := &jwtCustomClaims{
			userModel.Email,
			userModel.ID,
			jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid input")
		}

		response := model.AuthResponse{
			Phone: user.Phone,
			Token: t,
		}

		return c.JSON(http.StatusOK, response)
	}
}

func AuthPhoneRegister() echo.HandlerFunc {
	return func(c echo.Context) error {
		user := new(model.AuthPhoneRequest)

		if c.Request().Header.Get("Content-Type") != "application/json" {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid content type")
		}

		if err := c.Bind(user); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
		}

		if err := c.Validate(user); err != nil {
			fmt.Println(err)
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
		}

		db := config.DB()
		userModel := &model.User{}

		db.First(&userModel, "phone = ?", user.Phone)

		fmt.Println(`dua `, userModel)

		if userModel.ID != 0 {
			return echo.NewHTTPError(http.StatusConflict, "Phone is exist")
		}

		userModel = &model.User{
			Phone:    user.Phone,
			Password: string(hashedPassword),
		}

		if err := db.Create(&userModel).Error; err != nil {
			fmt.Println("err ", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
		}

		claims := &jwtCustomClaims{
			userModel.Email,
			userModel.ID,
			jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid input")
		}

		response := model.AuthResponse{
			Phone: user.Phone,
			Token: t,
		}

		return c.JSON(http.StatusOK, response)
	}
}
