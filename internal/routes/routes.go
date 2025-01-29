package routes

import (
	"net/http"
	"time"
	"tutuplapak-auth/internal/handler"
	"tutuplapak-auth/internal/middleware"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type jwtCustomClaims struct {
	Email string `json:"email"`
	ID    int    `json:"id"`
	jwt.RegisteredClaims
}

type CustomValidator struct {
	validator *validator.Validate
}

// Validate implements echo.Validator.
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func healthCheck(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World tes!")
	})

	e.GET("/healthz", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"status": "healthy",
			"time":   time.Now().String(),
		})
	})
}

func NewRouter() *echo.Echo {
	e := echo.New()

	healthCheck(e)
	authRoute(e)
	profileRoute(e)

	return e
}

func authRoute(e *echo.Echo) {
	validate := validator.New()
	e.Validator = &CustomValidator{validator: validate}

	e.POST("v1/login/email", handler.AuthEmailLogin())
	e.POST("v1/register/email", handler.AuthEmailRegister())
	e.POST("v1/login/phone", handler.AuthPhoneLogin())
	e.POST("v1/register/phone", handler.AuthPhoneRegister())
}

func profileRoute(e *echo.Echo) {
	validate := validator.New()
	e.Validator = &CustomValidator{validator: validate}

	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(jwtCustomClaims)
		},
		SigningKey: []byte("secret"),
	}
	user := e.Group("/v1/user")
	user.Use(echojwt.WithConfig(config))
	user.GET("", middleware.AuthMiddleware(handler.GetProfilehandler()))
	user.PUT("", middleware.AuthMiddleware(handler.PutProfilehandler()))
	user.POST("/link/phone", middleware.AuthMiddleware(handler.PostLinkPhone()))
	user.POST("/link/email", middleware.AuthMiddleware(handler.PostLinkEmail()))
}
