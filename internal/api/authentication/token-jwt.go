package authentication

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/loremcookie/go-home/backend/internal/api/models"
	"net/http"
	"os"
	"strings"
	"time"
)

var (
	ErrMalformedToken          error = errors.New("malformed token")
	ErrUnexpectedSigningMethod error = errors.New("unexpected signing method")
	ErrUnusableToken           error = errors.New("unusable token")
)

//NewToken creates a new jwt token.
//The function generates the token signs it and maps the claims.
func NewToken(metadata *models.TokenClaims) (map[string]string, error) {
	var err error

	//Create unsigned access accessToken
	accessToken := jwt.New(jwt.SigningMethodHS256)

	//Create map accessClaims to store access jwt metadata
	accessClaims := accessToken.Claims.(jwt.MapClaims)

	//Store access jwt metadata
	accessClaims["username"] = metadata.Username
	accessClaims["admin"] = metadata.Admin
	accessClaims["exp"] = time.Now().Add(24 * time.Hour).Unix()

	//Sign accessToken
	signedAccessToken, err := accessToken.SignedString([]byte(os.Getenv("API_SECRET")))
	if err != nil {
		return map[string]string{}, err
	}

	//Create unsigned refreshToken
	refreshToken := jwt.New(jwt.SigningMethodHS256)

	//Create private claims for refresh token
	refreshClaims := refreshToken.Claims.(jwt.MapClaims)

	//Fill metadata for in claims
	refreshClaims["username"] = metadata.Username
	refreshClaims["exp"] = time.Now().Add(504 * time.Hour)

	//Sign refreshToken
	signedRefreshToken, err := refreshToken.SignedString([]byte(os.Getenv("API_SECRET")))
	if err != nil {
		return map[string]string{}, err
	}

	return map[string]string{
		"access_token":  signedAccessToken,
		"refresh_token": signedRefreshToken,
	}, nil
}

//ParseTokenString parses string to jwt token
func ParseTokenString(tokenString string) (*jwt.Token, error) {
	var err error

	//Parse token and verify signing
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrUnexpectedSigningMethod
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}

//GetTokenFromHeader returns the token and verify's the signing method
func GetTokenFromHeader(r *http.Request) (*jwt.Token, error) {
	var err error

	//Gets  header ands splits it
	authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
	if len(authHeader) != 2 {
		return nil, ErrMalformedToken
	}

	//Get jwt token out of list
	jwtToken := authHeader[1]

	//Parse token
	token, err := ParseTokenString(jwtToken)
	if err != nil {
		return nil, err
	}

	return token, nil
}

//VerifyToken verifies the usability of the token.
//Reasons a token could be unusable is eg. expired
func VerifyToken(token *jwt.Token) error {
	//Check if token is usable
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return ErrUnusableToken
	}

	return nil
}

//GetTokenMetadata retrieves the username and permission from a accessToken
func GetTokenMetadata(token *jwt.Token) *models.TokenClaims {
	var ok bool

	//Make token claims
	claims, ok := token.Claims.(jwt.MapClaims)
	//Verify the claims
	if !ok && !token.Valid {
		return nil
	}

	//Create empty private token claims model to save metadata in
	var metadata models.TokenClaims

	//Get username
	metadata.Username, ok = claims["username"].(string)
	if !ok {
		return nil
	}

	//Get if the user is in the admin group
	metadata.Admin, ok = claims["admin"].(bool)
	if !ok {
		return nil
	}

	return &metadata
}
