package security

import (
	"net/http"
	"time"

	"github.com/MishaNiki/lsait/backend/internal/app/utils"
	"github.com/dgrijalva/jwt-go"
)

type Security struct {
	expMinute time.Duration
	tokenKey  []byte
}

type Config struct {
	ExpMinute int    `json:"expMin"`
	TokenKey  string `json:"tokenKey"`
}

type Tokens struct {
	AccessToken  string
	RefrashToken string
}

type claims struct {
	Value int `json:"value"`
	jwt.StandardClaims
}

func New(config *Config) *Security {
	return &Security{
		expMinute: time.Duration(config.ExpMinute),
		tokenKey:  []byte(config.TokenKey),
	}
}

// NewTokens ...
func (sec *Security) NewTokens(value int) (*Tokens, error) {

	var err error
	tokens := &Tokens{}

	if tokens.AccessToken, err = sec.newAccesToken(value); err != nil {
		return nil, err
	}

	tokens.RefrashToken = utils.RandString(32)

	return tokens, nil
}

// newAccesToken ...
func (sec *Security) newAccesToken(value int) (string, error) {

	expirationTime := time.Now().Add(sec.expMinute * time.Minute)
	claims := &claims{
		Value: value,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(sec.tokenKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// VerificationAccessToken ...
func VerificationAccessToken(tokenKey []byte, token string) (int, int, error) {

	cl := &claims{}
	tkn, err := jwt.ParseWithClaims(token, cl, func(token *jwt.Token) (interface{}, error) {
		return tokenKey, nil
	})

	if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			return 0, http.StatusUpgradeRequired, err
		} else {
			return 0, http.StatusUnauthorized, err
		}
	}
	if tkn.Valid {
		return cl.Value, http.StatusOK, nil
	}
	return 0, http.StatusUnauthorized, err
}
