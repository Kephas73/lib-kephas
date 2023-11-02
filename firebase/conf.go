package firebase

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Type                    string `json:"type"`
	ProjectID               string `json:"project_id"`
	PrivateKey              string `json:"private_key"`
	PrivateKeyID            string `json:"private_key_id"`
	ClientEmail             string `json:"client_email"`
	ClientID                string `json:"client_id"`
	AuthURI                 string `json:"auth_uri"`
	TokenURI                string `json:"token_uri"`
	AuthProviderX509CertURL string `json:"auth_provider_x509_cert_url"`
	ClientX509CertURL       string `json:"client_x509_cert_url"`
}

var firebaseConfig *Config

func getFirebaseConfigFromEnv(configKeys ...string) {
	configKey := "Firebase"
	for _, envKey := range configKeys {
		envKeyTrim := strings.TrimSpace(envKey)
		if envKeyTrim != "" {
			configKey = envKeyTrim
		}
	}

	firebaseConfig = &Config{}
	if err := viper.UnmarshalKey(configKey, &firebaseConfig); err != nil {
		err := fmt.Errorf("not found config name with env %q for social authorization with error: %+v", configKey, err)
		panic(err)
	}

	if firebaseConfig.Type == "" {
		err := fmt.Errorf("not found type property of config for %q", configKey)
		panic(err)
	}

	if firebaseConfig.ProjectID == "" {
		err := fmt.Errorf("not found project_id property of config for %q", configKey)
		panic(err)
	}

	if firebaseConfig.PrivateKey == "" {
		err := fmt.Errorf("not found private_key property of config for %q", configKey)
		panic(err)
	}

	if firebaseConfig.PrivateKeyID == "" {
		err := fmt.Errorf("not found private_key_id property of config for %q", configKey)
		panic(err)
	}

	if firebaseConfig.ClientEmail == "" {
		err := fmt.Errorf("not found client_email property of config for %q", configKey)
		panic(err)
	}

	if firebaseConfig.ClientID == "" {
		err := fmt.Errorf("not found client_id property of config for %q", configKey)
		panic(err)
	}

	if firebaseConfig.AuthURI == "" {
		err := fmt.Errorf("not found auth_uri property of config for %q", configKey)
		panic(err)
	}

	if firebaseConfig.TokenURI == "" {
		err := fmt.Errorf("not found token_uri property of config for %q", configKey)
		panic(err)
	}

	if firebaseConfig.AuthProviderX509CertURL == "" {
		err := fmt.Errorf("not found auth_provider_x509_cert_url property of config for %q", configKey)
		panic(err)
	}

	if firebaseConfig.ClientX509CertURL == "" {
		err := fmt.Errorf("not found client_x509_cert_url property of config for %q", configKey)
		panic(err)
	}
}
