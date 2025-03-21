package JWToken

import (
	"fmt"
	"github.com/Kephas73/lib-kephas/base"
	"github.com/Kephas73/lib-kephas/error_code"
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
	KTokenClientSite     string = "client_site"
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
	DeviceIP           string        `json:"device_ip,omitempty"`
	ClientSite         int           `json:"client_site,omitempty"`
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

func NewToken(accountID string, deviceKind int, deviceIP string, clientSite int, configs ...*JWTConfig) (*TokenDetails, error) {
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
	td.DeviceIP = deviceIP
	td.ClientSite = clientSite

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
	atClaims[KTokenClientSite] = clientSite
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
	rtClaims[KTokenClientSite] = clientSite
	td.RefreshToken = jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)

	_, _, err := Sign(td, conf.SecretKey)

	return td, err
}

func ExtractToken(c echo.Context, shortLive bool, configs ...*JWTConfig) (accessUUID string, accountID string, deviceKind int, deviceIP string, clientSite int, bearToken string, errError *error_code.ErrorCode) {
	if conf == nil {
		if len(configs) == 0 {
			getJWTConfigFromEnv()
		} else {
			conf = configs[0]
		}
	}

	token, b, errError := GetToken(c, configs...)
	if errError != nil {
		return
	}

	bearToken = b
	// Check Token Valid Expire
	if token == nil {
		errError = error_code.NewError(error_code.ERROR_NEED_NEW_TOKEN, "not found or can not parse token", base.GetFunc())
		return
	}

	if !token.Valid {
		errError = error_code.NewError(error_code.ERROR_TOKEN_INVALID, "token is invalid", base.GetFunc())
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUUID, ok = claims[KTokenAccessUUIDKey].(string)
		if !ok {
			errError = error_code.NewError(error_code.ERROR_TOKEN_INVALID, "token Invalid or Expire - not found access uuid key", base.GetFunc())
			return
		}

		accountID, ok = claims[KTokenUserIDKey].(string)
		if !ok {
			errError = error_code.NewError(error_code.ERROR_TOKEN_INVALID, "invalid Bearer Token - not found user uuid key", base.GetFunc())
			return
		}

		if dv, err := base.StringToInt64(fmt.Sprintf("%v", claims[KTokenDeviceKindKey])); err == nil {
			deviceKind = int(dv)
		}

		if deviceKind < 0 {
			errError = error_code.NewError(error_code.ERROR_TOKEN_INVALID, fmt.Sprintf("no device info: %s", accountID), base.GetFunc())
			return
		}

		if ip, ok := claims[KTokenDeviceIPKey].(string); ok && ip != "" {
			deviceIP = ip
		}

		if site, err := base.StringToInt64(fmt.Sprintf("%v", claims[KTokenClientSite])); err == nil {
			clientSite = int(site)
		}

		if clientSite < 0 {
			errError = error_code.NewError(error_code.ERROR_TOKEN_INVALID, fmt.Sprintf("no client site info: %s", accountID), base.GetFunc())
			return
		}

		expired, ok := claims[KTokenExpKey].(float64)
		if !ok {
			errError = error_code.NewError(error_code.ERROR_TOKEN_INVALID, "invalid Bearer Token - not found exp", base.GetFunc())
			return
		}

		if float64(time.Now().Unix()) > expired {
			errError = error_code.NewError(error_code.ERROR_NEED_NEW_TOKEN, "invalid Bearer Token - token is expired", base.GetFunc())
			return
		}

		if shortLive {
			shortLiveExp, ok := claims[KTokenSecureExpKey].(float64)
			if !ok {
				errError = error_code.NewError(error_code.ERROR_TOKEN_INVALID, "invalid Bearer Token - can not get secure expiration", base.GetFunc())
				return
			} else {
				if int64(shortLiveExp) < time.Now().Unix() {
					errError = error_code.NewError(error_code.ERROR_NEED_NEW_TOKEN, "invalid Bearer Token - token was expired", base.GetFunc())
				}
			}
		}

		return
	}

	return
}

func ExtractRefreshToken(c echo.Context, configs ...*JWTConfig) (refreshUUID, accountID string, deviceKind, clientSite int, errError *error_code.ErrorCode) {
	if conf == nil {
		if len(configs) == 0 {
			getJWTConfigFromEnv()
		} else {
			conf = configs[0]
		}
	}

	token, _, errError := GetToken(c, configs...)
	if errError != nil {
		return
	}

	if token == nil {
		errError = error_code.NewError(error_code.ERROR_REFRESH_TOKEN_INVALID, "token Invalid or Expire", base.GetFunc())
		return
	}

	// Check Token Valid Expire
	if !token.Valid {
		errError = error_code.NewError(error_code.ERROR_REFRESH_TOKEN_INVALID, "invalid Bearer Token", base.GetFunc())
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		refreshUUID, ok = claims[KTokenRefreshUUIDKey].(string)
		if !ok {
			errError = error_code.NewError(error_code.ERROR_REFRESH_TOKEN_INVALID, "invalid Bearer Token", base.GetFunc())
			return
		}

		accountID, ok = claims[KTokenUserIDKey].(string)
		if !ok {
			errError = error_code.NewError(error_code.ERROR_REFRESH_TOKEN_INVALID, "invalid Bearer Token", base.GetFunc())
			return
		}

		if dv, err := base.StringToInt64(fmt.Sprintf("%v", claims[KTokenDeviceKindKey])); err == nil {
			deviceKind = int(dv)
		}

		if deviceKind < 0 {
			errError = error_code.NewError(error_code.ERROR_REFRESH_TOKEN_INVALID, fmt.Sprintf("no device info: %s", accountID), base.GetFunc())
			return
		}

		if site, err := base.StringToInt64(fmt.Sprintf("%v", claims[KTokenClientSite])); err == nil {
			clientSite = int(site)
		}

		if clientSite < 0 {
			errError = error_code.NewError(error_code.ERROR_REFRESH_TOKEN_INVALID, fmt.Sprintf("no client site info: %s", accountID), base.GetFunc())
			return
		}

		return
	}

	errError = error_code.NewError(error_code.ERROR_REFRESH_TOKEN_INVALID, "invalid Bearer Token", base.GetFunc())
	return
}

func GetToken(c echo.Context, configs ...*JWTConfig) (token *jwt.Token, bearToken string, errCode *error_code.ErrorCode) {
	// Get token from Echo
	cookie := c.Request().Header.Get("Authorization")
	if len(cookie) < 1 {
		errCode = error_code.NewError(error_code.ERROR_AUTHORIZE, "invalid Bearer Token - not found Authorization key", base.GetFunc())
		return
	}

	if !strings.Contains(cookie, " ") {
		errCode = error_code.NewError(error_code.ERROR_UNAUTHORIZED_USER, fmt.Sprintf("AuthToken Have No Space Split: %s", cookie), base.GetFunc())
		return
	}

	bearToken = strings.Split(cookie, " ")[1]
	tokenParse, err := jwt.Parse(bearToken, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// Signing Key
		signKey := conf.SecretKey
		return []byte(signKey), nil
	})

	if err != nil {
		errCode = error_code.NewError(error_code.ERROR_TOKEN_INVALID, err.Error(), base.GetFunc())
		return
	}

	token = tokenParse
	return
}
