package env

import(
	"fmt"
	"os"
)

func Get(key string) string {
	if val :=os.Getenv(key); val == "" {
		panic(fmt.Sprintf("No value found for key: %s", key))
	} else {
		return val
	}
}

func GetDefault(key, def string) string {
	if val := os.Getenv(key); val == "" {
		return def
	} else {
		return val
	}
}