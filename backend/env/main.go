package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

// ENVIRONMENT VARIABLES IN GO
//
// In production backends, config (DB URLs, API keys, ports)
// NEVER gets hardcoded. It lives in environment variables.
//
// Go's os package provides:
//   os.Getenv("KEY")             → returns value or ""
//   os.LookupEnv("KEY")         → returns value + bool (exists?)
//   os.Setenv("KEY", "val")     → set a variable

type Config struct {
	Port     string
	DBUrl    string
	APIKey   string
	Debug    bool
	MaxConns int
}

// getEnv reads an env var, falls back to a default
func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}

// getEnvBool reads a bool env var (e.g. "true" / "false")
func getEnvBool(key string, fallback bool) bool {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	b, err := strconv.ParseBool(val)
	if err != nil {
		return fallback
	}
	return b
}

// getEnvInt reads an int env var
func getEnvInt(key string, fallback int) int {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	n, err := strconv.Atoi(val)
	if err != nil {
		return fallback
	}
	return n
}

// LoadConfig reads all config from env at startup
func LoadConfig() Config {
	return Config{
		Port:     getEnv("PORT", "8080"),
		DBUrl:    getEnv("DATABASE_URL", "postgres://localhost/mydb"),
		APIKey:   getEnv("API_KEY", ""),
		Debug:    getEnvBool("DEBUG", false),
		MaxConns: getEnvInt("MAX_CONNECTIONS", 10),
	}
}

func main() {

	// --- 1. Basic read/write ---
	fmt.Println("=== os.Getenv ===")
	os.Setenv("APP_NAME", "GoAPI")

	name := os.Getenv("APP_NAME")
	fmt.Println("APP_NAME:", name)

	missing := os.Getenv("DOES_NOT_EXIST")
	fmt.Printf("DOES_NOT_EXIST: %q (empty string if not set)\n", missing)

	// --- 2. LookupEnv — check if a key exists ---
	fmt.Println("\n=== LookupEnv (safer) ===")
	if val, ok := os.LookupEnv("APP_NAME"); ok {
		fmt.Println("Found APP_NAME:", val)
	}

	if _, ok := os.LookupEnv("DOES_NOT_EXIST"); !ok {
		fmt.Println("DOES_NOT_EXIST is not set")
	}

	// --- 3. Require a key (crash if missing) ---
	fmt.Println("\n=== Required env var ===")
	requireEnv := func(key string) string {
		val := os.Getenv(key)
		if val == "" {
			log.Fatalf("Required env var %q is not set", key)
		}
		return val
	}
	_ = requireEnv // don't call in demo — it would crash

	// --- 4. Load full config at startup ---
	fmt.Println("\n=== Config struct ===")
	cfg := LoadConfig()
	fmt.Printf("Port:     %s\n", cfg.Port)
	fmt.Printf("DB URL:   %s\n", cfg.DBUrl)
	fmt.Printf("Debug:    %v\n", cfg.Debug)
	fmt.Printf("MaxConns: %d\n", cfg.MaxConns)
	if cfg.APIKey == "" {
		fmt.Println("API Key:  (not set)")
	}

	// --- 5. How to use a .env file (in real projects) ---
	// Use the godotenv package: github.com/joho/godotenv
	// go get github.com/joho/godotenv
	//
	// Then in main():
	//   import "github.com/joho/godotenv"
	//   godotenv.Load()  // loads .env into env vars
	//   cfg := LoadConfig()
	//
	// Your .env file looks like:
	//   PORT=9090
	//   DATABASE_URL=postgres://user:pass@localhost/mydb
	//   API_KEY=my-secret-key
	//   DEBUG=true
	fmt.Println("\n=== .env pattern (with godotenv) ===")
	fmt.Println("  go get github.com/joho/godotenv")
	fmt.Println("  godotenv.Load() // call once at startup")

	// ============================================================
	// QUICK REFERENCE:
	//
	//  os.Getenv("KEY")             → value or ""
	//  os.LookupEnv("KEY")         → value, bool
	//  os.Setenv("KEY", "value")   → set
	//
	//  Best practice:
	//  1. Load all env vars into a Config struct at startup
	//  2. Use getEnv(key, default) for optional vars
	//  3. Crash loudly if required vars are missing
	//  4. Never commit .env to git — add it to .gitignore!
	// ============================================================
}
