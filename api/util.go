package api

import "github.com/labstack/echo/v4"

func GetContextDataString(ctx echo.Context, key string, defaultValues ...string) string {
	defaultValue := ""
	if len(defaultValues) > 0 {
		defaultValue = defaultValues[0]
	}

	userUUIDRaw := ctx.Get(key)
	if userUUIDRaw != nil {
		if res, ok := userUUIDRaw.(string); ok {
			return res
		}
	}

	return defaultValue
}

func SetContextDataString(ctx echo.Context, key string, value interface{}) {
	ctx.Set(key, value)
}