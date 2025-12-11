package env

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func LoadEnvs() {
	fmt.Println("Loading envs")

	loadEnvs(".env")
	loadEnvs(getEnvFileName()) // .env.development, .env.production, etc
	loadEnvs(".env.local")
	loadEnvs(getLocalEnvFileName()) // .env.development.local, .env.production.local, etc
}

func loadEnvs(fileName string) {
	envPath := findEnvFile(fileName)
	if envPath == "" {
		return
	}

	file, err := os.Open(envPath)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		os.Setenv(key, value)
	}
}

func findEnvFile(fileName string) string {
	dir, err := os.Getwd()
	if err != nil {
		return ""
	}

	for {
		envPath := filepath.Join(dir, fileName)
		if _, err := os.Stat(envPath); err == nil {
			return envPath
		}

		parentDir := filepath.Dir(dir)
		if parentDir == dir {
			return ""
		}
		dir = parentDir
	}
}

func getEnvFileName() string {
	envName := os.Getenv("GO_ENV")
	if envName == "" {
		envName = "development"
	}

	return fmt.Sprintf(".env.%s", envName)
}

func getLocalEnvFileName() string {
	envName := os.Getenv("GO_ENV")
	if envName == "" {
		envName = "development"
	}

	return fmt.Sprintf(".env.%s.local", envName)
}
