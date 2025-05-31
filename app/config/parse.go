package config

import (
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/spf13/viper"
)

// Parse parses all configuration to a single Config object.
//
// `cfgDir` is the path of the directory where the config file is located.
// The config file should be named `config.yaml` and should be in the YAML format.
//
// The precedence of picking up of configuration values is as follows:
//  1. Environment variables
//  2. Config file
//  3. Default values
func Parse(cfgDir string) (*Config, error) {
	var config Config

	// set config directory
	dir := getConfigDir(cfgDir)

	// initialize viper
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(dir)

	viper.AutomaticEnv()
	viper.SetEnvPrefix("CATALYST")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// bind environment variables
	keys := keysOfStruct(config, "")
	for _, key := range keys {
		_ = viper.BindEnv(key)
	}

	// default values
	viper.SetDefault("app.metrics.enabled", true)
	viper.SetDefault("app.metrics.port", 8001)
	viper.SetDefault("app.metrics.route", "/metrics")

	// try to read the config file
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("using configs from ", viper.ConfigFileUsed())
	}

	// unmarshal into the struct
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("unable to decode into struct: %w", err)
	}

	return &config, nil
}

// getConfigDir returns config directory path after analyzing and correcting.
func getConfigDir(dir string) string {
	// get last char of dir path
	c := dir[len(dir)-1]
	if os.IsPathSeparator(c) {
		return dir
	}

	return dir + string(os.PathSeparator)
}

// keysOfStruct traverses a struct and generates all keys by concatenating mapstructure tags with '.'
func keysOfStruct(input any, prefix string) []string {
	var keys []string

	// Get the reflect type and value of the input
	t := reflect.TypeOf(input)
	v := reflect.ValueOf(input)

	// Ensure input is a struct
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}

	if t.Kind() != reflect.Struct {
		return keys
	}

	// Traverse each field of the struct
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)

		// Get the mapstructure tag
		tag, ok := field.Tag.Lookup("mapstructure")
		if !ok || tag == "" {
			continue
		}

		// Generate the full key by combining the prefix with the current tag
		fullKey := tag
		if prefix != "" {
			fullKey = prefix + "." + tag
		}

		// If the field is a struct, recursively traverse it
		if field.Type.Kind() == reflect.Struct || (field.Type.Kind() == reflect.Ptr && field.Type.Elem().Kind() == reflect.Struct) {
			childKeys := keysOfStruct(value.Interface(), fullKey)
			keys = append(keys, childKeys...)
		} else {
			// Otherwise, add the key to the result
			keys = append(keys, fullKey)
		}
	}

	return keys
}
