package env

import (
	"os"
	"strings"
)

func Get(key string) string {
	return strings.TrimSpace(os.Getenv(key))
}
