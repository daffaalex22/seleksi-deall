package middlewares

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/daffaalex22/seleksi-deall/helper/err"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	RolePublic = 0
	RoleUser   = 1
	RoleAdmin  = 2

	PublicLevel = []int{RolePublic}
	UserLevel   = []int{RolePublic, RoleUser}
	AdminLevel  = []int{RoleAdmin, RoleUser, RolePublic}
)

type JwtCustomClaims struct {
	UserId string /*`json:"userId"`*/
	jwt.StandardClaims
	Roles []int
}

type ConfigJWT struct {
	SecretJWT       string
	ExpiresDuration int
}

func (jwtConf *ConfigJWT) Init() middleware.JWTConfig {
	return middleware.JWTConfig{
		Claims:     &JwtCustomClaims{},
		SigningKey: []byte(jwtConf.SecretJWT),
	}
}

func (jwtConf *ConfigJWT) GenerateTokenJWT(UserId string, isAdmin bool) (string, error) {
	var Roles []int
	if isAdmin {
		Roles = AdminLevel
	} else {
		Roles = UserLevel
	}

	claims := JwtCustomClaims{
		UserId,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(int64(jwtConf.ExpiresDuration))).Unix(),
		},
		Roles,
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, res := t.SignedString([]byte(jwtConf.SecretJWT))

	return token, res
}

func ExtractJWT(ctx echo.Context) (payload *JwtCustomClaims, res error) {
	header := ctx.Request().Header
	bearerToken := header.Get("Authorization")

	// extract bearerToken if exist
	tokenPayload := &JwtCustomClaims{}
	if len(bearerToken) > 0 {
		token := strings.Split(bearerToken, " ")[1]
		token = strings.Split(token, ".")[1]
		tokenByte, res := jwt.DecodeSegment(token)
		if res != nil {
			return nil, res
		}
		if res := json.Unmarshal(tokenByte, tokenPayload); res != nil {
			return nil, res
		}
	}
	return tokenPayload, nil
}

func ValidateAuthorization(ctx echo.Context, allowedRoles []int) (payload *JwtCustomClaims, res error) {
	tokenPayload, res := ExtractJWT(ctx)
	if res != nil {
		return nil, res
	}

	if len(allowedRoles) == 0 || len(allowedRoles) == 1 && allowedRoles[0] == PublicLevel[0] {
		return tokenPayload, nil
	}

	if tokenPayload.UserId == "" {
		return nil, err.ErrUnauthorized
	}

	if validateRole(tokenPayload.Roles, allowedRoles) {
		return tokenPayload, nil
	}

	return nil, err.ErrUnauthorized
}

func validateRole(roles []int, allowedRoles []int) bool {
	for _, role := range roles {
		for _, allowedRole := range allowedRoles {
			if role == allowedRole {
				return true
			}
		}
	}
	return false
}
