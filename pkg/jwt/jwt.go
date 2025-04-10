package jwt

import (
	app_model "authService/internal/models/app-model"
	user_model "authService/internal/models/user-model"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func NewToken(user *user_model.User, app *app_model.App, duration time.Duration) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["userID"] = user.ID
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(duration).Unix()
	claims["app_id"] = app.ID

	tokenString, err := token.SignedString([]byte(app.Secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
