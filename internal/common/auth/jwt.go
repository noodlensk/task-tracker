package auth

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/cristalhq/jwt/v4"
)

type User struct {
	UUID  string
	Email string
	Role  string

	Name string
}

type UserClaims struct {
	jwt.RegisteredClaims
	UserUID   string
	UserName  string
	UserEmail string
	UserRole  string
}

func GetToken(secret string, user User) (string, error) {
	key := []byte(secret)

	signer, err := jwt.NewSignerHS(jwt.HS256, key)
	if err != nil {
		return "", err
	}

	claims := &UserClaims{
		RegisteredClaims: jwt.RegisteredClaims{},
		UserUID:          user.UUID,
		UserName:         user.Name,
		UserEmail:        user.Email,
		UserRole:         user.Role,
	}

	builder := jwt.NewBuilder(signer)

	token, err := builder.Build(claims)
	if err != nil {
		return "", err
	}

	return token.String(), nil
}

func ParseToken(secret, token string) (user User, err error) {
	key := []byte(secret)

	verifier, err := jwt.NewVerifierHS(jwt.HS256, key)
	if err != nil {
		return User{}, err
	}

	parsedToken, err := jwt.Parse([]byte(token), verifier)
	if err != nil {
		return User{}, err
	}

	var claims UserClaims

	if err := json.Unmarshal(parsedToken.Claims(), &claims); err != nil {
		return User{}, err
	}

	if !claims.IsValidAt(time.Now()) {
		return User{}, errors.New("token expired")
	}

	return User{
		UUID:  claims.UserUID,
		Email: claims.UserEmail,
		Role:  claims.UserRole,
		Name:  claims.UserName,
	}, nil
}
