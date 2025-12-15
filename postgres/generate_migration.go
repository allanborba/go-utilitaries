//go:build ignore

// use sample:
// go run generate_migration.go create_user_table

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run generate_migration.go <migration_name>")
		fmt.Println("Example: go run generate_migration.go create_user_table")
		os.Exit(1)
	}

	migrationName := os.Args[1]
	today := time.Now().Format("20060102")

	// Count existing migrations with today's date
	entries, err := os.ReadDir("migrations")
	if err != nil {
		fmt.Printf("Error reading migrations directory: %v\n", err)
		os.Exit(1)
	}

	count := 0
	for _, entry := range entries {
		if strings.HasPrefix(entry.Name(), today) {
			count++
		}
	}

	sequence := fmt.Sprintf("%03d", count+1)
	filename := fmt.Sprintf("%s_%s_%s.sql", today, sequence, migrationName)
	path := filepath.Join("migrations", filename)

	file, err := os.Create(path)
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	fmt.Printf("Created migration: %s\n", path)
}
