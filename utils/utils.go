package utils

import (
	"log"
	"os"
)

func GetEnv(name string, default_ string) string {
	env_var := os.Getenv(name)
	if env_var == "" && default_ == "" {
		log.Panic("Could not find required environment variable")
	} else if env_var == "" {
		env_var = default_
	}
	return env_var
}
