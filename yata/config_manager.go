package yata

import (
	"encoding/json"
	"fmt"
	"reflect"
	"sort"

	"strings"

	"github.com/fatih/structs"
	"golang.org/x/oauth2"
)

// Config TODO docs
type Config struct {
	GoogleDrive GoogleDriveJSON `json:"googleDrive"`
	Dropbox     DropboxJSON     `json:"dropbox"`
}

// GoogleDriveJSON TODO docs
// URL to generate secret file: https://developers.google.com/drive/v3/web/quickstart/go
type GoogleDriveJSON struct {
	OAuthToken *oauth2.Token `json:"oauthToken"`
	SecretFile string        `json:"secretFile"`
}

// DropboxJSON TODO docs
type DropboxJSON struct {
	APIKey string `json:"apiKey"`
}

// ConfigManager TODO docs
type ConfigManager struct {
	Config Config
}

// NewConfigManager TODO docs
func NewConfigManager() *ConfigManager {
	manager := &ConfigManager{}
	manager.LoadConfig()
	return manager
}

// GetKeys TODO docs
func (m ConfigManager) GetKeys() (keys []string, err error) {
	all, err := m.GetAll()
	if err != nil {
		return nil, err
	}

	keys = make([]string, 0)
	for k := range all {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys, nil
}

// GetValueForKey TODO docs
func (m ConfigManager) GetValueForKey(key string) (value string, err error) {
	configMap, err := m.GetAll()
	if err != nil {
		return "", err
	}

	return configMap[strings.ToLower(key)], nil
}

// GetAll TODO docs
func (m ConfigManager) GetAll() (values map[string]string, err error) {
	dat, err := GetDirectory().Config()
	if err != nil {
		return nil, err
	}

	var config Config
	if err = json.Unmarshal(dat, &config); err != nil {
		return nil, err
	}

	values = make(map[string]string)
	flatten("", config, values)

	return values, nil
}

// SetKey TODO docs
func (m ConfigManager) SetKey(key string, value interface{}) error {
	configMap := m.Config.Mapify()
	keys := strings.Split(strings.ToLower(key), ".")
	writeConfigKey(keys, value, configMap)

	dat, err := json.Marshal(configMap)
	if err != nil {
		return err
	}
	json.Unmarshal(dat, &m.Config)

	dat, err = json.MarshalIndent(m.Config, "", "\t")
	if err != nil {
		return err
	}

	GetDirectory().WriteConfig(dat)

	return nil
}

// LoadConfig TODO docs
func (m *ConfigManager) LoadConfig() {
	GetLogger().Verbose("Loading config")

	dat, err := GetDirectory().Config()
	if err != nil {
		m.Config = DefaultConfig()
	}

	var config Config
	if err = json.Unmarshal(dat, &config); err != nil {
		m.Config = DefaultConfig()
	}

	m.Config = config

	GetLogger().Verbose("Config loaded")
}

// DefaultConfig TODO docs
func DefaultConfig() Config {
	return Config{}
}

// Mapify TODO docs
func (c Config) Mapify() map[string]interface{} {
	return structs.Map(c)
}

// flatten TODO docs
func flatten(prefix string, v interface{}, m map[string]string) {
	val := reflect.ValueOf(v)
	if val.Kind() != reflect.Struct {
		return
	}

	var ns string

	for i := 0; i < val.Type().NumField(); i++ {
		field := val.Field(i)

		ns = strings.ToLower(val.Type().Field(i).Name)
		if prefix != "" {
			ns = fmt.Sprintf("%s.%s", prefix, ns)
		}

		if field.Kind() == reflect.Struct {
			flatten(ns, field.Interface(), m)
			continue
		}

		m[ns] = field.String()
	}
}

// writeConfigKey TODO docs
func writeConfigKey(keys []string, value interface{}, configMap map[string]interface{}) error {
	if len(keys) == 0 {
		return fmt.Errorf("You cannot update a config value for an unspecified key")
	}

	ck := keys[0]
	if len(keys) == 1 {
		for k := range configMap {
			if ck == strings.ToLower(k) &&
				reflect.ValueOf(configMap[k]).Kind() != reflect.Map {
				configMap[k] = value
			}
		}
		return nil
	}

	for k := range configMap {
		if ck != strings.ToLower(k) {
			continue
		}
		ck = k

		if reflect.ValueOf(configMap[ck]).Kind() != reflect.Map {
			return fmt.Errorf("Whoops! That key does not and cannot be made to exist due to conflicts")
		}

		return writeConfigKey(keys[1:], value, configMap[ck].(map[string]interface{}))
	}

	return fmt.Errorf("Key %s could not be found", ck)
}
