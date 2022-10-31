package env

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"strings"
)

const (
	ConfigDefault    string = "config.json"
	EmptyConfig      string = ""
	EmptyValueConfig int    = 0
)

type ConfigEnv struct {
	Debug  bool `json:"debug,omitempty"`
	Logger struct {
		Path        string `json:"path,omitempty"`
		Prefix      string `json:"prefix,omitempty"`
		SentryDns   string `json:"sentry_dns,omitempty"`
		Description string `json:"description,omitempty"`
	} `json:"logger,omitempty"`
	Server struct {
		Port        string `json:"port,omitempty"`
		Host        string `json:"host,omitempty"`
		Environment string `json:"environment,omitempty"`
		ServiceName string `json:"service_name,omitempty"`
		RegionDC    string `json:"region_dc,omitempty"`
		Timeout     int    `json:"timeout,omitempty"`
		Description string `json:"description,omitempty"`
	} `json:"server,omitempty"`
	Context struct {
		Timeout int `json:"timeout,omitempty"`
	} `json:"context,omitempty"`
	SettingAPI struct {
		Path        string `json:"path,omitempty"`
		PathPrivate string `json:"path_private,omitempty"`
		Version     string `json:"version,omitempty"`
	} `json:"setting_api,omitempty"`
	HostPrivate struct {
	} `json:"host_private,omitempty"`
	DocsAPI struct {
		SwaggerPath string `json:"swagger_path,omitempty"`
	} `json:"docs_api,omitempty"`
	JWToken struct {
		SecretKey       string `json:"secret_key,omitempty"`
		AccessTokenTTL  int    `json:"access_token_ttl,omitempty"`
		RefreshTokenTTL int    `json:"refresh_token_ttl,omitempty"`
	} `json:"jw_token,omitempty"`
	Encrypt struct {
		AES struct {
			Key string `json:"key,omitempty"`
			IV  string `json:"iv,omitempty"`
		} `json:"aes,omitempty"`
	} `json:"encrypt,omitempty"`
}

var Environment *ConfigEnv

func SetupConfigEnv(in ...string) {
	cf := ConfigDefault

	for _, v := range in {
		cf = v
		break
	}

	viper.SetConfigFile(cf)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	Environment = &ConfigEnv{}
	if err := viper.Unmarshal(&Environment); err != nil {
		panic(err)
	}

	if Environment.Server.Port == EmptyConfig {
		err := errors.New("not found config with env SERVER.PORT for service")
		panic(err)
	}

	if Environment.Server.Environment == EmptyConfig {
		err := errors.New("not found config with SERVER.ENVIRONMENT for service")
		panic(err)
	}

	if Environment.Server.Timeout == EmptyValueConfig {
		err := errors.New("not found config with SERVER.TIMEOUT for service")
		panic(err)
	}

	if Environment.Logger.Prefix == EmptyConfig {
		err := errors.New("not found config with LOGGER.PREFIX for service")
		panic(err)
	}

	if Environment.Context.Timeout == EmptyValueConfig {
		err := errors.New("not found config with CONTEXT.TIMEOUT for service")
		panic(err)
	}

	if Environment.SettingAPI.Path == EmptyConfig {
		err := errors.New("not found config with SETTING_API.PATH for service")
		panic(err)
	}

	if Environment.SettingAPI.PathPrivate == EmptyConfig {
		err := errors.New("not found config with SETTING_API.PATH_PRIVATE for service")
		panic(err)
	}

	if Environment.SettingAPI.Version == EmptyConfig {
		err := errors.New("not found config with SETTING_API.VERSION for service")
		panic(err)
	}

	welcomeService := fmt.Sprintf("[Config Success] - Welcome service %s (ver api: %s)", strings.ToUpper(Environment.Server.ServiceName), Environment.SettingAPI.Version)
	fmt.Println(welcomeService)
}
