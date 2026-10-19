package config

import (
	"testing"
)

func TestLoadDefaults(t *testing.T) {
	t.Setenv("SERVER_PORT", "")
	t.Setenv("DB_HOST", "")
	t.Setenv("DB_PORT", "")
	t.Setenv("DB_USER", "")
	t.Setenv("DB_PASSWORD", "")
	t.Setenv("DB_NAME", "")
	t.Setenv("DB_SSLMODE", "")

	cfg, err := Load()
	if err != nil {
		t.Fatalf("Load() returned an error: %v", err)
	}

	if cfg.ServerPort != "8080" {
		t.Errorf("ServerPort: got %q, want %q", cfg.ServerPort, "8080")
	}

	if cfg.DBHost != "localhost" {
		t.Errorf("DBHost: got %q, want %q", cfg.DBHost, "localhost")
	}

	if cfg.DBPort != "5432" {
		t.Errorf("DBPort: got %q, want %q", cfg.DBPort, "5432")
	}

	if cfg.DBUser != "citysesi" {
		t.Errorf("DBUser: got %q, want %q", cfg.DBUser, "citysesi")
	}

	if cfg.DBPassword != "" {
		t.Errorf("DBPassword: expected an empty string")
	}

	if cfg.DBName != "citysesi" {
		t.Errorf("DBName: got %q, want %q", cfg.DBName, "citysesi")
	}

	if cfg.DBSSLMode != "disable" {
		t.Errorf("DBSSLMode: got %q, want %q", cfg.DBSSLMode, "disable")
	}
}

func TestLoadCustomValues(t *testing.T) {
	t.Setenv("SERVER_PORT", "9090")
	t.Setenv("DB_HOST", "db.example.local")

	cfg, err := Load()
	if err != nil {
		t.Fatalf("Load() returned an error: %v", err)
	}

	if cfg.ServerPort != "9090" {
		t.Errorf("ServerPort: got %q, want %q", cfg.ServerPort, "9090")
	}

	if cfg.DBHost != "db.example.local" {
		t.Errorf("DBHost: got %q, want %q", cfg.DBHost, "db.example.local")
	}
}
