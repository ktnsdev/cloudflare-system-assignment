package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"time"
)

var Env *EnvironmentVariables

type EnvironmentVariables struct {
	Port      string
	JwtSecret string
	JwtPublic string
}

func Load() (err error) {
	err = godotenv.Load(".env")
	if err != nil {
		return err
	}

	receiveDotEnv()
	return nil
}

func getEnvironmentVariable(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func receiveDotEnv() {
	Env = &EnvironmentVariables{
		Port:      getEnvironmentVariable("PORT", ":3000"),
		JwtSecret: getEnvironmentVariable("JWT_SECRET", ""),
		JwtPublic: getEnvironmentVariable("JWT_PUBLIC", ""),
	}
}

func Logger(path string, statusCode int, log string) {
	fmt.Printf("%s\t%s\t%v %s\n", time.Now().Format(time.StampMilli), path, statusCode, log)
}
