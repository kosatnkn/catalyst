package config

import (
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/spf13/viper"
)

// Parse parses all configuration values from environment vars and the `config.yaml` file
// in to the passed in config struct.
//
// The precedence for picking up of configuration values is as follows,
//  1. Environment variables prefixed with `Settings.Prefix`
//  2. `config.yaml` file in `Settings.Dir` (the file MUST be a YAML file with the name `config.yaml`)
//  3. Default values in `Settings.Defaults`
//
// The config struct 'MUST' contain `mapstructure` tags in order for parser to do the mapping.
// like so,
//
//	type MetricConfig struct {
//	  Enabled bool   `mapstructure:"enabled"`
//	  Port    int    `mapstructure:"port"`
//	  Route   string `mapstructure:"route"`
//	}
//
// Keys of the `Settings.Default` map should be dot (.) concatenated.
// For a config struct like this,
//
//	type Config struct {
//		App AppConfig `mapstructure:"app"`
//	}
//
//	type AppConfig struct {
//	 Name     string       `mapstructure:"name"`
//		Metrics  MetricConfig `mapstructure:"metrics"`
//	}
//
//	type MetricConfig struct {
//		Enabled bool   `mapstructure:"enabled"`
//		Port    int    `mapstructure:"port"`
//		Route   string `mapstructure:"route"`
//	}
//
// Use `app.metrics.enabled` as the key to set a default value for the metrics enabled configuration.
func Parse(config any, s Settings) (any, error) {
	// set config directory
	dir := getConfigDir(s.Dir)

	// initialize viper
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(dir)

	viper.AutomaticEnv()
	viper.SetEnvPrefix(s.Prefix)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// bind environment variables
	keys := keysOfStruct(config, "")
	for _, key := range keys {
		_ = viper.BindEnv(key)
	}

	// default values
	for k, v := range s.Defaults {
		viper.SetDefault(k, v)
	}

	// try to read the config file
	if err := viper.ReadInConfig(); err != nil {
		return config, fmt.Errorf("config: unable to read, %w", err)
	}

	// unmarshal into the struct
	if err := viper.Unmarshal(&config); err != nil {
		return config, fmt.Errorf("config: unable to unmarshal, %w", err)
	}

	return config, nil
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

// keysOfStruct traverses a struct and generates a slice of keys by concatenating 'mapstructure' tags with '.'
func keysOfStruct(input any, prefix string) []string {
	var keys []string

	// get the reflect type and value of the input
	t := reflect.TypeOf(input)
	v := reflect.ValueOf(input)

	// ensure input is a struct
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}

	if t.Kind() != reflect.Struct {
		return keys
	}

	// traverse each field of the struct
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)

		// get the 'mapstructure' tag
		tag, ok := field.Tag.Lookup("mapstructure")
		if !ok || tag == "" {
			continue
		}

		// generate the full key by combining the prefix with the current tag
		fullKey := tag
		if prefix != "" {
			fullKey = prefix + "." + tag
		}

		// if the field is a struct, recursively traverse it
		if field.Type.Kind() == reflect.Struct || (field.Type.Kind() == reflect.Ptr && field.Type.Elem().Kind() == reflect.Struct) {
			childKeys := keysOfStruct(value.Interface(), fullKey)
			keys = append(keys, childKeys...)
		} else {
			// otherwise, add the key to the result
			keys = append(keys, fullKey)
		}
	}

	return keys
}
