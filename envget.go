// Package envget provides functionality that wraps os.Getenv.
package envget

import (
	"os"
	"strconv"
	"strings"
	"time"
)

// GetString retrieves the value of the environment variable named by the key as string.
// If the variable is empty or not exists, fallbacks to the given.
func GetString(key string, fallback string) string {
	v := os.Getenv(key)
	if v == "" {
		return fallback
	}
	return v
}

// GetInt retrieves the value of the environment variable named by the key as int.
// If the variable is empty or not exists, fallbacks to the given.
func GetInt(key string, fallback int) int {
	v := os.Getenv(key)
	if v == "" {
		return fallback
	}
	_v, err := strconv.Atoi(v)
	if err != nil {
		return fallback
	}
	return _v
}

// GetBool retrieves the value of the environment variable named by the key as bool.
// If the variable is empty or not exists, fallbacks to the given.
func GetBool(key string, fallback bool) bool {
	v := os.Getenv(key)
	if v == "" {
		return fallback
	}
	_v, err := strconv.ParseBool(v)
	if err != nil {
		return fallback
	}
	return _v
}

// GetDuration retrieves the value of the environment variable named by the key as time.Duration.
// If the variable is empty or not exists, fallbacks to the given.
func GetDuration(key string, fallback time.Duration) time.Duration {
	v := os.Getenv(key)
	if v == "" {
		return fallback
	}
	_v, err := time.ParseDuration(v)
	if err != nil {
		return fallback
	}
	return _v
}

// GetStringSlice retrieves the value of the environment variable named by the key as []string.
// If the variable is not exists, fallbacks to the given.
func GetStringSlice(key string, fallback []string) []string {
	v := os.Getenv(key)
	if v == "" {
		return fallback
	}
	list := strings.Split(v, ",")
	_list := make([]string, 0, len(list))
	for _, s := range list {
		if s != "" {
			_list = append(_list, strings.TrimSpace(s))
		}
	}
	return _list
}
