package yata

import (
	"encoding/json"
	"fmt"
	"reflect"
	"sort"

	"strings"

	"github.com/fatih/structs"
)

// Config TODO docs
type Config struct {
	GoogleDrive GoogleDrive `json:"drive,omitempty"`
	Dropbox     Dropbox     `json:"dropbox,omitempty"`
}

// GoogleDrive TODO docs
type GoogleDrive struct {
	APIKey string `json:"apiKey"`
}

// Dropbox TODO docs
type Dropbox struct {
	APIKey string `json:"apiKey"`
}

// ConfigManager TODO docs
type ConfigManager struct {
	config *Config
}

// GetValueForKey TODO docs
func (m ConfigManager) GetValueForKey(key string) {
	// configMap := m.config.mapify()
	// configMap[key]
}

// GetAll TODO docs
func (m ConfigManager) GetAll() (values map[string]interface{}, err error) {
	dirService := NewDirectoryService()
	dat, err := dirService.GetConfig()
	if err != nil {
		return nil, err
	}

	var config Config
	if err = json.Unmarshal(dat, &config); err != nil {
		return nil, err
	}

	values = make(map[string]interface{})

	return values, nil
}

// SetKey TODO docs
func (m ConfigManager) SetKey(key string) {

}

// GetListOfKeys TODO docs
func (c Config) GetListOfKeys() []string {
	keys := make([]string, 0)
	extractFlatKeys("", c, &keys)
	sort.Strings(keys)
	return keys
}

// mapify TODO docs
func (c Config) mapify() map[string]interface{} {
	return structs.Map(c)
}

func extractFlatKeys(prefix string, v interface{}, keys *[]string) {
	val := reflect.ValueOf(v)
	if val.Kind() != reflect.Struct {
		return
	}

	var ns string

	for i := 0; i < val.Type().NumField(); i++ {
		field := val.Field(i)

		if prefix == "" {
			ns = val.Type().Field(i).Name
		} else {
			ns = fmt.Sprintf("%s.%s", prefix, strings.ToLower(val.Type().Field(i).Name))
		}

		if field.Kind() == reflect.Struct {
			extractFlatKeys(ns, field.Interface(), keys)
			continue
		}

		*keys = append(*keys, ns)
	}
}
