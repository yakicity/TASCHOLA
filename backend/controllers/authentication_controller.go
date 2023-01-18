package controllers

import (
	"log"
	"net/http"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"

	"taschola/models"
)

type LoginForm struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

const IdentityKey = "id"

func JWTInit() (*jwt.GinJWTMiddleware, error) {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "taschola",
		Key:         []byte("taschola"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: IdentityKey,

		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*LoginForm); ok {
				return jwt.MapClaims{
					IdentityKey: v.Name,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(ctx *gin.Context) interface{} {
			claims := jwt.ExtractClaims(ctx)

			return &LoginForm{
				Name: claims[IdentityKey].(string),
			}
		},
		// Authenticator: for Login function
		Authenticator: func(ctx *gin.Context) (interface{}, error) {
			var loginValues LoginForm
			if err := ctx.BindJSON(&loginValues); err != nil {
				log.Println("authenticator: BindJSON Error: " + err.Error())
				return "", jwt.ErrMissingLoginValues
			}
			name := loginValues.Name
			password := loginValues.Password

			user, err := models.GetUserByName(name)
			if err != nil {
				log.Println("authenticator: GetUserByName Error: " + err.Error())
				return nil, jwt.ErrFailedAuthentication
			}
			// check password is correct
			if string(hash(password)) != string(user.Password) {
				log.Println("authenticator: password is incorrect")
				log.Println("post password: " + password)
				log.Println("db password: " + string(user.Password))
				log.Println("hash password: " + string(hash(password)))

				return nil, jwt.ErrFailedAuthentication
			}

			return &LoginForm{
				Name:     name,
				Password: password,
			}, nil
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if v, ok := data.(*LoginForm); ok && v.Name == "admin" {
				return true
			}

			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message + " " + c.Request.URL.Path + " " + c.Request.Method,
			})
		},
		// Login
		//
		// @Summary Login
		// @Description Login
		// @Tags Authentication
		// @Accept  json
		// @Produce  json
		// @Param login body LoginForm true "Login"
		// @Success 200 {object} LoginResponse
		// @Failure 401 {object} ErrorResponse
		// @Failure 500 {object} ErrorResponse
		// @Router /v1/login [post]
		LoginResponse: func(c *gin.Context, code int, token string, expire time.Time) {
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"token":   token,
				"expire":  expire.Format(time.RFC3339),
				"message": "login success",
			})
		},
		LogoutResponse: func(c *gin.Context, code int) {
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": "logout success",
			})
		},
		RefreshResponse: func(c *gin.Context, code int, token string, expire time.Time) {
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"token":   token,
				"expire":  expire.Format(time.RFC3339),
				"message": "refresh success",
			})
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	errInit := authMiddleware.MiddlewareInit()
	if errInit != nil {
		log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
	}

	return authMiddleware, nil
}
