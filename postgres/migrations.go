package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type Migrator struct {
	db *sql.DB
}

func NewMigrator(db *sql.DB) *Migrator {
	return &Migrator{db}
}

func (this *Migrator) MigrateIfNeeded() {
	this.createMigrationTableIfNotExists()
	files := this.getMigrationFiles()

	for _, file := range files {
		this.executeMigrationFromFile(file)
	}
}

func (this *Migrator) createMigrationTableIfNotExists() {
	query := `
		CREATE TABLE IF NOT EXISTS migrations (
			id SERIAL PRIMARY KEY,
			name VARCHAR(255) NOT NULL UNIQUE,
			executed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`
	_, err := this.db.Exec(query)
	if err != nil {
		panic(err)
	}
}

func (this *Migrator) getMigrationFiles() []string {
	files, err := os.ReadDir(this.getMigrationsDir())
	if err != nil {
		log.Fatalf("Erro ao ler diretório de migrações: %v", err)
	}

	var migrations []string
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".sql") {
			migrations = append(migrations, file.Name())
		}
	}

	sort.Strings(migrations)

	return migrations
}

func (this *Migrator) executeMigrationFromFile(file string) {
	alreadyExecuted := this.verifyIfAlredyExecuted(file)
	if alreadyExecuted {
		return
	}

	fmt.Println("Executing migration:", file)
	content := this.getContentFromFile(file)

	tx := this.begginMigration()
	defer tx.Rollback()

	this.execMigration(tx, content)
	this.insertOnMigrationsTable(tx, file)

	tx.Commit()

	log.Printf("Migration executed successfully: %s", file)
}

func (this *Migrator) verifyIfAlredyExecuted(file string) bool {
	var exists bool

	err := this.db.QueryRow("SELECT EXISTS(SELECT 1 FROM migrations WHERE name = $1)", file).Scan(&exists)
	if err != nil {
		panic(err)
	}

	return exists
}

func (this *Migrator) getContentFromFile(file string) []byte {
	content, err := os.ReadFile(filepath.Join(this.getMigrationsDir(), file))
	if err != nil {
		log.Fatalf("Erro ao ler arquivo %s: %v", file, err)
	}
	return content
}

func (this *Migrator) execMigration(tx *sql.Tx, content []byte) {
	_, err := tx.Exec(string(content))
	if err != nil {
		panic(err)
	}
}

func (this *Migrator) insertOnMigrationsTable(tx *sql.Tx, file string) {
	_, err := tx.Exec("INSERT INTO migrations (name) VALUES ($1)", file)
	if err != nil {
		panic(err)
	}
}

func (this *Migrator) begginMigration() *sql.Tx {
	tx, err := this.db.Begin()
	if err != nil {
		panic(err)
	}
	return tx
}

func (this *Migrator) getMigrationsDir() string {
	dir, _ := os.Getwd()
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return filepath.Join(dir, "migrations")
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			panic("Root do projeto não encontrado")
		}
		dir = parent
	}
}
