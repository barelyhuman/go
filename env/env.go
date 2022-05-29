package env

import "os"

// Get the value from the env with a fallback default value
func Get(key string, defaultValue string) (result string) {
	result = os.Getenv(key)
	if len(result) == 0 {
		result = defaultValue
	}
	return
}

// Set - wrapper around os.Setenv to maintain consistency with the env API
func Set(key string, value string) error {
	return os.Setenv(key, value)
}
