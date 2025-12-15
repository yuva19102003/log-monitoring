package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"
)

/* ---------- LOAD .ENV ---------- */
func loadEnv() {
	file, err := os.Open(".env")
	if err != nil {
		log.Println("No .env file found, using environment variables")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			os.Setenv(parts[0], parts[1])
		}
	}
}

/* ---------- CREATE LOGGER ---------- */
func createLogger(filePath string) *log.Logger {
	dir := filepath.Dir(filePath)

	if err := os.MkdirAll(dir, 0755); err != nil {
		panic(err)
	}

	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	return log.New(file, "", 0)
}

/* ---------- SYSLOG FORMAT ---------- */
func writeLog(logger *log.Logger, service, level, message string) {
	hostname, _ := os.Hostname()
	pid := os.Getpid()
	timestamp := time.Now().Format("2006-01-02T15:04:05.000000-07:00")

	line := fmt.Sprintf(
		"%s %s %s[%d]: %s %s",
		timestamp,
		hostname,
		service,
		pid,
		level,
		message,
	)

	logger.Println(line)
}

func main() {
	loadEnv()
	rand.Seed(time.Now().UnixNano())

	frontend := createLogger(os.Getenv("FRONTEND_LOG"))
	backend := createLogger(os.Getenv("BACKEND_LOG"))
	database := createLogger(os.Getenv("DATABASE_LOG"))
	auth := createLogger(os.Getenv("AUTH_LOG"))
	system := createLogger(os.Getenv("SYSTEM_LOG"))

	for {
		/* ---------- FRONTEND ---------- */
		writeLog(frontend, "frontend", "INFO", "Page rendered")
		writeLog(frontend, "frontend", "DEBUG", "Assets loaded")

		/* ---------- BACKEND ---------- */
		writeLog(backend, "backend", "INFO", "API request handled")
		if rand.Intn(10) > 6 {
			writeLog(backend, "backend", "ERROR", "Service panic occurred")
		}

		/* ---------- DATABASE ---------- */
		writeLog(database, "database", "DEBUG", "Query executed")
		if rand.Intn(10) > 7 {
			writeLog(database, "database", "ERROR", "Transaction rollback")
		}

		/* ---------- AUTH ---------- */
		writeLog(auth, "auth", "INFO", "Login attempt")
		if rand.Intn(10) > 6 {
			writeLog(auth, "auth", "WARN", "Invalid password")
		}

		/* ---------- SYSTEM ---------- */
		writeLog(system, "system", "DEBUG", "CPU metrics collected")
		if rand.Intn(50) == 25 {
			writeLog(system, "system", "FATAL", "Out of memory")
		}

		time.Sleep(2 * time.Second)
	}
}
