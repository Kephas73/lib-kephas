package JWToken

import (
	"errors"
	"fmt"
	"github.com/Kephas73/lib-kephas/base"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"strings"
	"time"
)

const (
	KTokenAuthorizedKey  string = "authorized"
	KTokenAccessUUIDKey  string = "access_uuid"
	KTokenRefreshUUIDKey string = "refresh_uuid"
	KTokenUserIDKey      string = "user_id"
	KTokenExpKey         string = "exp"
	KTokenSecureExpKey   string = "secure_exp"
	KTokenDeviceKindKey  string = "device_kind"
	KTokenDeviceIPKey    string = "device_ip"
)

const (
	KCacheExpiresInOneHour  = time.Hour
	KCacheExpiresInOneDay   = 24 * KCacheExpiresInOneHour
	KCacheExpiresInOneWeek  = 7 * KCacheExpiresInOneDay
	KCacheExpiresInOneMonth = 30 * KCacheExpiresInOneDay
)

type TokenDetails struct {
	UserID             string        `json:"user_id,omitempty"`
	DeviceKind         int           `json:"device_kind,omitempty"`
	AccessUUID         uuid.UUID     `json:"access_uuid,omitempty"`
	AccessToken        *jwt.Token    `json:"access_token,omitempty"`
	RefreshUUID        uuid.UUID     `json:"refresh_uuid,omitempty"`
	RefreshToken       *jwt.Token    `json:"refresh_token,omitempty"`
	SignedAccessToken  string        `json:"signed_access_token,omitempty"`
	SignedRefreshToken string        `json:"signed_refresh_token,omitempty"`
	AtExpires          time.Duration `json:"at_expires,omitempty"`
	RtExpires          time.Duration `json:"rt_expires,omitempty"`
	SecureExpire       time.Duration `json:"secure_expire,omitempty"`
}

func Sign(token *TokenDetails, signer string) (string, string, error) {
	signedAt, err := token.AccessToken.SignedString([]byte(signer))
	if err != nil {
		return "", "", err
	}
	token.SignedAccessToken = signedAt

	signedRt, err := token.RefreshToken.SignedString([]byte(signer))
	if err != nil {
		return "", "", err
	}
	token.SignedRefreshToken = signedRt
	return signedAt, signedRt, nil
}

func NewToken(accountID string, deviceKind int, deviceIP string, configs ...*JWTConfig) (*TokenDetails, error) {
	if conf == nil {
		if len(configs) == 0 {
			getJWTConfigFromEnv()
		} else {
			conf = configs[0]
		}
	}

	td := &TokenDetails{}
	td.UserID = accountID
	td.DeviceKind = deviceKind

	td.AtExpires = time.Duration(conf.AccessTokenTTL) * time.Second
	td.SecureExpire = time.Duration(conf.SecureTokenTTL) * time.Second
	td.RtExpires = time.Duration(conf.RefreshTokenTTL) * time.Second
	accessUUID := uuid.New()
	td.AccessUUID = accessUUID
	atClaims := jwt.MapClaims{}
	atClaims[KTokenAuthorizedKey] = true
	atClaims[KTokenAccessUUIDKey] = accessUUID.String()
	atClaims[KTokenUserIDKey] = accountID
	atClaims[KTokenExpKey] = time.Now().Add(td.AtExpires).Unix()
	atClaims[KTokenDeviceKindKey] = deviceKind
	atClaims[KTokenDeviceIPKey] = deviceIP
	atClaims[KTokenSecureExpKey] = time.Now().Add(td.SecureExpire).Unix()
	td.AccessToken = jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)

	refreshUUID := uuid.New()
	td.RefreshUUID = refreshUUID
	rtClaims := jwt.MapClaims{}
	rtClaims[KTokenRefreshUUIDKey] = refreshUUID.String()
	rtClaims[KTokenUserIDKey] = accountID
	rtClaims[KTokenExpKey] = time.Now().Add(td.RtExpires).Unix()
	rtClaims[KTokenDeviceKindKey] = deviceKind
	rtClaims[KTokenDeviceIPKey] = deviceIP
	td.RefreshToken = jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)

	_, _, err := Sign(td, conf.SecretKey)

	return td, err
}

func ExtractToken(c echo.Context, shortLive bool, configs ...*JWTConfig) (accessUUID string, accountID string, deviceKind int, deviceIP string, bearToken string, err error) {
	if conf == nil {
		if len(configs) == 0 {
			getJWTConfigFromEnv()
		} else {
			conf = configs[0]
		}
	}

	token, b, err := GetToken(c, configs...)
	bearToken = b

	// Check Token Valid Expire
	if token == nil {
		err = errors.New("not found or can not parse token")
		return
	}

	if !token.Valid {
		err = errors.New("token is invalid")
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUUID, ok = claims[KTokenAccessUUIDKey].(string)
		if !ok {
			err = errors.New("token Invalid or Expire - not found access uuid key")
			return
		}

		accountID, ok = claims[KTokenUserIDKey].(string)
		if !ok {
			err = errors.New("invalid Bearer Token - not found user uuid key")
			return
		}

		if dv, err := base.StringToInt64(fmt.Sprintf("%v", claims[KTokenDeviceKindKey])); err == nil {
			deviceKind = int(dv)
		}

		if deviceKind < 0 {
			err = errors.New(fmt.Sprintf("no device info: %s", accountID))
			return
		}

		if ip, ok := claims[KTokenDeviceIPKey].(string); ok && ip != "" {
			deviceIP = ip
		}

		if shortLive {
			shortLiveExp, ok := claims[KTokenSecureExpKey].(float64)
			if !ok {
				err = errors.New("invalid Bearer Token - can not get secure expiration")
				return
			} else {
				if int64(shortLiveExp) < time.Now().Unix() {
					err = errors.New("invalid Bearer Token - token was expired")
				}
			}
		}

		return
	}

	return
}

func ExtractRefreshToken(c echo.Context, configs ...*JWTConfig) (refreshUUID string, accountID string, err error) {
	if conf == nil {
		if len(configs) == 0 {
			getJWTConfigFromEnv()
		} else {
			conf = configs[0]
		}
	}

	token, _, err := GetToken(c, configs...)

	if token == nil {
		err = errors.New("token Invalid or Expire")
		return
	}

	// Check Token Valid Expire
	if !token.Valid {
		err = errors.New("invalid Bearer Token")
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		refreshUUID, ok = claims[KTokenRefreshUUIDKey].(string)
		if !ok {
			err = errors.New("invalid Bearer Token")
			return
		}

		accountID, ok = claims[KTokenUserIDKey].(string)
		if !ok {
			err = errors.New("invalid Bearer Token")
			return
		}

		return
	}

	err = errors.New("invalid Bearer Token")
	return
}

func GetToken(c echo.Context, configs ...*JWTConfig) (token *jwt.Token, bearToken string, err error) {
	// Get token from Echo
	cookie := c.Request().Header.Get("Authorization")
	if len(cookie) < 1 {
		err = errors.New("invalid Bearer Token - not found Authorization key")
		return
	}

	if !strings.Contains(cookie, " ") {
		errorLog := fmt.Sprintf("AuthToken Have No Space Split: %s", cookie)
		err = errors.New(errorLog)
		return
	}

	bearToken = strings.Split(cookie, " ")[1]
	token, err = jwt.Parse(bearToken, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// Signing Key
		signKey := conf.SecretKey
		return []byte(signKey), nil
	})

	return
}
